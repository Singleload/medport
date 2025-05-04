package service

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"
	"github.com/svenskhalsovard/api/internal/entity"
)

// ServiceRepository defines the interface for service data operations
type ServiceRepository interface {
	GetServices(ctx context.Context) ([]entity.Service, error)
	GetServiceByID(ctx context.Context, id int64) (*entity.Service, error)
	GetServiceWithFeatures(ctx context.Context, id int64) (*entity.ServiceWithFeatures, error)
	GetServicesByIDs(ctx context.Context, ids []int64) (map[int64]entity.Service, error)
}

// ServiceService provides business logic for services
type ServiceService struct {
	repo ServiceRepository
}

// NewServiceService creates a new ServiceService
func NewServiceService(repo ServiceRepository) *ServiceService {
	return &ServiceService{
		repo: repo,
	}
}

// GetServices retrieves all active services
func (s *ServiceService) GetServices(ctx context.Context) ([]entity.Service, error) {
	services, err := s.repo.GetServices(ctx)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get services")
		return nil, fmt.Errorf("failed to get services: %w", err)
	}

	return services, nil
}

// GetServiceByID retrieves a service by ID
func (s *ServiceService) GetServiceByID(ctx context.Context, id int64) (*entity.Service, error) {
	service, err := s.repo.GetServiceByID(ctx, id)
	if err != nil {
		log.Error().Err(err).Int64("serviceID", id).Msg("Failed to get service by ID")
		return nil, fmt.Errorf("failed to get service by ID: %w", err)
	}

	if service == nil {
		return nil, fmt.Errorf("service not found")
	}

	return service, nil
}

// GetServiceWithFeatures retrieves a service with its features by ID
func (s *ServiceService) GetServiceWithFeatures(ctx context.Context, id int64) (*entity.ServiceWithFeatures, error) {
	serviceWithFeatures, err := s.repo.GetServiceWithFeatures(ctx, id)
	if err != nil {
		log.Error().Err(err).Int64("serviceID", id).Msg("Failed to get service with features")
		return nil, fmt.Errorf("failed to get service with features: %w", err)
	}

	if serviceWithFeatures == nil {
		return nil, fmt.Errorf("service not found")
	}

	return serviceWithFeatures, nil
}

// GetServicesByIDs retrieves multiple services by their IDs
func (s *ServiceService) GetServicesByIDs(ctx context.Context, ids []int64) (map[int64]entity.Service, error) {
	if len(ids) == 0 {
		return make(map[int64]entity.Service), nil
	}

	services, err := s.repo.GetServicesByIDs(ctx, ids)
	if err != nil {
		log.Error().Err(err).Interface("serviceIDs", ids).Msg("Failed to get services by IDs")
		return nil, fmt.Errorf("failed to get services by IDs: %w", err)
	}

	return services, nil
}

// ValidateServices checks if all service IDs exist and are active
func (s *ServiceService) ValidateServices(ctx context.Context, serviceIDs []int64) (map[int64]entity.Service, error) {
	if len(serviceIDs) == 0 {
		return nil, fmt.Errorf("no services specified")
	}

	services, err := s.GetServicesByIDs(ctx, serviceIDs)
	if err != nil {
		return nil, err
	}

	// Check if all services exist
	for _, id := range serviceIDs {
		if _, ok := services[id]; !ok {
			return nil, fmt.Errorf("service with ID %d not found", id)
		}
	}

	return services, nil
}