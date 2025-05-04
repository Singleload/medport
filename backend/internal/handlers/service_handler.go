package handlers

import (
	"errors"
	"net/http"

	"github.com/rs/zerolog/log"
	"github.com/svenskhalsovard/api/internal/dto"
	"github.com/svenskhalsovard/api/internal/entity"
)

// ServiceHandler handles service-related requests
type ServiceHandler struct {
	service ServiceService
}

// ServiceService defines the interface for service business logic
type ServiceService interface {
	GetServices(ctx http.Context) ([]entity.Service, error)
	GetServiceByID(ctx http.Context, id int64) (*entity.Service, error)
	GetServiceWithFeatures(ctx http.Context, id int64) (*entity.ServiceWithFeatures, error)
}

// NewServiceHandler creates a new ServiceHandler
func NewServiceHandler(service ServiceService) *ServiceHandler {
	return &ServiceHandler{
		service: service,
	}
}

// GetServices handles the request to get all services
func (h *ServiceHandler) GetServices(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	services, err := h.service.GetServices(ctx)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get services")
		RespondError(w, err)
		return
	}

	// Simple implementation without features
	response := make([]dto.ServiceResponse, 0, len(services))
	for _, service := range services {
		response = append(response, dto.MapServiceToResponse(service, nil))
	}

	RespondJSON(w, http.StatusOK, dto.NewSuccessResponse(dto.ServiceListResponse{
		Services: response,
		Count:    len(response),
	}))
}

// GetServiceByID handles the request to get a service by ID
func (h *ServiceHandler) GetServiceByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id, err := ParseIDParam(r, "id")
	if err != nil {
		RespondJSON(w, http.StatusBadRequest, dto.NewErrorResponse(
			dto.ErrorCodeInvalidRequest,
			err.Error(),
			nil,
		))
		return
	}

	serviceWithFeatures, err := h.service.GetServiceWithFeatures(ctx, id)
	if err != nil {
		log.Error().Err(err).Int64("serviceID", id).Msg("Failed to get service")
		RespondError(w, err)
		return
	}

	if serviceWithFeatures == nil {
		RespondJSON(w, http.StatusNotFound, dto.NewErrorResponse(
			dto.ErrorCodeResourceNotFound,
			"Service not found",
			nil,
		))
		return
	}

	response := dto.MapServiceWithFeaturesToResponse(*serviceWithFeatures)
	RespondJSON(w, http.StatusOK, dto.NewSuccessResponse(response))
}