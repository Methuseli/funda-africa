package services

import (
	"gin-funda-africa/internal/modules/education/dto"
	"gin-funda-africa/internal/modules/education/repository"
	"gin-funda-africa/internal/pkg/utils"

	"github.com/google/uuid"
	"github.com/jinzhu/copier"
)

type SubTopicService struct {}

func NewSubTopicService() *SubTopicService {
	return &SubTopicService{}
}

func (s *SubTopicService) GetSubTopicByID(id uuid.UUID) (*dto.SubTopicResponse, error) {
	objectives, err := repository.GetObjectivesBySubTopic(id)
	contents, err := repository.GetContentBySubTopic(id)
	prerequisites, err := repository.GetPrerequisitesBySubTopic(id)
	suggestedResources, err := repository.GetSuggestedResourceBySubTopic(id)
	suggestedNotes, err := repository.GetSuggestedNotesBySubTopic(id)
	chunks, err := repository.GetChunksBySubTopic(id)
	subTopic, err := repository.GetSubTopicByID(id)

	if err != nil {
		return nil, err
	}

	var subTopicResponse dto.SubTopicResponse
	err = copier.Copy(&subTopicResponse, subTopic)
	if err != nil {
		return nil, err
	}
	subTopicResponse.Chunks = chunks
	subTopicResponse.Contents = contents
	subTopicResponse.Objectives = objectives
	subTopicResponse.Prerequisites = prerequisites
	subTopicResponse.SuggestedResources = suggestedResources
	subTopicResponse.SuggestedNotes = suggestedNotes

	return &subTopicResponse, nil
}

func (s *SubTopicService) GetSubTopics(params utils.Params, page, pageSize int) (*repository.PaginatedSubTopics, error) {
	return repository.GetSubTopics(params, page, pageSize)
}
