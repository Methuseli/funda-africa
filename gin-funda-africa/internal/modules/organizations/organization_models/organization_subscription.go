package organization_models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OrganizationSubscription struct {
	gorm.Model
	OrganizationID uuid.UUID    `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	PlanID         uuid.UUID    `gorm:"not null"`
	StartDate      *time.Time   `gorm:"not null"`
	EndDate        *time.Time   `gorm:"not null"`
	IsActive       bool         `gorm:"default:true"`
	OrganizationSize int 
	Organization   Organization `gorm:"foreignKey:OrganizationID"`
	CreatedAt      time.Time    `gorm:"autoCreateTime"`
	UpdatedAt      time.Time    `gorm:"autoUpdateTime"`
}