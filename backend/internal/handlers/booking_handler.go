package handlers

import (
	"errors"
	"net/http"

	"github.com/rs/zerolog/log"
	"github.com/svenskhalsovard/api/internal/dto"
	"github.com/svenskhalsovard/api/internal/entity"
)

// BookingHandler handles booking-related requests
type BookingHandler struct {
	service BookingService
}

// BookingService defines the interface for booking business logic
type BookingService interface {
	CreateBooking(ctx http.Context, paymentID int64, customer *entity.Customer) (*entity.Booking, error)
	GetBooking(ctx http.Context, id int64) (*entity.BookingWithItems, error)
}

// NewBookingHandler creates a new BookingHandler
func NewBookingHandler(service BookingService) *BookingHandler {
	return &BookingHandler{
		service: service,
	}
}

// BookingRequest represents a request to create a booking
type BookingRequest struct {
	PaymentID int64                `json:"paymentId" validate:"required,min=1"`
	Customer  dto.CustomerRequest  `json:"customer"`
}

// BookingResponse represents a booking in the API response
type BookingResponse struct {
	ID            int64                   `json:"id"`
	BookingNumber string                  `json:"bookingNumber"`
	Status        string                  `json:"status"`
	TotalAmount   float64                 `json:"totalAmount"`
	Customer      *CustomerResponse       `json:"customer,omitempty"`
	Items         []BookingItemResponse   `json:"items,omitempty"`
	CreatedAt     string                  `json:"createdAt"`
}

// CustomerResponse represents a customer in the API response
type CustomerResponse struct {
	FirstName      string `json:"firstName"`
	LastName       string `json:"lastName"`
	Email          string `json:"email"`
	Phone          string `json:"phone"`
	StreetAddress  string `json:"streetAddress"`
	PostalCode     string `json:"postalCode"`
	City           string `json:"city"`
	AdditionalInfo string `json:"additionalInfo,omitempty"`
}

// BookingItemResponse represents a booking item in the API response
type BookingItemResponse struct {
	ServiceID      int64   `json:"serviceId"`
	ServiceName    string  `json:"serviceName"`
	Quantity       int     `json:"quantity"`
	UnitPrice      float64 `json:"unitPrice"`
	TotalPrice     float64 `json:"totalPrice"`
	PurchaseType   string  `json:"purchaseType"`
	IsSubscription bool    `json:"isSubscription"`
}

// CreateBooking handles the request to create a booking
func (h *BookingHandler) CreateBooking(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req BookingRequest
	if err := ParseJSON(r, &req); err != nil {
		log.Debug().Err(err).Msg("Invalid booking request")
		RespondJSON(w, http.StatusBadRequest, dto.NewErrorResponse(
			dto.ErrorCodeInvalidRequest,
			"Invalid booking request",
			err.Error(),
		))
		return
	}

	// Map customer request to entity if provided
	var customerEntity *entity.Customer
	if (dto.CustomerRequest{}) != req.Customer {
		customer := entity.Customer{
			FirstName:      req.Customer.FirstName,
			LastName:       req.Customer.LastName,
			Email:          req.Customer.Email,
			Phone:          req.Customer.Phone,
			StreetAddress:  req.Customer.StreetAddress,
			PostalCode:     req.Customer.PostalCode,
			City:           req.Customer.City,
			AdditionalInfo: req.Customer.AdditionalInfo,
		}
		customerEntity = &customer
	}

	booking, err := h.service.CreateBooking(ctx, req.PaymentID, customerEntity)
	if err != nil {
		log.Error().Err(err).Int64("paymentID", req.PaymentID).Msg("Failed to create booking")
		
		// Handle different error types
		var statusCode int
		var errorCode string
		
		if errors.Is(err, entity.ErrPaymentNotFound) {
			statusCode = http.StatusNotFound
			errorCode = dto.ErrorCodeResourceNotFound
		} else if errors.Is(err, entity.ErrPaymentNotSuccessful) {
			statusCode = http.StatusBadRequest
			errorCode = dto.ErrorCodeInvalidRequest
		} else if errors.Is(err, entity.ErrBookingAlreadyExists) {
			statusCode = http.StatusConflict
			errorCode = dto.ErrorCodeInvalidRequest
		} else {
			statusCode = http.StatusInternalServerError
			errorCode = dto.ErrorCodeInternalServerError
		}
		
		RespondJSON(w, statusCode, dto.NewErrorResponse(
			errorCode,
			err.Error(),
			nil,
		))
		return
	}

	// Get the full booking with items to include in the response
	bookingWithItems, err := h.service.GetBooking(ctx, booking.ID)
	if err != nil {
		log.Error().Err(err).Int64("bookingID", booking.ID).Msg("Failed to get booking details")
		// Even if we can't get full booking details, we'll return a basic response
		response := BookingResponse{
			ID:            booking.ID,
			BookingNumber: booking.BookingNumber,
			Status:        booking.Status,
			TotalAmount:   booking.TotalAmount,
			CreatedAt:     booking.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		}
		RespondJSON(w, http.StatusCreated, dto.NewSuccessResponse(response))
		return
	}

	// Map booking to response
	response := mapBookingToResponse(*bookingWithItems)
	RespondJSON(w, http.StatusCreated, dto.NewSuccessResponse(response))
}

// mapBookingToResponse maps a BookingWithItems to a BookingResponse
func mapBookingToResponse(booking entity.BookingWithItems) BookingResponse {
	response := BookingResponse{
		ID:            booking.Booking.ID,
		BookingNumber: booking.Booking.BookingNumber,
		Status:        booking.Booking.Status,
		TotalAmount:   booking.Booking.TotalAmount,
		CreatedAt:     booking.Booking.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		Customer: &CustomerResponse{
			FirstName:      booking.Customer.FirstName,
			LastName:       booking.Customer.LastName,
			Email:          booking.Customer.Email,
			Phone:          booking.Customer.Phone,
			StreetAddress:  booking.Customer.StreetAddress,
			PostalCode:     booking.Customer.PostalCode,
			City:           booking.Customer.City,
			AdditionalInfo: booking.Customer.AdditionalInfo,
		},
		Items: make([]BookingItemResponse, 0, len(booking.Items)),
	}

	for _, item := range booking.Items {
		response.Items = append(response.Items, BookingItemResponse{
			ServiceID:      item.ServiceID,
			ServiceName:    item.ServiceName,
			Quantity:       item.Quantity,
			UnitPrice:      item.UnitPrice,
			TotalPrice:     item.TotalPrice,
			PurchaseType:   item.PurchaseType,
			IsSubscription: item.IsSubscription,
		})
	}

	return response
}