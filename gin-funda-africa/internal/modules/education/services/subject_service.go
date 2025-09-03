package services

import (
	"gin-funda-africa/internal/modules/education/models"
	"gin-funda-africa/internal/modules/education/repository"

	"github.com/google/uuid"
)

type SubjectService struct {}

func NewSubjectService() *SubjectService {
	return &SubjectService{}
}

func (s *SubjectService) GetSubjectByID(id uuid.UUID) (*models.Subject, error){
	return repository.GetSubjectByID(id)
}

func (s *SubjectService) GetSubjectByCurriculumID(curriculum uuid.UUID, page, pageSize int) (*repository.PaginatedSubjects, error) {
	return repository.GetSubjectByCurriculumID(curriculum, page, pageSize)
}

func (s *SubjectService) GetSubjectByCode(code string, page, pageSize int) (*repository.PaginatedSubjects, error) {
	return repository.GetSubjectByCode(code, page, pageSize)
}

func (s *SubjectService) GetSubjectByLevel(level string, page, pageSize int) (*repository.PaginatedSubjects, error) {
	return repository.GetSubjectByLevel(level, page, pageSize)
}