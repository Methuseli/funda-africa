package repository

import (
	"gin-funda-africa/internal/modules/education/dto"
	"gin-funda-africa/internal/modules/education/models"
	"gin-funda-africa/internal/modules/users"
	"gin-funda-africa/internal/pkg/utils"
	"reflect"

	"github.com/google/uuid"
)

func GetTopicByID(id uuid.UUID) (*models.Topic, error) {
	var topic models.Topic
	if err := users.Database.First(&topic, utils.IdQuery, id).Error; err != nil {
		return nil, err
	}
	return &topic, nil
}

type PaginatedTopics struct {
	Data       []dto.TopicResponse `json:"data"`
	TotalCount int64               `json:"total_count"`
	Page       int                 `json:"page"`
	PageSize   int                 `json:"page_size"`
}

func GetTopicsBySubjectID(subject uuid.UUID, page, pageSize int) (*PaginatedTopics, error) {
	var subjects []*models.Topic
	var responses []*dto.TopicResponse

	results, err := utils.GetPaginatedData(&utils.PaginateData{
		Model:          reflect.TypeOf(subjects),
		Filter:         "subject_id = ?",
		Dto:            reflect.TypeOf(responses),
		Page:           page,
		PageSize:       pageSize,
		FilterVariable: subject,
	})

	return &PaginatedTopics{
		Data:       utils.ConvertResults[dto.TopicResponse](results.Data),
		TotalCount: results.TotalCount,
		Page:       results.Page,
		PageSize:   results.PageSize,
	}, err
}
