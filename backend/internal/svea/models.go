package svea

import "time"

// OrderRequest represents a request to create an order in Svea Ekonomi
type OrderRequest struct {
	OrderReference    string      `json:"orderReference"`
	Currency          string      `json:"currency"`
	ClientOrderNumber string      `json:"clientOrderNumber"`
	PaymentType       string      `json:"paymentType"`
	Customer          Customer    `json:"customer"`
	Items             []OrderItem `json:"items"`
}

// Customer represents a customer in a Svea order
type Customer struct {
	FirstName     string  `json:"firstName"`
	LastName      string  `json:"lastName"`
	Email         string  `json:"email"`
	PhoneNumber   string  `json:"phoneNumber"`
	Address       Address `json:"address"`
	CompanyName   string  `json:"companyName,omitempty"`
	OrgNumber     string  `json:"orgNumber,omitempty"`
	CountryCode   string  `json:"countryCode,omitempty"`
	NationalId    string  `json:"nationalId,omitempty"`
	IsBusinessB2B bool    `json:"isBusinessB2B,omitempty"`
}

// Address represents a customer address in a Svea order
type Address struct {
	StreetAddress string `json:"streetAddress"`
	PostalCode    string `json:"postalCode"`
	City          string `json:"city"`
	CountryCode   string `json:"countryCode"`
}

// OrderItem represents an item in a Svea order
type OrderItem struct {
	ArticleNumber string `json:"articleNumber"`
	Name          string `json:"name"`
	Quantity      int    `json:"quantity"`
	UnitPrice     int    `json:"unitPrice"` // In cents/öre
	VatPercent    int    `json:"vatPercent"`
	Unit          string `json:"unit"`
	Discount      int    `json:"discount,omitempty"`
}

// OrderResponse represents a response from creating an order in Svea Ekonomi
type OrderResponse struct {
	OrderID      string         `json:"orderId"`
	Status       string         `json:"status"`
	CheckoutURL  string         `json:"checkoutUrl"`
	CheckoutUI   CheckoutUIData `json:"checkoutUI"`
	ClientID     string         `json:"clientId"`
	ClientSecret string         `json:"clientSecret"`
}

// CheckoutUIData represents the UI data for Svea checkout
type CheckoutUIData struct {
	HTML      string `json:"html"`
	JavaScript string `json:"javascript"`
	Snippet   string `json:"snippet"`
}

// Order represents an order in Svea Ekonomi
type Order struct {
	ID               string      `json:"id"`
	OrderReference   string      `json:"orderReference"`
	Status           string      `json:"status"`
	Currency         string      `json:"currency"`
	ClientOrderNumber string     `json:"clientOrderNumber"`
	PaymentType      string      `json:"paymentType"`
	Customer         Customer    `json:"customer"`
	Items            []OrderItem `json:"items"`
	CreatedAt        time.Time   `json:"createdAt"`
	UpdatedAt        time.Time   `json:"updatedAt"`
	Payments         []Payment   `json:"payments"`
}

// Payment represents a payment in Svea Ekonomi
type Payment struct {
	ID             string    `json:"id"`
	PaymentMethod  string    `json:"paymentMethod"`
	Status         string    `json:"status"`
	Amount         int       `json:"amount"`
	Currency       string    `json:"currency"`
	TransactionID  string    `json:"transactionId"`
	PaymentDetails interface{} `json:"paymentDetails"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}

// PaymentRequest represents a request to create a payment in Svea Ekonomi
type PaymentRequest struct {
	PaymentMethod string `json:"paymentMethod"`
}

// PaymentResponse represents a response from creating a payment in Svea Ekonomi
type PaymentResponse struct {
	ID            string    `json:"id"`
	OrderID       string    `json:"orderId"`
	PaymentMethod string    `json:"paymentMethod"`
	Status        string    `json:"status"`
	Amount        int       `json:"amount"`
	Currency      string    `json:"currency"`
	TransactionID string    `json:"transactionId"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}

// RefundRequest represents a request to refund a payment in Svea Ekonomi
type RefundRequest struct {
	Amount  int    `json:"amount"`
	Comment string `json:"comment,omitempty"`
}

// RefundResponse represents a response from refunding a payment in Svea Ekonomi
type RefundResponse struct {
	ID            string    `json:"id"`
	OrderID       string    `json:"orderId"`
	PaymentID     string    `json:"paymentId"`
	Amount        int       `json:"amount"`
	Currency      string    `json:"currency"`
	Status        string    `json:"status"`
	TransactionID string    `json:"transactionId"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}

// OrderStatus constants
const (
	OrderStatusCreated   = "Created"
	OrderStatusPending   = "Pending"
	OrderStatusCompleted = "Completed"
	OrderStatusCancelled = "Cancelled"
	OrderStatusFailed    = "Failed"
)

// PaymentMethodType constants
const (
	PaymentMethodCard      = "card"
	PaymentMethodInvoice   = "invoice"
	PaymentMethodDirectDebit = "direct_debit"
	PaymentMethodSwish     = "swish"
)

// PaymentType constants
const (
	PaymentTypeOneTime    = "onetime"
	PaymentTypeRecurring  = "recurring"
)