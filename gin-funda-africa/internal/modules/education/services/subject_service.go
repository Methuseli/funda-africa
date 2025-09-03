package services

import (
	"gin-funda-africa/internal/modules/education/models"
	"gin-funda-africa/internal/modules/education/repository"
	"gin-funda-africa/internal/pkg/utils"

	"github.com/google/uuid"
)

type SubjectService struct {}

func NewSubjectService() *SubjectService {
	return &SubjectService{}
}

func (s *SubjectService) GetSubjectByID(id uuid.UUID) (*models.Subject, error){
	return repository.GetSubjectByID(id)
}

func (s *SubjectService) GetSubjects(params utils.Params, page, pageSize int) (*repository.PaginatedSubjects, error) {
	return repository.GetSubjects(params, page, pageSize)
}

