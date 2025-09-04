package organization_models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Organization struct {
	gorm.Model
	ID          uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name        string    `gorm:"not null"`
	Type        string    `gorm:"not null"`
	CountryCode string    `gorm:"size:2;not null"`
	City        string
	UserID      uuid.UUID `gorm:"type:uuid"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}
