package service

import (
	"context"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"github.com/svenskhalsovard/api/internal/entity"
	"github.com/svenskhalsovard/api/internal/svea"
)

// PaymentRepository defines the interface for payment data operations
type PaymentRepository interface {
	CreatePayment(ctx context.Context, tx *sqlx.Tx, payment *entity.Payment, items []entity.PaymentItem) error
	GetPaymentByID(ctx context.Context, id int64) (*entity.Payment, error)
	GetPaymentByExternalID(ctx context.Context, externalID string) (*entity.Payment, error)
	GetPaymentByOrderReference(ctx context.Context, orderReference string) (*entity.Payment, error)
	GetPaymentWithItems(ctx context.Context, id int64) (*entity.PaymentWithItems, error)
	UpdatePaymentStatus(ctx context.Context, id int64, status string, errorMessage string) error
	UpdatePaymentExternalID(ctx context.Context, id int64, externalID string) error
	FindIncompletePayments(ctx context.Context, maxAge string) ([]entity.Payment, error)
	Transaction(fn func(*sqlx.Tx) error) error
}

// SveaClient defines the interface for Svea Ekonomi API client
type SveaClient interface {
	CreateOrder(ctx context.Context, order *svea.OrderRequest) (*svea.OrderResponse, error)
	GetOrder(ctx context.Context, orderID string) (*svea.Order, error)
	FinalizePayment(ctx context.Context, orderID string, paymentMethod string) (*svea.PaymentResponse, error)
}

// PaymentService provides business logic for payments
type PaymentService struct {
	repo       PaymentRepository
	sveaClient SveaClient
}

// NewPaymentService creates a new PaymentService
func NewPaymentService(repo PaymentRepository, sveaClient SveaClient) *PaymentService {
	return &PaymentService{
		repo:       repo,
		sveaClient: sveaClient,
	}
}

// InitiatePayment starts the payment process
func (s *PaymentService) InitiatePayment(ctx context.Context, customer *entity.Customer, items []entity.PaymentItem, options *PaymentOptions) (*entity.Payment, error) {
	// Generate unique order reference
	orderReference := fmt.Sprintf("ORD-%d", time.Now().UnixNano())

	// Create payment record in database
	payment := &entity.Payment{
		CustomerID:      customer.ID,
		Amount:          options.TotalAmount,
		Currency:        entity.CurrencySEK,
		Status:          entity.PaymentStatusInitiated,
		OrderReference:  orderReference,
		TransactionType: options.TransactionType,
	}

	// Start database transaction
	err := s.repo.Transaction(func(tx *sqlx.Tx) error {
		if err := s.repo.CreatePayment(ctx, tx, payment, items); err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		log.Error().Err(err).Msg("Failed to initiate payment")
		return nil, fmt.Errorf("failed to initiate payment: %w", err)
	}

	return payment, nil
}

// CreateSveaOrder creates a new order in Svea Ekonomi
func (s *PaymentService) CreateSveaOrder(ctx context.Context, paymentID int64) (*svea.OrderResponse, error) {
	// Get payment data
	paymentWithItems, err := s.repo.GetPaymentWithItems(ctx, paymentID)
	if err != nil {
		log.Error().Err(err).Int64("paymentID", paymentID).Msg("Failed to get payment data")
		return nil, fmt.Errorf("failed to get payment data: %w", err)
	}

	if paymentWithItems == nil {
		return nil, fmt.Errorf("payment not found")
	}

	// Prepare Svea order request
	orderRequest := &svea.OrderRequest{
		OrderReference: paymentWithItems.Payment.OrderReference,
		Currency:       paymentWithItems.Payment.Currency,
		ClientOrderNumber: fmt.Sprintf("ORD-%d", paymentWithItems.Payment.ID),
		PaymentType:    s.mapTransactionType(paymentWithItems.Payment.TransactionType),
		Customer: svea.Customer{
			FirstName:     paymentWithItems.Customer.FirstName,
			LastName:      paymentWithItems.Customer.LastName,
			Email:         paymentWithItems.Customer.Email,
			PhoneNumber:   paymentWithItems.Customer.Phone,
			Address: svea.Address{
				StreetAddress: paymentWithItems.Customer.StreetAddress,
				PostalCode:    paymentWithItems.Customer.PostalCode,
				City:          paymentWithItems.Customer.City,
				CountryCode:   "SE",
			},
		},
		Items: make([]svea.OrderItem, 0, len(paymentWithItems.Items)),
	}

	// Add order items
	for _, item := range paymentWithItems.Items {
		orderRequest.Items = append(orderRequest.Items, svea.OrderItem{
			ArticleNumber: fmt.Sprintf("SRV-%d", item.ServiceID),
			Name:          item.ServiceName,
			Quantity:      item.Quantity,
			UnitPrice:     int(item.UnitPrice * 100), // Convert to cents/öre
			VatPercent:    25,                       // Assuming 25% VAT
			Unit:          "st",
		})
	}

	// Create order in Svea
	orderResponse, err := s.sveaClient.CreateOrder(ctx, orderRequest)
	if err != nil {
		// Update payment status to failed
		_ = s.repo.UpdatePaymentStatus(ctx, paymentID, entity.PaymentStatusFailed, err.Error())
		log.Error().Err(err).Int64("paymentID", paymentID).Msg("Failed to create Svea order")
		return nil, fmt.Errorf("failed to create Svea order: %w", err)
	}

	// Update payment with external ID
	err = s.repo.UpdatePaymentExternalID(ctx, paymentID, orderResponse.OrderID)
	if err != nil {
		log.Error().Err(err).Int64("paymentID", paymentID).Str("orderID", orderResponse.OrderID).Msg("Failed to update payment with external ID")
	}

	// Update payment status to pending
	err = s.repo.UpdatePaymentStatus(ctx, paymentID, entity.PaymentStatusPending, "")
	if err != nil {
		log.Error().Err(err).Int64("paymentID", paymentID).Msg("Failed to update payment status")
	}

	return orderResponse, nil
}

// ProcessPayment processes a payment with Svea Ekonomi
func (s *PaymentService) ProcessPayment(ctx context.Context, paymentID int64, paymentMethod string) error {
	// Get payment data
	payment, err := s.repo.GetPaymentByID(ctx, paymentID)
	if err != nil {
		log.Error().Err(err).Int64("paymentID", paymentID).Msg("Failed to get payment")
		return fmt.Errorf("failed to get payment: %w", err)
	}

	if payment == nil {
		return fmt.Errorf("payment not found")
	}

	if payment.Status != entity.PaymentStatusPending {
		return fmt.Errorf("payment cannot be processed (status: %s)", payment.Status)
	}

	if payment.ExternalPaymentID == "" {
		return fmt.Errorf("payment has no external ID")
	}

	// Finalize payment in Svea
	_, err = s.sveaClient.FinalizePayment(ctx, payment.ExternalPaymentID, paymentMethod)
	if err != nil {
		// Update payment status to failed
		_ = s.repo.UpdatePaymentStatus(ctx, paymentID, entity.PaymentStatusFailed, err.Error())
		log.Error().Err(err).Int64("paymentID", paymentID).Msg("Failed to process payment")
		return fmt.Errorf("failed to process payment: %w", err)
	}

	// Update payment status and method
	payment.Status = entity.PaymentStatusSuccess
	payment.PaymentMethod = paymentMethod
	err = s.repo.UpdatePaymentStatus(ctx, paymentID, payment.Status, "")
	if err != nil {
		log.Error().Err(err).Int64("paymentID", paymentID).Msg("Failed to update payment status")
		return fmt.Errorf("payment processed but failed to update status: %w", err)
	}

	return nil
}

// VerifyPayment checks the status of a payment with Svea Ekonomi
func (s *PaymentService) VerifyPayment(ctx context.Context, paymentID int64) (*entity.Payment, error) {
	// Get payment data
	payment, err := s.repo.GetPaymentByID(ctx, paymentID)
	if err != nil {
		log.Error().Err(err).Int64("paymentID", paymentID).Msg("Failed to get payment")
		return nil, fmt.Errorf("failed to get payment: %w", err)
	}

	if payment == nil {
		return nil, fmt.Errorf("payment not found")
	}

	// If payment is already in a final state, return it
	if payment.Status == entity.PaymentStatusSuccess || payment.Status == entity.PaymentStatusFailed {
		return payment, nil
	}

	// If no external ID, payment hasn't been sent to Svea yet
	if payment.ExternalPaymentID == "" {
		return payment, nil
	}

	// Check order status in Svea
	order, err := s.sveaClient.GetOrder(ctx, payment.ExternalPaymentID)
	if err != nil {
		log.Error().Err(err).Int64("paymentID", paymentID).Str("externalID", payment.ExternalPaymentID).Msg("Failed to get order from Svea")
		return payment, nil // Return current payment info, don't update status on error
	}

	// Map Svea order status to our payment status
	newStatus := s.mapSveaOrderStatus(order.Status)
	if newStatus != payment.Status {
		// Update payment status
		err = s.repo.UpdatePaymentStatus(ctx, paymentID, newStatus, "")
		if err != nil {
			log.Error().Err(err).Int64("paymentID", paymentID).Str("status", newStatus).Msg("Failed to update payment status")
			return payment, nil // Return current payment info, don't fail on status update error
		}
		payment.Status = newStatus
	}

	return payment, nil
}

// Helper methods

// PaymentOptions contains options for initiating a payment
type PaymentOptions struct {
	TotalAmount     float64
	TransactionType string
}

// mapTransactionType maps our transaction type to Svea's payment type
func (s *PaymentService) mapTransactionType(transactionType string) string {
	switch transactionType {
	case entity.TransactionTypeSubscription:
		return "recurring"
	default:
		return "onetime"
	}
}

// mapSveaOrderStatus maps Svea order status to our payment status
func (s *PaymentService) mapSveaOrderStatus(sveaStatus string) string {
	switch sveaStatus {
	case "Created", "Pending":
		return entity.PaymentStatusPending
	case "Completed", "Paid":
		return entity.PaymentStatusSuccess
	case "Failed", "Cancelled":
		return entity.PaymentStatusFailed
	default:
		return entity.PaymentStatusPending
	}
}