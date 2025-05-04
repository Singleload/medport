package dto

import "github.com/svenskhalsovard/api/internal/entity"

// ServiceResponse represents a service in the API response
type ServiceResponse struct {
	ID                  int64    `json:"id"`
	Name                string   `json:"name"`
	ShortDescription    string   `json:"shortDescription"`
	Description         string   `json:"description"`
	Price               float64  `json:"price"`
	DiscountedPrice     *float64 `json:"discountedPrice,omitempty"`
	IsSubscription      bool     `json:"isSubscription"`
	SubscriptionInterval string   `json:"subscriptionInterval,omitempty"`
	Image               string   `json:"image"`
	Features            []string `json:"features,omitempty"`
}

// ServiceListResponse represents a list of services in the API response
type ServiceListResponse struct {
	Services []ServiceResponse `json:"services"`
	Count    int               `json:"count"`
}

// MapServiceToResponse maps an entity.Service to a ServiceResponse
func MapServiceToResponse(service entity.Service, features []entity.ServiceFeature) ServiceResponse {
	featureStrings := make([]string, 0, len(features))
	for _, feature := range features {
		featureStrings = append(featureStrings, feature.Feature)
	}

	return ServiceResponse{
		ID:                  service.ID,
		Name:                service.Name,
		ShortDescription:    service.ShortDescription,
		Description:         service.Description,
		Price:               service.Price,
		DiscountedPrice:     service.DiscountedPrice,
		IsSubscription:      service.IsSubscription,
		SubscriptionInterval: service.SubscriptionInterval,
		Image:               service.Image,
		Features:            featureStrings,
	}
}

// MapServiceWithFeaturesToResponse maps an entity.ServiceWithFeatures to a ServiceResponse
func MapServiceWithFeaturesToResponse(serviceWithFeatures entity.ServiceWithFeatures) ServiceResponse {
	return MapServiceToResponse(serviceWithFeatures.Service, serviceWithFeatures.Features)
}

// MapServicesToResponse maps a slice of entity.Service to a ServiceListResponse
func MapServicesToResponse(services []entity.Service, featuresMap map[int64][]entity.ServiceFeature) ServiceListResponse {
	responseServices := make([]ServiceResponse, 0, len(services))
	for _, service := range services {
		features := featuresMap[service.ID]
		responseServices = append(responseServices, MapServiceToResponse(service, features))
	}

	return ServiceListResponse{
		Services: responseServices,
		Count:    len(responseServices),
	}
}