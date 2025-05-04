package handlers

import (
	"errors"
	"net/http"

	"github.com/rs/zerolog/log"
	"github.com/svenskhalsovard/api/internal/dto"
	"github.com/svenskhalsovard/api/internal/entity"
	"github.com/svenskhalsovard/api/internal/service"
)

// CheckoutHandler handles checkout-related requests
type CheckoutHandler struct {
	service CheckoutService
}

// CheckoutService defines the interface for checkout business logic
type CheckoutService interface {
	InitiateCheckout(ctx http.Context, req *service.CheckoutRequest) (*service.CheckoutResult, error)
	ProcessPayment(ctx http.Context, paymentID int64, paymentMethod string) error
	VerifyPayment(ctx http.Context, paymentID int64) (*entity.Payment, error)
}

// NewCheckoutHandler creates a new CheckoutHandler
func NewCheckoutHandler(service CheckoutService) *CheckoutHandler {
	return &CheckoutHandler{
		service: service,
	}
}

// InitiateCheckout handles the request to initiate a checkout
func (h *CheckoutHandler) InitiateCheckout(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req dto.CheckoutRequest
	if err := ParseJSON(r, &req); err != nil {
		log.Debug().Err(err).Msg("Invalid checkout request")
		RespondJSON(w, http.StatusBadRequest, dto.NewErrorResponse(
			dto.ErrorCodeInvalidRequest,
			"Invalid checkout request",
			err.Error(),
		))
		return
	}

	// Map DTO to service model
	serviceReq := &service.CheckoutRequest{
		Customer: entity.Customer{
			FirstName:      req.Customer.FirstName,
			LastName:       req.Customer.LastName,
			Email:          req.Customer.Email,
			Phone:          req.Customer.Phone,
			StreetAddress:  req.Customer.StreetAddress,
			PostalCode:     req.Customer.PostalCode,
			City:           req.Customer.City,
			AdditionalInfo: req.Customer.AdditionalInfo,
		},
		Items:       make([]service.CheckoutItem, 0, len(req.Items)),
		TotalAmount: req.TotalAmount,
	}

	for _, item := range req.Items {
		serviceReq.Items = append(serviceReq.Items, service.CheckoutItem{
			ServiceID:    item.ServiceID,
			Quantity:     item.Quantity,
			PurchaseType: item.PurchaseType,
			Price:        item.Price,
		})
	}

	result, err := h.service.InitiateCheckout(ctx, serviceReq)
	if err != nil {
		log.Error().Err(err).Msg("Failed to initiate checkout")
		
		// Handle business logic errors
		var statusCode int
		var errorCode string
		
		if errors.Is(err, service.ErrInvalidServices) || 
		   errors.Is(err, service.ErrInvalidPurchaseType) || 
		   errors.Is(err, service.ErrInvalidPrice) {
			statusCode = http.StatusBadRequest
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

	// Map service model to DTO
	response := dto.CheckoutResponse{
		PaymentID:      result.PaymentID,
		SveaOrderID:    result.SveaOrderID,
		Status:         result.Status,
		OrderReference: result.PaymentID,
		CheckoutUI: dto.CheckoutUIResponse{
			HTML:      result.SveaCheckoutUI.HTML,
			JavaScript: result.SveaCheckoutUI.JavaScript,
			Snippet:   result.SveaCheckoutUI.Snippet,
		},
	}

	RespondJSON(w, http.StatusCreated, dto.NewSuccessResponse(response))
}

// ProcessPayment handles the request to process a payment
func (h *CheckoutHandler) ProcessPayment(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req dto.PaymentRequest
	if err := ParseJSON(r, &req); err != nil {
		log.Debug().Err(err).Msg("Invalid payment request")
		RespondJSON(w, http.StatusBadRequest, dto.NewErrorResponse(
			dto.ErrorCodeInvalidRequest,
			"Invalid payment request",
			err.Error(),
		))
		return
	}

	err := h.service.ProcessPayment(ctx, req.PaymentID, req.PaymentMethod)
	if err != nil {
		log.Error().Err(err).Int64("paymentID", req.PaymentID).Str("method", req.PaymentMethod).Msg("Failed to process payment")
		
		// Handle business logic errors
		var statusCode int
		var errorCode string
		
		if errors.Is(err, service.ErrPaymentNotFound) || 
		   errors.Is(err, service.ErrInvalidPaymentStatus) {
			statusCode = http.StatusBadRequest
			errorCode = dto.ErrorCodeInvalidRequest
		} else if errors.Is(err, service.ErrPaymentFailed) {
			statusCode = http.StatusBadRequest
			errorCode = dto.ErrorCodePaymentFailed
		} else if errors.Is(err, service.ErrBookingFailed) {
			statusCode = http.StatusInternalServerError
			errorCode = dto.ErrorCodeBookingFailed
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

	RespondJSON(w, http.StatusOK, dto.NewSuccessResponse(map[string]string{
		"status": "success",
	}))
}

// VerifyPayment handles the request to verify a payment
func (h *CheckoutHandler) VerifyPayment(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	paymentID, err := ParseIDParam(r, "paymentId")
	if err != nil {
		RespondJSON(w, http.StatusBadRequest, dto.NewErrorResponse(
			dto.ErrorCodeInvalidRequest,
			err.Error(),
			nil,
		))
		return
	}

	payment, err := h.service.VerifyPayment(ctx, paymentID)
	if err != nil {
		log.Error().Err(err).Int64("paymentID", paymentID).Msg("Failed to verify payment")
		
		var statusCode int
		var errorCode string
		
		if errors.Is(err, service.ErrPaymentNotFound) {
			statusCode = http.StatusNotFound
			errorCode = dto.ErrorCodeResourceNotFound
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

	response := dto.MapPaymentToResponse(*payment)
	RespondJSON(w, http.StatusOK, dto.NewSuccessResponse(response))
}