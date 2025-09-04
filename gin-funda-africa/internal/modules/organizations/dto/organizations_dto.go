package dto

import (
	"time"

	"github.com/google/uuid"
)

type OrganizationResponse struct {
	ID          uuid.UUID `json:"id" example:"123e4567-e89b-12d3-a456-426614174000"`
	Name        string    `json:"name" example:"NGO"`
	Type        string    `json:"type" example:"NGO"`
	CountryCode string    `json:"country_code" example:"ZW"`
	UserID      uuid.UUID `json:"user_id" example:"123e4567-e89b-12d3-a456-426614174000"`
	City        string    `json:"city" example:"Bulawayo"`
	CreatedAt   time.Time `json:"created_at" example:"2025-01-01T12:00:00Z"`
	UpdatedAt   time.Time `json:"updated_at" example:"2025-08-01T12:00:00Z"`
}
