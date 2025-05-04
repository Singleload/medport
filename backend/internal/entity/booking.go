package entity

import "time"

// Booking represents a customer booking for a health service
type Booking struct {
	ID            int64      `db:"id" json:"id"`
	PaymentID     int64      `db:"payment_id" json:"paymentId"`
	CustomerID    int64      `db:"customer_id" json:"customerId"`
	Status        string     `db:"status" json:"status"`
	TotalAmount   float64    `db:"total_amount" json:"totalAmount"`
	BookingNumber string     `db:"booking_number" json:"bookingNumber"`
	Notes         string     `db:"notes" json:"notes,omitempty"`
	CreatedAt     time.Time  `db:"created_at" json:"createdAt"`
	UpdatedAt     time.Time  `db:"updated_at" json:"updatedAt"`
	DeletedAt     *time.Time `db:"deleted_at" json:"deletedAt,omitempty"`
}

// BookingItem represents an item in a booking
type BookingItem struct {
	ID             int64      `db:"id" json:"id"`
	BookingID      int64      `db:"booking_id" json:"bookingId"`
	ServiceID      int64      `db:"service_id" json:"serviceId"`
	ServiceName    string     `db:"service_name" json:"serviceName"`
	Quantity       int        `db:"quantity" json:"quantity"`
	UnitPrice      float64    `db:"unit_price" json:"unitPrice"`
	TotalPrice     float64    `db:"total_price" json:"totalPrice"`
	PurchaseType   string     `db:"purchase_type" json:"purchaseType"`
	IsSubscription bool       `db:"is_subscription" json:"isSubscription"`
	CreatedAt      time.Time  `db:"created_at" json:"createdAt"`
	UpdatedAt      time.Time  `db:"updated_at" json:"updatedAt"`
	DeletedAt      *time.Time `db:"deleted_at" json:"deletedAt,omitempty"`
}

// Customer represents a customer who makes bookings
type Customer struct {
	ID             int64      `db:"id" json:"id"`
	FirstName      string     `db:"first_name" json:"firstName"`
	LastName       string     `db:"last_name" json:"lastName"`
	Email          string     `db:"email" json:"email"`
	Phone          string     `db:"phone" json:"phone"`
	StreetAddress  string     `db:"street_address" json:"streetAddress"`
	PostalCode     string     `db:"postal_code" json:"postalCode"`
	City           string     `db:"city" json:"city"`
	AdditionalInfo string     `db:"additional_info" json:"additionalInfo,omitempty"`
	CreatedAt      time.Time  `db:"created_at" json:"createdAt"`
	UpdatedAt      time.Time  `db:"updated_at" json:"updatedAt"`
	DeletedAt      *time.Time `db:"deleted_at" json:"deletedAt,omitempty"`
}

// BookingStatus represents the possible status values for a booking
const (
	BookingStatusPending   = "pending"
	BookingStatusConfirmed = "confirmed"
	BookingStatusCancelled = "cancelled"
	BookingStatusCompleted = "completed"
)

// BookingWithItems represents a booking with its items
type BookingWithItems struct {
	Booking  Booking       `json:"booking"`
	Customer Customer      `json:"customer"`
	Items    []BookingItem `json:"items"`
}