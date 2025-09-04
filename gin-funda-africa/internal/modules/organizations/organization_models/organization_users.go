package organization_models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OrganizationUser struct {
	gorm.Model
	UserID         uuid.UUID `gorm:"type:uuid;not null"`
	OrganizationID uuid.UUID `gorm:"type:uuid;not null"`
	Position       string
	IsActive       bool         `gorm:"default:true"`
	Organization   Organization `gorm:"foreignKey:OrganizationID"`
	CreatedAt      time.Time    `gorm:"autoCreateTime"`
	UpdatedAt      time.Time    `gorm:"autoUpdateTime"`
}
