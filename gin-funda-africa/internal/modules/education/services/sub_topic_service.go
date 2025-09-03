package services

import (
	"gin-funda-africa/internal/modules/education/dto"
	"gin-funda-africa/internal/modules/education/models"
	"gin-funda-africa/internal/modules/education/repository"

	"github.com/google/uuid"
)

type SubTopicService struct {}

func NewSubTopicService() *SubTopicService {
	return &SubTopicService{}
}

func (s *SubTopicService) GetSubTopicByID(id uuid.UUID) (*models.SubTopic, error) {
	return repository.GetSubTopicByID(id)
}

func (s *SubTopicService) GetSubTopicByTopicID(topic uuid.UUID, page, pageSize int) (*repository.PaginatedSubTopics, error) {
	return repository.GetSubTopicByTopicID(topic, page, pageSize)
}

func (s *SubTopicService) GetObjectivesBySubTopic(subTopic uuid.UUID) ([]dto.ObjectivesResponse, error) {
	return repository.GetObjectivesBySubTopic(subTopic)
}

func (s *SubTopicService) GetContentBySubTopic(subTopic uuid.UUID) ([]dto.ContentResponse, error) {
	return repository.GetContentBySubTopic(subTopic)
}

func (s *SubTopicService) GetPrerequisitesBySubTopic(subTopic uuid.UUID) ([]dto.PrerequisiteResponse, error) {
	return repository.GetPrerequisitesBySubTopic(subTopic)
}

func (s *SubTopicService) GetSuggestedResourceBySubTopic(subTopic uuid.UUID) ([]dto.SuggestedResourceResponse, error) {
	return repository.GetSuggestedResourceBySubTopic(subTopic)
}

func (s *SubTopicService) GetSuggestedNotesBySubTopic(subTopic uuid.UUID) ([]dto.SuggestedNoteResponse, error) {
	return repository.GetSuggestedNotesBySubTopic(subTopic)
}

func (s *SubTopicService) GetChunksBySubTopic(subTopic uuid.UUID) ([]dto.ChunkResponse, error) {
	return repository.GetChunksBySubTopic(subTopic)
}