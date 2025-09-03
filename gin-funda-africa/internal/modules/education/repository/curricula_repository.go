package repository

import (
	"gin-funda-africa/internal/modules/education/dto"
	"gin-funda-africa/internal/modules/education/models"
	"gin-funda-africa/internal/modules/users"
	"gin-funda-africa/internal/pkg/utils"
	"reflect"

	"github.com/google/uuid"
)

func GetCurriculum(id uuid.UUID) (*models.Curriculum, error) {
	var curriculum models.Curriculum
	if err := users.Database.First(&curriculum, utils.IdQuery, id).Error; err != nil {
		return nil, err
	}
	return &curriculum, nil
}

type PaginatedCurriculums struct {
	Data       []dto.CurriculumResponse `json:"data"`
	TotalCount int64                    `json:"total_count"`
	Page       int                      `json:"page"`
	PageSize   int                      `json:"page_size"`
}

func GetCurricula(params utils.Params, page, pageSize int) (*PaginatedCurriculums, error) {
	var curricula []*models.Curriculum
	var responses []*dto.CurriculumResponse

	results, err := utils.GetPaginatedData(&utils.PaginateData{
		Model:          reflect.TypeOf(curricula),
		Dto:            reflect.TypeOf(responses),
		Page:           page,
		PageSize:       pageSize,
		Params: params,
	})

	return &PaginatedCurriculums{
		Data:       utils.ConvertResults[dto.CurriculumResponse](results.Data),
		TotalCount: results.TotalCount,
		Page:       results.Page,
		PageSize:   results.PageSize,
	}, err
}

