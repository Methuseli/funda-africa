package dto

import (
	"time"

	"github.com/google/uuid"
)

type OrganizationUserResponse struct {
	UserID         uuid.UUID `json:"user_id" example:"123e4567-e89b-12d3-a456-426614174000"`
	OrganizationID uuid.UUID `json:"organization_id" example:"123e4567-e89b-12d3-a456-426614174000"`
	Position       string    `json:"position" example:"Teacher"`
	CreatedAt      time.Time `json:"created_at" example:"2025-01-01T12:00:00Z"`
	UpdatedAt      time.Time `json:"updated_at" example:"2025-08-01T12:00:00Z"`
}
