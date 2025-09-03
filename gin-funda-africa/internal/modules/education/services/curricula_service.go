package services

import (
	"gin-funda-africa/internal/modules/education/models"
	"gin-funda-africa/internal/modules/education/repository"

	"github.com/google/uuid"
)

type CurriculaService struct {}

func NewCurriculaService() *CurriculaService {
	return &CurriculaService{}
}

func (s *CurriculaService) GetCurriculumById(id uuid.UUID) (*models.Curriculum, error) {
	return repository.GetCurriculum(id)
}

func (s *CurriculaService) GetCurriculaByCountry(country string, page, pageSize int) (*repository.PaginatedCurriculums, error) {
	return repository.GetCurriculaByCountry(country, page, pageSize)
}

func (s *CurriculaService) GetCurriculaByBoard(board string, page, pageSize int) (*repository.PaginatedCurriculums, error) {
	return repository.GetCurriculaByBoard(board, page, pageSize)
}

func (s *CurriculaService) GetCurriculaByVersion(version uuid.UUID, page, pageSize int) (*repository.PaginatedCurriculums, error) {
	return repository.GetCurriculaByVersion(version, page, pageSize)
}
