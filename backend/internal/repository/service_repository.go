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

// ServiceRepository handles database operations for services
type ServiceRepository struct {
	db *sqlx.DB
}

// NewServiceRepository creates a new ServiceRepository
func NewServiceRepository(database *Database) *ServiceRepository {
	return &ServiceRepository{
		db: database.DB,
	}
}

// GetServices retrieves all active services
func (r *ServiceRepository) GetServices(ctx context.Context) ([]entity.Service, error) {
	query := `
		SELECT id, name, short_description, description, price, discounted_price, 
		       is_subscription, subscription_interval, image, is_active, 
		       created_at, updated_at, deleted_at
		FROM services
		WHERE ` + softDeleteCondition("services") + `
		AND is_active = true
		ORDER BY id
	`

	var services []entity.Service
	if err := r.db.SelectContext(ctx, &services, query); err != nil {
		return nil, fmt.Errorf("failed to get services: %w", err)
	}

	return services, nil
}

// GetServiceByID retrieves a service by ID
func (r *ServiceRepository) GetServiceByID(ctx context.Context, id int64) (*entity.Service, error) {
	query := `
		SELECT id, name, short_description, description, price, discounted_price, 
		       is_subscription, subscription_interval, image, is_active, 
		       created_at, updated_at, deleted_at
		FROM services
		WHERE ` + softDeleteCondition("services") + `
		AND is_active = true
		AND id = ?
	`

	var service entity.Service
	if err := r.db.GetContext(ctx, &service, query, id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // Service not found
		}
		return nil, fmt.Errorf("failed to get service by ID: %w", err)
	}

	return &service, nil
}

// GetServiceWithFeatures retrieves a service with its features by ID
func (r *ServiceRepository) GetServiceWithFeatures(ctx context.Context, id int64) (*entity.ServiceWithFeatures, error) {
	service, err := r.GetServiceByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if service == nil {
		return nil, nil // Service not found
	}

	features, err := r.GetServiceFeatures(ctx, id)
	if err != nil {
		return nil, err
	}

	return &entity.ServiceWithFeatures{
		Service:  *service,
		Features: features,
	}, nil
}

// GetServiceFeatures retrieves the features for a service
func (r *ServiceRepository) GetServiceFeatures(ctx context.Context, serviceID int64) ([]entity.ServiceFeature, error) {
	query := `
		SELECT id, service_id, feature, created_at, updated_at, deleted_at
		FROM service_features
		WHERE ` + softDeleteCondition("service_features") + `
		AND service_id = ?
		ORDER BY id
	`

	var features []entity.ServiceFeature
	if err := r.db.SelectContext(ctx, &features, query, serviceID); err != nil {
		return nil, fmt.Errorf("failed to get service features: %w", err)
	}

	return features, nil
}

// GetServicesByIDs retrieves services by their IDs
func (r *ServiceRepository) GetServicesByIDs(ctx context.Context, ids []int64) (map[int64]entity.Service, error) {
	if len(ids) == 0 {
		return make(map[int64]entity.Service), nil
	}

	query, args, err := sqlx.In(`
		SELECT id, name, short_description, description, price, discounted_price, 
		       is_subscription, subscription_interval, image, is_active, 
		       created_at, updated_at, deleted_at
		FROM services
		WHERE `+softDeleteCondition("services")+`
		AND is_active = true
		AND id IN (?)
	`, ids)

	if err != nil {
		return nil, fmt.Errorf("failed to build query: %w", err)
	}

	query = r.db.Rebind(query)
	var services []entity.Service
	if err := r.db.SelectContext(ctx, &services, query, args...); err != nil {
		return nil, fmt.Errorf("failed to get services by IDs: %w", err)
	}

	// Map services by ID for easy lookup
	serviceMap := make(map[int64]entity.Service, len(services))
	for _, service := range services {
		serviceMap[service.ID] = service
	}

	return serviceMap, nil
}

// CreateService creates a new service
func (r *ServiceRepository) CreateService(ctx context.Context, service *entity.Service) error {
	query := `
		INSERT INTO services (
			name, short_description, description, price, discounted_price,
			is_subscription, subscription_interval, image, is_active,
			created_at, updated_at
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	now := now()
	service.CreatedAt = now
	service.UpdatedAt = now

	result, err := r.db.ExecContext(
		ctx,
		query,
		service.Name,
		service.ShortDescription,
		service.Description,
		service.Price,
		service.DiscountedPrice,
		service.IsSubscription,
		service.SubscriptionInterval,
		service.Image,
		service.IsActive,
		service.CreatedAt,
		service.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to create service: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("failed to get last insert ID: %w", err)
	}

	service.ID = id
	return nil
}

// UpdateService updates an existing service
func (r *ServiceRepository) UpdateService(ctx context.Context, service *entity.Service) error {
	query := `
		UPDATE services
		SET name = ?,
			short_description = ?,
			description = ?,
			price = ?,
			discounted_price = ?,
			is_subscription = ?,
			subscription_interval = ?,
			image = ?,
			is_active = ?,
			updated_at = ?
		WHERE id = ?
		AND ` + softDeleteCondition("services")
	
	service.UpdatedAt = now()

	result, err := r.db.ExecContext(
		ctx,
		query,
		service.Name,
		service.ShortDescription,
		service.Description,
		service.Price,
		service.DiscountedPrice,
		service.IsSubscription,
		service.SubscriptionInterval,
		service.Image,
		service.IsActive,
		service.UpdatedAt,
		service.ID,
	)
	if err != nil {
		return fmt.Errorf("failed to update service: %w", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rows == 0 {
		return fmt.Errorf("service not found or already deleted")
	}

	return nil
}

// DeleteService soft-deletes a service
func (r *ServiceRepository) DeleteService(ctx context.Context, id int64) error {
	query := `
		UPDATE services
		SET deleted_at = ?,
			updated_at = ?
		WHERE id = ?
		AND ` + softDeleteCondition("services")
	
	now := now()
	result, err := r.db.ExecContext(ctx, query, now, now, id)
	if err != nil {
		return fmt.Errorf("failed to delete service: %w", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rows == 0 {
		return fmt.Errorf("service not found or already deleted")
	}

	return nil
}

// AddServiceFeature adds a feature to a service
func (r *ServiceRepository) AddServiceFeature(ctx context.Context, feature *entity.ServiceFeature) error {
	query := `
		INSERT INTO service_features (
			service_id, feature, created_at, updated_at
		) VALUES (?, ?, ?, ?)
	`

	now := now()
	feature.CreatedAt = now
	feature.UpdatedAt = now

	result, err := r.db.ExecContext(
		ctx,
		query,
		feature.ServiceID,
		feature.Feature,
		feature.CreatedAt,
		feature.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to add service feature: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("failed to get last insert ID: %w", err)
	}

	feature.ID = id
	return nil
}

// UpdateServiceFeature updates a service feature
func (r *ServiceRepository) UpdateServiceFeature(ctx context.Context, feature *entity.ServiceFeature) error {
	query := `
		UPDATE service_features
		SET feature = ?,
			updated_at = ?
		WHERE id = ?
		AND ` + softDeleteCondition("service_features")
	
	feature.UpdatedAt = now()

	result, err := r.db.ExecContext(
		ctx,
		query,
		feature.Feature,
		feature.UpdatedAt,
		feature.ID,
	)
	if err != nil {
		return fmt.Errorf("failed to update service feature: %w", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rows == 0 {
		return fmt.Errorf("service feature not found or already deleted")
	}

	return nil
}

// DeleteServiceFeature soft-deletes a service feature
func (r *ServiceRepository) DeleteServiceFeature(ctx context.Context, id int64) error {
	query := `
		UPDATE service_features
		SET deleted_at = ?,
			updated_at = ?
		WHERE id = ?
		AND ` + softDeleteCondition("service_features")
	
	now := now()
	result, err := r.db.ExecContext(ctx, query, now, now, id)
	if err != nil {
		return fmt.Errorf("failed to delete service feature: %w", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rows == 0 {
		return fmt.Errorf("service feature not found or already deleted")
	}

	return nil
}