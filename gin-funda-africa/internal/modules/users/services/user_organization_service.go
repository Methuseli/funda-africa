package services

import (
	"gin-funda-africa/internal/modules/users/models"
	"gin-funda-africa/internal/modules/users/repository"

	"github.com/google/uuid"
)

type UserOrganizationService struct{}

func NewUserOrganizationService() *UserOrganizationService {
	return &UserOrganizationService{}
}

func (s *UserOrganizationService) CreateUserOrganization(uo *models.UserOrganization) (*models.UserOrganization, error) {
	return repository.SaveUserOrganization(uo)
}

func (s *UserOrganizationService) GetUserOrganization(userID, orgID uuid.UUID) (*models.UserOrganization, error) {
	return repository.GetUserOrganization(userID, orgID)
}

func (s *UserOrganizationService) ListUserOrganizations(page, pageSize int) (*repository.PaginatedUserOrganizations, error) {
	return repository.FindAllUserOrganizations(page, pageSize)
}

func (s *UserOrganizationService) UpdateUserOrganizationPartial(userID, orgID uuid.UUID, patch *repository.UserOrganizationPatch) (*models.UserOrganization, error) {
	return repository.UpdateUserOrganization(userID, orgID, patch)
}

func (s *UserOrganizationService) DeleteUserOrganization(userID, orgID uuid.UUID) error {
	return repository.DeleteUserOrganization(userID, orgID)
}
