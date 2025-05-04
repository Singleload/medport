package dto

import (
	"github.com/svenskhalsovard/api/internal/entity"
	"github.com/svenskhalsovard/api/internal/svea"
)

// CustomerRequest represents a customer in the API request
type CustomerRequest struct {
	FirstName      string `json:"firstName" validate:"required"`
	LastName       string `json:"lastName" validate:"required"`
	Email          string `json:"email" validate:"required,email"`
	Phone          string `json:"phone" validate:"required"`
	StreetAddress  string `json:"streetAddress" validate:"required"`
	PostalCode     string `json:"postalCode" validate:"required"`
	City           string `json:"city" validate:"required"`
	AdditionalInfo string `json:"additionalInfo"`
}

// CheckoutItemRequest represents an item in a checkout request
type CheckoutItemRequest struct {
	ServiceID    int64   `json:"serviceId" validate:"required,min=1"`
	Quantity     int     `json:"quantity" validate:"required,min=1,max=10"`
	PurchaseType string  `json:"purchaseType" validate:"required,oneof=one-time subscription"`
	Price        float64 `json:"price" validate:"required,min=0"`
}

// CheckoutRequest represents a checkout request
type CheckoutRequest struct {
	Customer    CustomerRequest      `json:"customer" validate:"required"`
	Items       []CheckoutItemRequest `json:"items" validate:"required,min=1,dive"`
	TotalAmount float64              `json:"totalAmount" validate:"required,min=0"`
}

// PaymentRequest represents a payment request
type PaymentRequest struct {
	PaymentID     int64  `json:"paymentId" validate:"required,min=1"`
	PaymentMethod string `json:"paymentMethod" validate:"required,oneof=card invoice direct_debit swish"`
}

// PaymentVerifyRequest represents a request to verify a payment
type PaymentVerifyRequest struct {
	PaymentID int64 `json:"paymentId" validate:"required,min=1"`
}

// CheckoutResponse represents a checkout response
type CheckoutResponse struct {
	PaymentID      int64              `json:"paymentId"`
	SveaOrderID    string             `json:"sveaOrderId"`
	CheckoutUI     CheckoutUIResponse `json:"checkoutUI"`
	Status         string             `json:"status"`
	OrderReference string             `json:"orderReference"`
}

// CheckoutUIResponse represents the UI data for checkout
type CheckoutUIResponse struct {
	HTML      string `json:"html"`
	JavaScript string `json:"javascript"`
	Snippet   string `json:"snippet"`
}

// PaymentResponse represents a payment response
type PaymentResponse struct {
	ID             int64   `json:"id"`
	Status         string  `json:"status"`
	Amount         float64 `json:"amount"`
	Currency       string  `json:"currency"`
	PaymentMethod  string  `json:"paymentMethod"`
	OrderReference string  `json:"orderReference"`
	CreatedAt      string  `json:"createdAt"`
}

// MapCustomerRequestToEntity maps a CustomerRequest to an entity.Customer
func MapCustomerRequestToEntity(req CustomerRequest) entity.Customer {
	return entity.Customer{
		FirstName:      req.FirstName,
		LastName:       req.LastName,
		Email:          req.Email,
		Phone:          req.Phone,
		StreetAddress:  req.StreetAddress,
		PostalCode:     req.PostalCode,
		City:           req.City,
		AdditionalInfo: req.AdditionalInfo,
	}
}

// MapCheckoutRequestToServiceItems converts a checkout request to service items for validation
func MapCheckoutRequestToServiceItems(req CheckoutRequest) []struct {
	ServiceID int64
	Quantity  int
} {
	items := make([]struct {
		ServiceID int64
		Quantity  int
	}, 0, len(req.Items))

	for _, item := range req.Items {
		items = append(items, struct {
			ServiceID int64
			Quantity  int
		}{
			ServiceID: item.ServiceID,
			Quantity:  item.Quantity,
		})
	}

	return items
}

// MapSveaCheckoutToResponse maps a Svea checkout response to a CheckoutResponse
func MapSveaCheckoutToResponse(paymentID int64, orderRef string, sveaResponse svea.OrderResponse) CheckoutResponse {
	return CheckoutResponse{
		PaymentID:      paymentID,
		SveaOrderID:    sveaResponse.OrderID,
		Status:         sveaResponse.Status,
		OrderReference: orderRef,
		CheckoutUI: CheckoutUIResponse{
			HTML:      sveaResponse.CheckoutUI.HTML,
			JavaScript: sveaResponse.CheckoutUI.JavaScript,
			Snippet:   sveaResponse.CheckoutUI.Snippet,
		},
	}
}

// MapPaymentToResponse maps an entity.Payment to a PaymentResponse
func MapPaymentToResponse(payment entity.Payment) PaymentResponse {
	return PaymentResponse{
		ID:             payment.ID,
		Status:         payment.Status,
		Amount:         payment.Amount,
		Currency:       payment.Currency,
		PaymentMethod:  payment.PaymentMethod,
		OrderReference: payment.OrderReference,
		CreatedAt:      payment.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}
}