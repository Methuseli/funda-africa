package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserOrganization struct {
    gorm.Model
    UserID         uuid.UUID `gorm:"type:uuid;not null"`
    OrganizationID uuid.UUID `gorm:"type:uuid;not null"`
    Position       string
    User           User        `gorm:"foreignKey:UserID"`
    Organization   Organization `gorm:"foreignKey:OrganizationID"`
	CreatedAt      time.Time `gorm:"autoCreateTime"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime"`
}
