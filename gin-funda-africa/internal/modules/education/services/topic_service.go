package services

import (
	"gin-funda-africa/internal/modules/education/models"
	"gin-funda-africa/internal/modules/education/repository"
	"gin-funda-africa/internal/pkg/utils"

	"github.com/google/uuid"
)

type TopicService struct {}

func NewTopicService() *TopicService {
	return &TopicService{}
}

func (s *TopicService) GetTopicByID(id uuid.UUID) (*models.Topic, error) {
	return repository.GetTopicByID(id)
}

func (s *TopicService) GetTopics(params utils.Params, page, pageSize int) (*repository.PaginatedTopics, error) {
	return repository.GetTopicsBySubjectID(params, page, pageSize)
}
