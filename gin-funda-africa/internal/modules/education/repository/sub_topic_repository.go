package repository

import (
	"gin-funda-africa/internal/modules/education/dto"
	"gin-funda-africa/internal/modules/education/education_models"
	"gin-funda-africa/internal/modules/users"
	"gin-funda-africa/internal/pkg/utils"
	"reflect"

	"github.com/google/uuid"
)

const subTopicIdQuery = "subject_id = ?"

func GetSubTopicByID(id uuid.UUID) (*education_models.SubTopic, error) {
	var subTopic education_models.SubTopic
	if err := users.Database.First(&subTopic, utils.IdQuery, id).Error; err != nil {
		return nil, err
	}

	return &subTopic, nil
}

type PaginatedSubTopics struct {
	Data       []dto.SubTopicResponse `json:"data"`
	TotalCount int64                  `json:"total_count"`
	Page       int                    `json:"page"`
	PageSize   int                    `json:"page_size"`
}

func GetSubTopics(params utils.Params, page, pageSize int) (*PaginatedSubTopics, error) {
	var topics []*education_models.SubTopic
	var responses []*dto.SubTopicResponse

	results, err := utils.GetPaginatedData(&utils.PaginateData{
		Model:          reflect.TypeOf(topics),
		Dto:            reflect.TypeOf(responses),
		Page:           page,
		PageSize:       pageSize,
		Params: params,
	})

	return &PaginatedSubTopics{
		Data:       utils.ConvertResults[dto.SubTopicResponse](results.Data),
		TotalCount: results.TotalCount,
		Page:       results.Page,
		PageSize:   results.PageSize,
	}, err
}

func GetObjectivesBySubTopic(subTopic uuid.UUID) ([]dto.ObjectivesResponse, error) {
	var objectives []*education_models.Objective
	var responses []*dto.ObjectivesResponse

	results, err := utils.QueryList(&utils.QueryListData{
		Model:          reflect.TypeOf(objectives),
		Filter:         subTopicIdQuery,
		Dto:            reflect.TypeOf(responses),
		FilterVariable: subTopic,
	})
	return utils.ConvertResults[dto.ObjectivesResponse](results), err
}

func GetContentBySubTopic(subTopic uuid.UUID) ([]dto.ContentResponse, error) {
	var contents []*education_models.Content
	var responses []*dto.ContentResponse

	results, err := utils.QueryList(&utils.QueryListData{
		Model:          reflect.TypeOf(contents),
		Filter:         subTopicIdQuery,
		Dto:            reflect.TypeOf(responses),
		FilterVariable: subTopic,
	})
	return utils.ConvertResults[dto.ContentResponse](results), err
}

func GetPrerequisitesBySubTopic(subTopic uuid.UUID) ([]dto.PrerequisiteResponse, error) {
	var prerequisites []*education_models.Prerequisite
	var responses []*dto.PrerequisiteResponse

	results, err := utils.QueryList(&utils.QueryListData{
		Model:          reflect.TypeOf(prerequisites),
		Filter:         subTopicIdQuery,
		Dto:            reflect.TypeOf(responses),
		FilterVariable: subTopic,
	})

	return utils.ConvertResults[dto.PrerequisiteResponse](results), err
}

func GetSuggestedResourceBySubTopic(subTopic uuid.UUID) ([]dto.SuggestedResourceResponse, error) {
	var suggestedResources []*education_models.SuggestedResource
	var responses []*dto.SuggestedResourceResponse

	results, err := utils.QueryList(&utils.QueryListData{
		Model:          reflect.TypeOf(suggestedResources),
		Filter:         subTopicIdQuery,
		Dto:            reflect.TypeOf(responses),
		FilterVariable: subTopic,
	})

	return utils.ConvertResults[dto.SuggestedResourceResponse](results), err
}

func GetSuggestedNotesBySubTopic(subTopic uuid.UUID) ([]dto.SuggestedNoteResponse, error) {
	var suggestedNotes []*education_models.SuggestedNote
	var responses []*dto.SuggestedNoteResponse

	results, err := utils.QueryList(&utils.QueryListData{
		Model:          reflect.TypeOf(suggestedNotes),
		Filter:         subTopicIdQuery,
		Dto:            reflect.TypeOf(responses),
		FilterVariable: subTopic,
	})

	return utils.ConvertResults[dto.SuggestedNoteResponse](results), err
}

func GetChunksBySubTopic(subTopic uuid.UUID) ([]dto.ChunkResponse, error) {
	var chunks []*education_models.Chunk
	var responses []*dto.ChunkResponse

	results, err := utils.QueryList(&utils.QueryListData{
		Model:          reflect.TypeOf(chunks),
		Filter:         subTopicIdQuery,
		Dto:            reflect.TypeOf(responses),
		FilterVariable: subTopic,
	})

	return utils.ConvertResults[dto.ChunkResponse](results), err
}
