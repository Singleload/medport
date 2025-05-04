package entity

import "time"

// Payment represents a payment for a booking
type Payment struct {
	ID                int64      `db:"id" json:"id"`
	ExternalPaymentID string     `db:"external_payment_id" json:"externalPaymentId"`
	CustomerID        int64      `db:"customer_id" json:"customerId"`
	Amount            float64    `db:"amount" json:"amount"`
	Currency          string     `db:"currency" json:"currency"`
	Status            string     `db:"status" json:"status"`
	PaymentMethod     string     `db:"payment_method" json:"paymentMethod"`
	OrderReference    string     `db:"order_reference" json:"orderReference"`
	TransactionType   string     `db:"transaction_type" json:"transactionType"`
	ErrorMessage      string     `db:"error_message" json:"errorMessage,omitempty"`
	CreatedAt         time.Time  `db:"created_at" json:"createdAt"`
	UpdatedAt         time.Time  `db:"updated_at" json:"updatedAt"`
	DeletedAt         *time.Time `db:"deleted_at" json:"deletedAt,omitempty"`
}

// PaymentStatus represents the possible status values for a payment
const (
	PaymentStatusInitiated = "initiated"
	PaymentStatusPending   = "pending"
	PaymentStatusSuccess   = "success"
	PaymentStatusFailed    = "failed"
	PaymentStatusCancelled = "cancelled"
	PaymentStatusRefunded  = "refunded"
)

// TransactionType represents the type of transaction
const (
	TransactionTypeOneTime     = "one-time"
	TransactionTypeSubscription = "subscription"
)

// PaymentMethod represents the payment methods supported
const (
	PaymentMethodCard      = "card"
	PaymentMethodInvoice   = "invoice"
	PaymentMethodDirectDebit = "direct_debit"
	PaymentMethodSwish     = "swish"
)

// Currency represents the currencies supported
const (
	CurrencySEK = "SEK"
	CurrencyEUR = "EUR"
)

// PaymentItem represents an item in a payment
type PaymentItem struct {
	ID              int64      `db:"id" json:"id"`
	PaymentID       int64      `db:"payment_id" json:"paymentId"`
	ServiceID       int64      `db:"service_id" json:"serviceId"`
	ServiceName     string     `db:"service_name" json:"serviceName"`
	Quantity        int        `db:"quantity" json:"quantity"`
	UnitPrice       float64    `db:"unit_price" json:"unitPrice"`
	TotalPrice      float64    `db:"total_price" json:"totalPrice"`
	PurchaseType    string     `db:"purchase_type" json:"purchaseType"`
	CreatedAt       time.Time  `db:"created_at" json:"createdAt"`
	UpdatedAt       time.Time  `db:"updated_at" json:"updatedAt"`
	DeletedAt       *time.Time `db:"deleted_at" json:"deletedAt,omitempty"`
}

// PaymentWithItems represents a payment with its items
type PaymentWithItems struct {
	Payment Payment       `json:"payment"`
	Items   []PaymentItem `json:"items"`
	Customer Customer     `json:"customer"`
}