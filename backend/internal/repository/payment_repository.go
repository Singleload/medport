package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"github.com/svenskhalsovard/api/internal/entity"
)

// PaymentRepository handles database operations for payments
type PaymentRepository struct {
	db *sqlx.DB
}

// NewPaymentRepository creates a new PaymentRepository
func NewPaymentRepository(database *Database) *PaymentRepository {
	return &PaymentRepository{
		db: database.DB,
	}
}

// CreatePayment creates a new payment with items
func (r *PaymentRepository) CreatePayment(ctx context.Context, tx *sqlx.Tx, payment *entity.Payment, items []entity.PaymentItem) error {
	// Create payment record
	if err := r.createPaymentRecord(ctx, tx, payment); err != nil {
		return fmt.Errorf("failed to create payment record: %w", err)
	}
	
	// Create payment items
	for i := range items {
		items[i].PaymentID = payment.ID
		if err := r.createPaymentItem(ctx, tx, &items[i]); err != nil {
			return fmt.Errorf("failed to create payment item: %w", err)
		}
	}
	
	return nil
}

// GetPaymentByID retrieves a payment by ID
func (r *PaymentRepository) GetPaymentByID(ctx context.Context, id int64) (*entity.Payment, error) {
	query := `
		SELECT id, external_payment_id, customer_id, amount, currency, status, 
		       payment_method, order_reference, transaction_type, error_message, 
		       created_at, updated_at, deleted_at
		FROM payments
		WHERE ` + softDeleteCondition("payments") + `
		AND id = ?
	`

	var payment entity.Payment
	if err := r.db.GetContext(ctx, &payment, query, id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // Payment not found
		}
		return nil, fmt.Errorf("failed to get payment by ID: %w", err)
	}

	return &payment, nil
}

// GetPaymentByExternalID retrieves a payment by external payment ID
func (r *PaymentRepository) GetPaymentByExternalID(ctx context.Context, externalID string) (*entity.Payment, error) {
	query := `
		SELECT id, external_payment_id, customer_id, amount, currency, status, 
		       payment_method, order_reference, transaction_type, error_message, 
		       created_at, updated_at, deleted_at
		FROM payments
		WHERE ` + softDeleteCondition("payments") + `
		AND external_payment_id = ?
	`

	var payment entity.Payment
	if err := r.db.GetContext(ctx, &payment, query, externalID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // Payment not found
		}
		return nil, fmt.Errorf("failed to get payment by external ID: %w", err)
	}

	return &payment, nil
}

// GetPaymentByOrderReference retrieves a payment by order reference
func (r *PaymentRepository) GetPaymentByOrderReference(ctx context.Context, orderReference string) (*entity.Payment, error) {
	query := `
		SELECT id, external_payment_id, customer_id, amount, currency, status, 
		       payment_method, order_reference, transaction_type, error_message, 
		       created_at, updated_at, deleted_at
		FROM payments
		WHERE ` + softDeleteCondition("payments") + `
		AND order_reference = ?
	`

	var payment entity.Payment
	if err := r.db.GetContext(ctx, &payment, query, orderReference); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // Payment not found
		}
		return nil, fmt.Errorf("failed to get payment by order reference: %w", err)
	}

	return &payment, nil
}

// GetPaymentWithItems retrieves a payment with its items
func (r *PaymentRepository) GetPaymentWithItems(ctx context.Context, id int64) (*entity.PaymentWithItems, error) {
	payment, err := r.GetPaymentByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if payment == nil {
		return nil, nil // Payment not found
	}

	// Get customer
	customerRepository := &BookingRepository{db: r.db}
	customer, err := customerRepository.GetCustomerByID(ctx, payment.CustomerID)
	if err != nil {
		return nil, fmt.Errorf("failed to get customer: %w", err)
	}

	// Get payment items
	items, err := r.GetPaymentItems(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get payment items: %w", err)
	}

	return &entity.PaymentWithItems{
		Payment:  *payment,
		Items:    items,
		Customer: *customer,
	}, nil
}

// GetPaymentItems retrieves the items for a payment
func (r *PaymentRepository) GetPaymentItems(ctx context.Context, paymentID int64) ([]entity.PaymentItem, error) {
	query := `
		SELECT id, payment_id, service_id, service_name, quantity, unit_price, 
		       total_price, purchase_type, created_at, updated_at, deleted_at
		FROM payment_items
		WHERE ` + softDeleteCondition("payment_items") + `
		AND payment_id = ?
		ORDER BY id
	`

	var items []entity.PaymentItem
	if err := r.db.SelectContext(ctx, &items, query, paymentID); err != nil {
		return nil, fmt.Errorf("failed to get payment items: %w", err)
	}

	return items, nil
}

// UpdatePaymentStatus updates the status of a payment
func (r *PaymentRepository) UpdatePaymentStatus(ctx context.Context, id int64, status string, errorMessage string) error {
	query := `
		UPDATE payments
		SET status = ?,
		    error_message = ?,
		    updated_at = ?
		WHERE id = ?
		AND ` + softDeleteCondition("payments")

	result, err := r.db.ExecContext(ctx, query, status, errorMessage, now(), id)
	if err != nil {
		return fmt.Errorf("failed to update payment status: %w", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rows == 0 {
		return fmt.Errorf("payment not found or already deleted")
	}

	return nil
}

// UpdatePaymentExternalID updates the external payment ID of a payment
func (r *PaymentRepository) UpdatePaymentExternalID(ctx context.Context, id int64, externalID string) error {
	query := `
		UPDATE payments
		SET external_payment_id = ?,
		    updated_at = ?
		WHERE id = ?
		AND ` + softDeleteCondition("payments")

	result, err := r.db.ExecContext(ctx, query, externalID, now(), id)
	if err != nil {
		return fmt.Errorf("failed to update payment external ID: %w", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rows == 0 {
		return fmt.Errorf("payment not found or already deleted")
	}

	return nil
}

// FindIncompletePayments finds payments with initiated or pending status older than the given duration
func (r *PaymentRepository) FindIncompletePayments(ctx context.Context, maxAge string) ([]entity.Payment, error) {
	query := `
		SELECT id, external_payment_id, customer_id, amount, currency, status, 
		       payment_method, order_reference, transaction_type, error_message, 
		       created_at, updated_at, deleted_at
		FROM payments
		WHERE ` + softDeleteCondition("payments") + `
		AND status IN (?, ?)
		AND created_at < DATE_SUB(NOW(), INTERVAL ` + maxAge + `)
		ORDER BY created_at
	`

	var payments []entity.Payment
	if err := r.db.SelectContext(
		ctx, 
		&payments, 
		query, 
		entity.PaymentStatusInitiated, 
		entity.PaymentStatusPending,
	); err != nil {
		return nil, fmt.Errorf("failed to find incomplete payments: %w", err)
	}

	return payments, nil
}

// Helper methods

// createPaymentRecord creates a new payment record
func (r *PaymentRepository) createPaymentRecord(ctx context.Context, tx *sqlx.Tx, payment *entity.Payment) error {
	query := `
		INSERT INTO payments (
			external_payment_id, customer_id, amount, currency, status, 
			payment_method, order_reference, transaction_type, error_message, 
			created_at, updated_at
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	now := now()
	payment.CreatedAt = now
	payment.UpdatedAt = now

	result, err := tx.ExecContext(
		ctx,
		query,
		payment.ExternalPaymentID,
		payment.CustomerID,
		payment.Amount,
		payment.Currency,
		payment.Status,
		payment.PaymentMethod,
		payment.OrderReference,
		payment.TransactionType,
		payment.ErrorMessage,
		payment.CreatedAt,
		payment.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to create payment: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("failed to get last insert ID: %w", err)
	}

	payment.ID = id
	return nil
}

// createPaymentItem creates a new payment item
func (r *PaymentRepository) createPaymentItem(ctx context.Context, tx *sqlx.Tx, item *entity.PaymentItem) error {
	query := `
		INSERT INTO payment_items (
			payment_id, service_id, service_name, quantity, unit_price, 
			total_price, purchase_type, created_at, updated_at
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	now := now()
	item.CreatedAt = now
	item.UpdatedAt = now

	result, err := tx.ExecContext(
		ctx,
		query,
		item.PaymentID,
		item.ServiceID,
		item.ServiceName,
		item.Quantity,
		item.UnitPrice,
		item.TotalPrice,
		item.PurchaseType,
		item.CreatedAt,
		item.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to create payment item: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("failed to get last insert ID: %w", err)
	}

	item.ID = id
	return nil
}