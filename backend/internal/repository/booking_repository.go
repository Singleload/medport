package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"github.com/svenskhalsovard/api/internal/entity"
)

// BookingRepository handles database operations for bookings
type BookingRepository struct {
	db *sqlx.DB
}

// NewBookingRepository creates a new BookingRepository
func NewBookingRepository(database *Database) *BookingRepository {
	return &BookingRepository{
		db: database.DB,
	}
}

// CreateBooking creates a new booking with items and customer data
func (r *BookingRepository) CreateBooking(ctx context.Context, tx *sqlx.Tx, booking *entity.Booking, customer *entity.Customer, items []entity.BookingItem) error {
	var err error
	
	// Create or update customer
	var customerID int64
	if customer.ID > 0 {
		// Update existing customer
		if err = r.updateCustomer(ctx, tx, customer); err != nil {
			return fmt.Errorf("failed to update customer: %w", err)
		}
		customerID = customer.ID
	} else {
		// Check if customer exists with the same email
		existingCustomer, err := r.getCustomerByEmail(ctx, tx, customer.Email)
		if err != nil {
			return fmt.Errorf("failed to check for existing customer: %w", err)
		}
		
		if existingCustomer != nil {
			// Update existing customer data
			customer.ID = existingCustomer.ID
			if err = r.updateCustomer(ctx, tx, customer); err != nil {
				return fmt.Errorf("failed to update existing customer: %w", err)
			}
			customerID = customer.ID
		} else {
			// Create new customer
			if err = r.createCustomer(ctx, tx, customer); err != nil {
				return fmt.Errorf("failed to create customer: %w", err)
			}
			customerID = customer.ID
		}
	}
	
	// Set customer ID in booking
	booking.CustomerID = customerID
	
	// Generate booking number
	bookingNumber, err := r.generateBookingNumber(ctx, tx)
	if err != nil {
		return fmt.Errorf("failed to generate booking number: %w", err)
	}
	booking.BookingNumber = bookingNumber
	
	// Create booking
	if err = r.createBookingRecord(ctx, tx, booking); err != nil {
		return fmt.Errorf("failed to create booking record: %w", err)
	}
	
	// Create booking items
	for i := range items {
		items[i].BookingID = booking.ID
		if err = r.createBookingItem(ctx, tx, &items[i]); err != nil {
			return fmt.Errorf("failed to create booking item: %w", err)
		}
	}
	
	return nil
}

// GetBookingByID retrieves a booking by ID
func (r *BookingRepository) GetBookingByID(ctx context.Context, id int64) (*entity.Booking, error) {
	query := `
		SELECT id, payment_id, customer_id, status, total_amount, booking_number, 
		       notes, created_at, updated_at, deleted_at
		FROM bookings
		WHERE ` + softDeleteCondition("bookings") + `
		AND id = ?
	`

	var booking entity.Booking
	if err := r.db.GetContext(ctx, &booking, query, id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // Booking not found
		}
		return nil, fmt.Errorf("failed to get booking by ID: %w", err)
	}

	return &booking, nil
}

// GetBookingByPaymentID retrieves a booking by payment ID
func (r *BookingRepository) GetBookingByPaymentID(ctx context.Context, paymentID int64) (*entity.Booking, error) {
	query := `
		SELECT id, payment_id, customer_id, status, total_amount, booking_number, 
		       notes, created_at, updated_at, deleted_at
		FROM bookings
		WHERE ` + softDeleteCondition("bookings") + `
		AND payment_id = ?
	`

	var booking entity.Booking
	if err := r.db.GetContext(ctx, &booking, query, paymentID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // Booking not found
		}
		return nil, fmt.Errorf("failed to get booking by payment ID: %w", err)
	}

	return &booking, nil
}

// GetBookingWithItems retrieves a booking with its items and customer info
func (r *BookingRepository) GetBookingWithItems(ctx context.Context, id int64) (*entity.BookingWithItems, error) {
	booking, err := r.GetBookingByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if booking == nil {
		return nil, nil // Booking not found
	}

	// Get customer
	customer, err := r.GetCustomerByID(ctx, booking.CustomerID)
	if err != nil {
		return nil, fmt.Errorf("failed to get customer: %w", err)
	}

	// Get booking items
	items, err := r.GetBookingItems(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get booking items: %w", err)
	}

	return &entity.BookingWithItems{
		Booking:  *booking,
		Customer: *customer,
		Items:    items,
	}, nil
}

// GetBookingItems retrieves the items for a booking
func (r *BookingRepository) GetBookingItems(ctx context.Context, bookingID int64) ([]entity.BookingItem, error) {
	query := `
		SELECT id, booking_id, service_id, service_name, quantity, unit_price, 
		       total_price, purchase_type, is_subscription, created_at, updated_at, deleted_at
		FROM booking_items
		WHERE ` + softDeleteCondition("booking_items") + `
		AND booking_id = ?
		ORDER BY id
	`

	var items []entity.BookingItem
	if err := r.db.SelectContext(ctx, &items, query, bookingID); err != nil {
		return nil, fmt.Errorf("failed to get booking items: %w", err)
	}

	return items, nil
}

// GetCustomerByID retrieves a customer by ID
func (r *BookingRepository) GetCustomerByID(ctx context.Context, id int64) (*entity.Customer, error) {
	query := `
		SELECT id, first_name, last_name, email, phone, street_address, 
		       postal_code, city, additional_info, created_at, updated_at, deleted_at
		FROM customers
		WHERE ` + softDeleteCondition("customers") + `
		AND id = ?
	`

	var customer entity.Customer
	if err := r.db.GetContext(ctx, &customer, query, id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // Customer not found
		}
		return nil, fmt.Errorf("failed to get customer by ID: %w", err)
	}

	return &customer, nil
}

// UpdateBookingStatus updates the status of a booking
func (r *BookingRepository) UpdateBookingStatus(ctx context.Context, bookingID int64, status string) error {
	query := `
		UPDATE bookings
		SET status = ?,
		    updated_at = ?
		WHERE id = ?
		AND ` + softDeleteCondition("bookings")

	result, err := r.db.ExecContext(ctx, query, status, now(), bookingID)
	if err != nil {
		return fmt.Errorf("failed to update booking status: %w", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rows == 0 {
		return fmt.Errorf("booking not found or already deleted")
	}

	return nil
}

// Helper methods

// createBookingRecord creates a new booking record
func (r *BookingRepository) createBookingRecord(ctx context.Context, tx *sqlx.Tx, booking *entity.Booking) error {
	query := `
		INSERT INTO bookings (
			payment_id, customer_id, status, total_amount, booking_number, 
			notes, created_at, updated_at
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?)
	`

	now := now()
	booking.CreatedAt = now
	booking.UpdatedAt = now

	result, err := tx.ExecContext(
		ctx,
		query,
		booking.PaymentID,
		booking.CustomerID,
		booking.Status,
		booking.TotalAmount,
		booking.BookingNumber,
		booking.Notes,
		booking.CreatedAt,
		booking.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to create booking: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("failed to get last insert ID: %w", err)
	}

	booking.ID = id
	return nil
}

// createBookingItem creates a new booking item
func (r *BookingRepository) createBookingItem(ctx context.Context, tx *sqlx.Tx, item *entity.BookingItem) error {
	query := `
		INSERT INTO booking_items (
			booking_id, service_id, service_name, quantity, unit_price, 
			total_price, purchase_type, is_subscription, created_at, updated_at
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	now := now()
	item.CreatedAt = now
	item.UpdatedAt = now

	result, err := tx.ExecContext(
		ctx,
		query,
		item.BookingID,
		item.ServiceID,
		item.ServiceName,
		item.Quantity,
		item.UnitPrice,
		item.TotalPrice,
		item.PurchaseType,
		item.IsSubscription,
		item.CreatedAt,
		item.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to create booking item: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("failed to get last insert ID: %w", err)
	}

	item.ID = id
	return nil
}

// createCustomer creates a new customer
func (r *BookingRepository) createCustomer(ctx context.Context, tx *sqlx.Tx, customer *entity.Customer) error {
	query := `
		INSERT INTO customers (
			first_name, last_name, email, phone, street_address, 
			postal_code, city, additional_info, created_at, updated_at
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	now := now()
	customer.CreatedAt = now
	customer.UpdatedAt = now

	result, err := tx.ExecContext(
		ctx,
		query,
		customer.FirstName,
		customer.LastName,
		customer.Email,
		customer.Phone,
		customer.StreetAddress,
		customer.PostalCode,
		customer.City,
		customer.AdditionalInfo,
		customer.CreatedAt,
		customer.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to create customer: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("failed to get last insert ID: %w", err)
	}

	customer.ID = id
	return nil
}

// updateCustomer updates an existing customer
func (r *BookingRepository) updateCustomer(ctx context.Context, tx *sqlx.Tx, customer *entity.Customer) error {
	query := `
		UPDATE customers
		SET first_name = ?,
			last_name = ?,
			phone = ?,
			street_address = ?,
			postal_code = ?,
			city = ?,
			additional_info = ?,
			updated_at = ?
		WHERE id = ?
		AND ` + softDeleteCondition("customers")
	
	now := now()
	customer.UpdatedAt = now

	result, err := tx.ExecContext(
		ctx,
		query,
		customer.FirstName,
		customer.LastName,
		customer.Phone,
		customer.StreetAddress,
		customer.PostalCode,
		customer.City,
		customer.AdditionalInfo,
		customer.UpdatedAt,
		customer.ID,
	)
	if err != nil {
		return fmt.Errorf("failed to update customer: %w", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rows == 0 {
		return fmt.Errorf("customer not found or already deleted")
	}

	return nil
}

// getCustomerByEmail retrieves a customer by email
func (r *BookingRepository) getCustomerByEmail(ctx context.Context, tx *sqlx.Tx, email string) (*entity.Customer, error) {
	query := `
		SELECT id, first_name, last_name, email, phone, street_address, 
		       postal_code, city, additional_info, created_at, updated_at, deleted_at
		FROM customers
		WHERE ` + softDeleteCondition("customers") + `
		AND email = ?
	`

	var customer entity.Customer
	if err := tx.GetContext(ctx, &customer, query, email); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // Customer not found
		}
		return nil, fmt.Errorf("failed to get customer by email: %w", err)
	}

	return &customer, nil
}

// generateBookingNumber generates a unique booking number
func (r *BookingRepository) generateBookingNumber(ctx context.Context, tx *sqlx.Tx) (string, error) {
	// Get current year and latest booking sequence number
	var seq int
	query := `
		SELECT COALESCE(MAX(CAST(SUBSTRING_INDEX(booking_number, '-', -1) AS UNSIGNED)), 0) AS seq
		FROM bookings
		WHERE booking_number LIKE CONCAT(?, '%')
	`

	year := now().Format("2006")
	if err := tx.GetContext(ctx, &seq, query, year); err != nil {
		return "", fmt.Errorf("failed to get latest booking sequence: %w", err)
	}

	// Increment sequence and format booking number
	seq++
	bookingNumber := fmt.Sprintf("%s-%06d", year, seq)
	return bookingNumber, nil
}