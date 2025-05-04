package entity

import "time"

// Service represents a healthcare service that customers can book
type Service struct {
	ID                  int64      `db:"id" json:"id"`
	Name                string     `db:"name" json:"name"`
	ShortDescription    string     `db:"short_description" json:"shortDescription"`
	Description         string     `db:"description" json:"description"`
	Price               float64    `db:"price" json:"price"`
	DiscountedPrice     *float64   `db:"discounted_price" json:"discountedPrice,omitempty"`
	IsSubscription      bool       `db:"is_subscription" json:"isSubscription"`
	SubscriptionInterval string     `db:"subscription_interval" json:"subscriptionInterval,omitempty"`
	Image               string     `db:"image" json:"image"`
	IsActive            bool       `db:"is_active" json:"isActive"`
	CreatedAt           time.Time  `db:"created_at" json:"createdAt"`
	UpdatedAt           time.Time  `db:"updated_at" json:"updatedAt"`
	DeletedAt           *time.Time `db:"deleted_at" json:"deletedAt,omitempty"`
}

// ServiceFeature represents a feature of a service
type ServiceFeature struct {
	ID        int64      `db:"id" json:"id"`
	ServiceID int64      `db:"service_id" json:"serviceId"`
	Feature   string     `db:"feature" json:"feature"`
	CreatedAt time.Time  `db:"created_at" json:"createdAt"`
	UpdatedAt time.Time  `db:"updated_at" json:"updatedAt"`
	DeletedAt *time.Time `db:"deleted_at" json:"deletedAt,omitempty"`
}

// ServiceWithFeatures represents a service with its features
type ServiceWithFeatures struct {
	Service  Service         `json:"service"`
	Features []ServiceFeature `json:"features"`
}