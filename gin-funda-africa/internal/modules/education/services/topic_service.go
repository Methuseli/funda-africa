package services

import (
	"gin-funda-africa/internal/modules/education/models"
	"gin-funda-africa/internal/modules/education/repository"

	"github.com/google/uuid"
)

type TopicService struct {}

func NewTopicService() *TopicService {
	return &TopicService{}
}

func (s *TopicService) GetTopicByID(id uuid.UUID) (*models.Topic, error) {
	return repository.GetTopicByID(id)
}

func (s *TopicService) GetTopicsBySubjectID(subject uuid.UUID, page, pageSize int) (*repository.PaginatedTopics, error) {
	return repository.GetTopicsBySubjectID(subject, page, pageSize)
}
