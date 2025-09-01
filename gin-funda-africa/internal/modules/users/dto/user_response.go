package dto

import (
	"time"

	"github.com/google/uuid"
)

type UserResponse struct {
	ID             uuid.UUID  `json:"id" example:"123e4567-e89b-12d3-a456-426614174000"`
	Username       string     `json:"username" example:"john_doe"`
	Email          string     `json:"email" example:"john.doe@example.com"`
	PhoneNumber    string     `json:"phone_number" example:"+263771234567"`
	FirstName      string     `json:"first_name" example:"John"`
	LastName       string     `json:"last_name" example:"Doe"`
	MiddleName     string     `json:"middle_name" example:"Michael"`
	Role           string     `json:"role" example:"student"`
	ProfilePicture string     `json:"profile_picture" example:"https://example.com/images/john_doe.png"`
	DateOfBirth    *time.Time `json:"date_of_birth" example:"2000-01-15T00:00:00Z"`
	Gender         string    `json:"gender" example:"male"`
	CountryCode    string     `json:"country_code" example:"ZW"`
	CreatedAt      time.Time  `json:"created_at" example:"2025-01-01T12:00:00Z"`
	UpdatedAt      time.Time  `json:"updated_at" example:"2025-08-01T12:00:00Z"`
}

type LoginResponse struct {
    User  UserResponse `json:"user"`
    Token *string      `json:"token,omitempty" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
}


