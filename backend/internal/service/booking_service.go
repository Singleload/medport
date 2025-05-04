package service

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"github.com/svenskhalsovard/api/internal/entity"
)

// BookingRepository defines the interface for booking data operations
type BookingRepository interface {
	CreateBooking(ctx context.Context, tx *sqlx.Tx, booking *entity.Booking, customer *entity.Customer, items []entity.BookingItem) error
	GetBookingByID(ctx context.Context, id int64) (*entity.Booking, error)
	GetBookingByPaymentID(ctx context.Context, paymentID int64) (*entity.Booking, error)
	GetBookingWithItems(ctx context.Context, id int64) (*entity.BookingWithItems, error)
	UpdateBookingStatus(ctx context.Context, bookingID int64, status string) error
	Transaction(fn func(*sqlx.Tx) error) error
}

// BookingService provides business logic for bookings
type BookingService struct {
	bookingRepo BookingRepository
	paymentRepo PaymentRepository
}

// NewBookingService creates a new BookingService
func NewBookingService(bookingRepo BookingRepository, paymentRepo PaymentRepository) *BookingService {
	return &BookingService{
		bookingRepo: bookingRepo,
		paymentRepo: paymentRepo,
	}
}

// CreateBooking creates a new booking based on a successful payment
func (s *BookingService) CreateBooking(ctx context.Context, paymentID int64, customer *entity.Customer) (*entity.Booking, error) {
	// Get payment data
	paymentWithItems, err := s.paymentRepo.GetPaymentWithItems(ctx, paymentID)
	if err != nil {
		log.Error().Err(err).Int64("paymentID", paymentID).Msg("Failed to get payment data")
		return nil, fmt.Errorf("failed to get payment data: %w", err)
	}

	if paymentWithItems == nil {
		return nil, fmt.Errorf("payment not found")
	}

	// Verify payment is successful
	if paymentWithItems.Payment.Status != entity.PaymentStatusSuccess {
		return nil, fmt.Errorf("payment is not successful (status: %s)", paymentWithItems.Payment.Status)
	}

	// Check if booking already exists for this payment
	existingBooking, err := s.bookingRepo.GetBookingByPaymentID(ctx, paymentID)
	if err != nil {
		log.Error().Err(err).Int64("paymentID", paymentID).Msg("Failed to check for existing booking")
		return nil, fmt.Errorf("failed to check for existing booking: %w", err)
	}

	if existingBooking != nil {
		// Booking already exists for this payment
		booking, err := s.bookingRepo.GetBookingWithItems(ctx, existingBooking.ID)
		if err != nil {
			log.Error().Err(err).Int64("bookingID", existingBooking.ID).Msg("Failed to get existing booking")
			return nil, fmt.Errorf("failed to get existing booking: %w", err)
		}
		return &booking.Booking, nil
	}

	// Create booking items from payment items
	bookingItems := make([]entity.BookingItem, 0, len(paymentWithItems.Items))
	for _, paymentItem := range paymentWithItems.Items {
		bookingItems = append(bookingItems, entity.BookingItem{
			ServiceID:      paymentItem.ServiceID,
			ServiceName:    paymentItem.ServiceName,
			Quantity:       paymentItem.Quantity,
			UnitPrice:      paymentItem.UnitPrice,
			TotalPrice:     paymentItem.TotalPrice,
			PurchaseType:   paymentItem.PurchaseType,
			IsSubscription: paymentItem.PurchaseType == entity.TransactionTypeSubscription,
		})
	}

	// Create booking record
	booking := &entity.Booking{
		PaymentID:   paymentID,
		Status:      entity.BookingStatusConfirmed,
		TotalAmount: paymentWithItems.Payment.Amount,
	}

	// Use the customer data from the request if provided, otherwise use the customer from the payment
	if customer == nil {
		customer = &paymentWithItems.Customer
	}

	// Start database transaction
	err = s.bookingRepo.Transaction(func(tx *sqlx.Tx) error {
		if err := s.bookingRepo.CreateBooking(ctx, tx, booking, customer, bookingItems); err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		log.Error().Err(err).Int64("paymentID", paymentID).Msg("Failed to create booking")
		return nil, fmt.Errorf("failed to create booking: %w", err)
	}

	return booking, nil
}

// GetBooking retrieves a booking by ID
func (s *BookingService) GetBooking(ctx context.Context, id int64) (*entity.BookingWithItems, error) {
	booking, err := s.bookingRepo.GetBookingWithItems(ctx, id)
	if err != nil {
		log.Error().Err(err).Int64("bookingID", id).Msg("Failed to get booking")
		return nil, fmt.Errorf("failed to get booking: %w", err)
	}

	if booking == nil {
		return nil, fmt.Errorf("booking not found")
	}

	return booking, nil
}

// UpdateBookingStatus updates the status of a booking
func (s *BookingService) UpdateBookingStatus(ctx context.Context, id int64, status string) error {
	if !isValidBookingStatus(status) {
		return fmt.Errorf("invalid booking status: %s", status)
	}

	err := s.bookingRepo.UpdateBookingStatus(ctx, id, status)
	if err != nil {
		log.Error().Err(err).Int64("bookingID", id).Str("status", status).Msg("Failed to update booking status")
		return fmt.Errorf("failed to update booking status: %w", err)
	}

	return nil
}

// Helper functions

// isValidBookingStatus checks if a booking status is valid
func isValidBookingStatus(status string) bool {
	validStatuses := map[string]bool{
		entity.BookingStatusPending:   true,
		entity.BookingStatusConfirmed: true,
		entity.BookingStatusCancelled: true,
		entity.BookingStatusCompleted: true,
	}
	return validStatuses[status]
}