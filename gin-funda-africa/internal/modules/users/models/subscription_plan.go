package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SubscriptionPlan struct {
	gorm.Model
	ID           uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name         string    `json:"name" gorm:"not null"`
	Description  string    `json:"description"`
	PriceUSD     float64   `json:"price_usd" gorm:"not null"`
	FeaturesJSON string    `json:"features"`
	IsActive     bool      `json:"is_active" gorm:"default:true"`
	CreatedAt    time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
