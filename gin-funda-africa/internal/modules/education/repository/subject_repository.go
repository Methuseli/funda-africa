package repository

import (
	"gin-funda-africa/internal/modules/education/dto"
	"gin-funda-africa/internal/modules/education/models"
	"gin-funda-africa/internal/modules/users"
	"gin-funda-africa/internal/pkg/utils"
	"reflect"

	"github.com/google/uuid"
)

func GetSubjectByID(id uuid.UUID) (*models.Subject, error) {
	var subject models.Subject
	if err := users.Database.First(&subject, utils.IdQuery, id).Error; err != nil {
		return nil, err
	}
	return &subject, nil
}

type PaginatedSubjects struct {
	Data       []dto.SubjectResponse `json:"data"`
	TotalCount int64                 `json:"total_count"`
	Page       int                   `json:"page"`
	PageSize   int                   `json:"page_size"`
}

func GetSubjects(params utils.Params, page, pageSize int) (*PaginatedSubjects, error) {
	var subjects []*models.Subject
	var responses []*dto.SubjectResponse

	results, err := utils.GetPaginatedData(&utils.PaginateData{
		Model:          reflect.TypeOf(subjects),
		Dto:            reflect.TypeOf(responses),
		Page:           page,
		PageSize:       pageSize,
		Params: params,
	})

	return &PaginatedSubjects{
		Data:       utils.ConvertResults[dto.SubjectResponse](results.Data),
		TotalCount: results.TotalCount,
		Page:       results.Page,
		PageSize:   results.PageSize,
	}, err
}

