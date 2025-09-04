package dto

import (
	"time"

	"github.com/google/uuid"
)

type OrganizationSubscriptionResponse struct {
	OrganizationID   uuid.UUID  `json:"organization_id" example:"550e8400-e29b-41d4-a716-446655440000"`
	PlanID           uuid.UUID  `json:"plan_id" example:"550e8400-e29b-41d4-a716-446655440fg4"`
	StartDate        *time.Time `json:"start_date" example:"2025-09-01T15:04:05Z"`
	EndDate          *time.Time `json:"end_date" example:"2025-09-01T15:04:05Z"`
	IsActive         bool       `json:"is_active" example:"true"`
	OrganizationSize int        `json:"organization_size" example:"50"`
	CreatedAt        time.Time  `json:"created_at" example:"2025-09-01T15:04:05Z"`
	UpdatedAt        time.Time  `json:"updated_at" example:"2025-09-01T15:04:05Z"`
}
