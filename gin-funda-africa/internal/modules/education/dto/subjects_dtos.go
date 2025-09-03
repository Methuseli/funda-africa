package dto

import (
	"time"

	"github.com/google/uuid"
)

type CurriculumResponse struct {
	ID        uuid.UUID  `json:"id"`
	Name      string     `json:"name"`
	Country   string     `json:"country"`
	Board     string     `json:"board"`
	VersionID uuid.UUID  `json:"version"`
	StartDate *time.Time `json:"start_date"`
	EndDate   *time.Time `json:"end_date"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

type SubjectResponse struct {
	ID           uuid.UUID `json:"id"`
	CurriculumID uuid.UUID `json:"curriculum"`
	Name         string    `json:"name"`
	Level        string    `json:"level"`
	Code         string    `json:"code"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type TopicResponse struct {
	ID         uuid.UUID `json:"id"`
	SubjectID  uuid.UUID `json:"subject_id"`
	Name       string    `json:"name"`
	OrderIndex string    `json:"order_index"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type SubTopicResponse struct {
	ID                 uuid.UUID                   `json:"id"`
	TopicID            uuid.UUID                   `json:"topic_id"`
	Name               string                      `json:"name"`
	OrderIndex         string                      `json:"order_index"`
	Code               string                      `json:"code"`
	CreatedAt          time.Time                   `json:"created_at"`
	UpdatedAt          time.Time                   `json:"updated_at"`
	Objectives         []ObjectivesResponse        `json:"objectives"`
	Contents           []ContentResponse           `json:"syllabus_contents"`
	Prerequisites      []PrerequisiteResponse      `json:"prerequisites"`
	SuggestedResources []SuggestedResourceResponse `json:"suggested_resources"`
	SuggestedNotes     []SuggestedNoteResponse     `json:"suggested_notes"`
	Chunks             []ChunkResponse             `json:"chunks"`
}

type ObjectivesResponse struct {
	ID         uuid.UUID `json:"id"`
	SubTopicID uuid.UUID `json:"sub_topic_id"`
	Text       string    `json:"text"`
	OrderIndex string    `json:"order_index"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type ContentResponse struct {
	ID         uuid.UUID `json:"id"`
	SubTopicID uuid.UUID `json:"sub_topic_id"`
	Text       string    `json:"text"`
	OrderIndex string    `json:"order_index"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type PrerequisiteResponse struct {
	ID         uuid.UUID `json:"id"`
	SubTopicID uuid.UUID `json:"sub_topic_id"`
	Text       string    `json:"text"`
	OrderIndex string    `json:"order_index"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type SuggestedResourceResponse struct {
	ID         uuid.UUID `json:"id"`
	SubTopicID uuid.UUID `json:"sub_topic_id"`
	Text       string    `json:"text"`
	OrderIndex string    `json:"order_index"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type SuggestedNoteResponse struct {
	ID         uuid.UUID `json:"id"`
	SubTopicID uuid.UUID `json:"sub_topic_id"`
	Text       string    `json:"text"`
	OrderIndex string    `json:"order_index"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type ChunkResponse struct {
	ID           uuid.UUID `json:"id"`
	SubTopicID   uuid.UUID `json:"sub_topic_id"`
	ReadingTime  string    `json:"reading_time"`
	Type         string    `json:"type"`
	ChunkContent []byte    `json:"chunk_content"`
	OrderIndex   string    `json:"order_index"`
	Difficulty   string    `json:"difficulty"`
	Language     string    `json:"language"`
	Title        string    `json:"title"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
