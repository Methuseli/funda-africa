package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Version struct {
	gorm.Model
	ID   uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	name string    `gorm:"not null"`
}

type Curriculum struct {
	gorm.Model
	ID        uuid.UUID  `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name      string     `gorm:"not null"`
	Country   string     `gorm:"not null"`
	Board     string     `gorm:"not null"`
	VersionID uuid.UUID  `gorm:"type:uuid;not null"`
	StartDate *time.Time `gorm:"not null"`
	EndDate   *time.Time `gorm:"not null"`
	CreatedAt time.Time  `gorm:"autoCreateTime"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime"`
	Version   Version    `gorm:"foreignKey:VersionId"`
}

type Subject struct {
	gorm.Model
	ID           uuid.UUID  `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	CurriculumID uuid.UUID  `gorm:"type:uuid;not null"`
	Name         string     `gorm:"not null"`
	Level        string     `gorm:"not null"`
	Code         string     `gorm:"not null"`
	Curriculum   Curriculum `gorm:"foreignKey:CurriculumID"`
	CreatedAt    time.Time  `gorm:"autoCreateTime"`
	UpdatedAt    time.Time  `gorm:"autoUpdateTime"`
}

type Topic struct {
	gorm.Model
	ID         uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	SubjectID  uuid.UUID `gorm:"type:uuid;not null"`
	Name       string    `gorm:"not null"`
	OrderIndex string    `gorm:"not null"`
	Subject    Subject   `gorm:"foreignKey:SubjectID"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`
}

type SubTopic struct {
	gorm.Model
	ID         uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	TopicID    uuid.UUID `gorm:"type:uuid;not null"`
	Name       string    `gorm:"not null"`
	OrderIndex string    `gorm:"not null"`
	Code       string    `gorm:"not null"`
	Topic      Topic     `gorm:"foreignKey:TopicID"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`
}

type Objective struct {
	gorm.Model
	ID         uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	SubTopicID uuid.UUID `gorm:"type:uuid;not null"`
	Text       string    `gorm:"not null"`
	OrderIndex string    `gorm:"not null"`
	SubTopic   SubTopic  `gorm:"foreignKey:SubTopicID"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`
}

type Content struct {
	gorm.Model
	ID         uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	SubTopicID uuid.UUID `gorm:"type:uuid;not null"`
	Text       string    `gorm:"not null"`
	OrderIndex string    `gorm:"not null"`
	SubTopic   SubTopic  `gorm:"foreignKey:SubTopicID"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`
}

type Prerequisite struct {
	gorm.Model
	ID         uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	SubTopicID uuid.UUID `gorm:"type:uuid;not null"`
	Text       string    `gorm:"not null"`
	OrderIndex string    `gorm:"not null"`
	SubTopic   SubTopic  `gorm:"foreignKey:SubTopicID"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`
}

type SuggestedResource struct {
	gorm.Model
	ID         uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	SubTopicID uuid.UUID `gorm:"type:uuid;not null"`
	Text       string    `gorm:"not null"`
	OrderIndex string    `gorm:"not null"`
	SubTopic   SubTopic  `gorm:"foreignKey:SubTopicID"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`
}

type SuggestedNote struct {
	gorm.Model
	ID         uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	SubTopicID uuid.UUID `gorm:"type:uuid;not null"`
	Text       string    `gorm:"not null"`
	OrderIndex string    `gorm:"not null"`
	SubTopic   SubTopic  `gorm:"foreignKey:SubTopicID"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`
}

type Chunk struct {
	gorm.Model
	ID           uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	SubTopicID   uuid.UUID `gorm:"type:uuid;not null"`
	ReadingTime  string    `gorm:"not null"`
	Type         string    `gorm:"not null"`
	ChunkContent []byte    `gorm:"type:jsonb;not null"`
	OrderIndex   string    `gorm:"not null"`
	Difficulty   string    `gorm:"not null"`
	Language     string    `gorm:"not null"`
	Title        string    `gorm:"not null"`
	SubTopic     SubTopic  `gorm:"foreignKey:SubTopicID"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime"`
}
