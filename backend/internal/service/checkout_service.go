package service

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"
	"github.com/svenskhalsovard/api/internal/entity"
	"github.com/svenskhalsovard/api/internal/svea"
)

// CheckoutService provides business logic for the checkout process
type CheckoutService struct {
	paymentService *PaymentService
	bookingService *BookingService
	serviceService *ServiceService
}

// NewCheckoutService creates a new CheckoutService
func NewCheckoutService(
	paymentService *PaymentService,
	bookingService *BookingService,
	serviceService *ServiceService,
) *CheckoutService {
	return &CheckoutService{
		paymentService: paymentService,
		bookingService: bookingService,
		serviceService: serviceService,
	}
}

// CheckoutItem represents an item in a checkout request
type CheckoutItem struct {
	ServiceID    int64   `json:"serviceId"`
	Quantity     int     `json:"quantity"`
	PurchaseType string  `json:"purchaseType"`
	Price        float64 `json:"price"`
}

// CheckoutRequest represents a checkout request
type CheckoutRequest struct {
	Customer    entity.Customer `json:"customer"`
	Items       []CheckoutItem  `json:"items"`
	TotalAmount float64         `json:"totalAmount"`
}

// CheckoutResult represents the result of a checkout
type CheckoutResult struct {
	PaymentID      int64                `json:"paymentID"`
	SveaOrderID    string               `json:"sveaOrderID"`
	SveaCheckoutUI svea.CheckoutUIData  `json:"sveaCheckoutUI"`
	Status         string               `json:"status"`
	Customer       entity.Customer      `json:"customer"`
	Items          []entity.PaymentItem `json:"items"`
}

// InitiateCheckout handles the checkout process
func (s *CheckoutService) InitiateCheckout(ctx context.Context, req *CheckoutRequest) (*CheckoutResult, error) {
	// Validate items
	if len(req.Items) == 0 {
		return nil, fmt.Errorf("no items in checkout")
	}

	// Get service IDs
	serviceIDs := make([]int64, 0, len(req.Items))
	for _, item := range req.Items {
		serviceIDs = append(serviceIDs, item.ServiceID)
	}

	// Validate services
	services, err := s.serviceService.ValidateServices(ctx, serviceIDs)
	if err != nil {
		log.Error().Err(err).Interface("serviceIDs", serviceIDs).Msg("Invalid services in checkout")
		return nil, fmt.Errorf("invalid services: %w", err)
	}

	// Prepare payment items
	paymentItems := make([]entity.PaymentItem, 0, len(req.Items))
	var hasSubscription bool

	for _, item := range req.Items {
		service, ok := services[item.ServiceID]
		if !ok {
			return nil, fmt.Errorf("service with ID %d not found", item.ServiceID)
		}

		// Validate purchase type based on service
		if item.PurchaseType == "subscription" && !service.IsSubscription {
			return nil, fmt.Errorf("service with ID %d does not support subscription", item.ServiceID)
		}

		// Track if we have any subscription items
		if item.PurchaseType == "subscription" {
			hasSubscription = true
		}

		// Validate price matches service
		expectedPrice := service.DiscountedPrice
		if expectedPrice == nil {
			expectedPrice = &service.Price
		}
		
		// Allow small difference in price to account for decimal precision issues
		if !almostEqual(item.Price, *expectedPrice, 0.01) {
			return nil, fmt.Errorf("price mismatch for service %s: expected %.2f, got %.2f", 
				service.Name, *expectedPrice, item.Price)
		}

		// Validate quantity
		if item.Quantity <= 0 || item.Quantity > 10 {
			return nil, fmt.Errorf("invalid quantity for service %s: %d", service.Name, item.Quantity)
		}

		// Create payment item
		paymentItems = append(paymentItems, entity.PaymentItem{
			ServiceID:    service.ID,
			ServiceName:  service.Name,
			Quantity:     item.Quantity,
			UnitPrice:    item.Price,
			TotalPrice:   item.Price * float64(item.Quantity),
			PurchaseType: item.PurchaseType,
		})
	}

	// Validate total amount
	calculatedTotal := 0.0
	for _, item := range paymentItems {
		calculatedTotal += item.TotalPrice
	}

	if !almostEqual(req.TotalAmount, calculatedTotal, 0.01) {
		return nil, fmt.Errorf("total amount mismatch: expected %.2f, got %.2f", 
			calculatedTotal, req.TotalAmount)
	}

	// Determine transaction type
	transactionType := entity.TransactionTypeOneTime
	if hasSubscription {
		transactionType = entity.TransactionTypeSubscription
	}

	// Initiate payment
	payment, err := s.paymentService.InitiatePayment(ctx, &req.Customer, paymentItems, &PaymentOptions{
		TotalAmount:     req.TotalAmount,
		TransactionType: transactionType,
	})
	if err != nil {
		log.Error().Err(err).Msg("Failed to initiate payment")
		return nil, fmt.Errorf("failed to initiate payment: %w", err)
	}

	// Create Svea order
	orderResponse, err := s.paymentService.CreateSveaOrder(ctx, payment.ID)
	if err != nil {
		log.Error().Err(err).Int64("paymentID", payment.ID).Msg("Failed to create Svea order")
		return nil, fmt.Errorf("failed to create Svea order: %w", err)
	}

	// Return checkout result
	return &CheckoutResult{
		PaymentID:      payment.ID,
		SveaOrderID:    orderResponse.OrderID,
		SveaCheckoutUI: orderResponse.CheckoutUI,
		Status:         payment.Status,
		Customer:       req.Customer,
		Items:          paymentItems,
	}, nil
}

// ProcessPayment processes a payment in Svea Ekonomi
func (s *CheckoutService) ProcessPayment(ctx context.Context, paymentID int64, paymentMethod string) error {
	// Process payment
	err := s.paymentService.ProcessPayment(ctx, paymentID, paymentMethod)
	if err != nil {
		log.Error().Err(err).Int64("paymentID", paymentID).Str("method", paymentMethod).Msg("Failed to process payment")
		return fmt.Errorf("failed to process payment: %w", err)
	}

	// Create booking for successful payment
	_, err = s.bookingService.CreateBooking(ctx, paymentID, nil)
	if err != nil {
		log.Error().Err(err).Int64("paymentID", paymentID).Msg("Failed to create booking after successful payment")
		return fmt.Errorf("payment successful but failed to create booking: %w", err)
	}

	return nil
}

// VerifyPayment checks the status of a payment
func (s *CheckoutService) VerifyPayment(ctx context.Context, paymentID int64) (*entity.Payment, error) {
	// Verify payment
	payment, err := s.paymentService.VerifyPayment(ctx, paymentID)
	if err != nil {
		log.Error().Err(err).Int64("paymentID", paymentID).Msg("Failed to verify payment")
		return nil, fmt.Errorf("failed to verify payment: %w", err)
	}

	// If payment is successful, ensure a booking exists
	if payment.Status == entity.PaymentStatusSuccess {
		// Check if booking already exists
		existingBooking, err := s.bookingService.bookingRepo.GetBookingByPaymentID(ctx, paymentID)
		if err != nil {
			log.Error().Err(err).Int64("paymentID", paymentID).Msg("Failed to check for existing booking")
			return payment, nil // Return payment info despite booking check error
		}

		if existingBooking == nil {
			// Create booking for successful payment
			_, err = s.bookingService.CreateBooking(ctx, paymentID, nil)
			if err != nil {
				log.Error().Err(err).Int64("paymentID", paymentID).Msg("Failed to create booking during payment verification")
				return payment, nil // Return payment info despite booking creation error
			}
		}
	}

	return payment, nil
}

// Helper function to compare floats with tolerance
func almostEqual(a, b, tolerance float64) bool {
	diff := a - b
	if diff < 0 {
		diff = -diff
	}
	return diff <= tolerance
}