package dto

import (
	"time"

	"github.com/google/uuid"
)

type SubscriptionPlanResponse struct {
	ID           uuid.UUID `json:"id" example:"550e8400-e29b-41d4-a716-446655440000"`
	Name         string    `json:"name" example:"Premium Plan"`
	Description  string    `json:"description" example:"Full access to all courses and premium features"`
	PriceUSD     float64   `json:"price_usd" example:"49.99"`
	FeaturesJSON string    `json:"features" example:"[\"Unlimited access to videos\", \"Offline downloads\", \"Priority support\"]"`
	IsActive     bool      `json:"is_active" example:"true"`
	CreatedAt    time.Time `json:"created_at" example:"2025-09-01T15:04:05Z"`
	UpdatedAt    time.Time `json:"updated_at" example:"2025-09-01T15:04:05Z"`
}

