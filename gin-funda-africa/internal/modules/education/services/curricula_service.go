package services

import (
	"gin-funda-africa/internal/modules/education/models"
	"gin-funda-africa/internal/modules/education/repository"
	"gin-funda-africa/internal/pkg/utils"

	"github.com/google/uuid"
)

type CurriculaService struct {}

func NewCurriculaService() *CurriculaService {
	return &CurriculaService{}
}

func (s *CurriculaService) GetCurriculumById(id uuid.UUID) (*models.Curriculum, error) {
	return repository.GetCurriculum(id)
}

func (s *CurriculaService) GetCurricula(params utils.Params, page, pageSize int) (*repository.PaginatedCurriculums, error) {
	return repository.GetCurricula(params, page, pageSize)
}

