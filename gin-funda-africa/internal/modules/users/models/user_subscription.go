package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserSubscription struct {
    gorm.Model
    UserID    uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
    PlanID    uint      `gorm:"not null"`
    StartDate time.Time `gorm:"not null"`
    EndDate   *time.Time
    IsActive  bool `gorm:"default:true"`
    CreatedAt      time.Time `gorm:"autoCreateTime"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime"`
}