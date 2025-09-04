package services

import (
	"gin-funda-africa/internal/modules/organizations/organization_models"
	"gin-funda-africa/internal/modules/organizations/repository"

	"github.com/google/uuid"
)

type OrganizationUserService struct{}

func NewOrganizationUserService() *OrganizationUserService {
	return &OrganizationUserService{}
}

func (s *OrganizationUserService) CreateOrganizationUser(uo *organization_models.OrganizationUser) (*organization_models.OrganizationUser, error) {
	return repository.SaveOrganizationUser(uo)
}

func (s *OrganizationUserService) GetOrganizationUser(userID, orgID uuid.UUID) (*organization_models.OrganizationUser, error) {
	return repository.GetOrganizationUser(userID, orgID)
}

func (s *OrganizationUserService) ListUserByOrganization(id uuid.UUID, page, pageSize int) (*repository.PaginatedOrganizationUsers, error) {
	return repository.FindByOrganization(id, page, pageSize)
}

func (s *OrganizationUserService) UpdateOrganizationUserPartial(userID, orgID uuid.UUID, patch *repository.OrganizationUserPatch) (*organization_models.OrganizationUser, error) {
	return repository.UpdateOrganizationUser(userID, orgID, patch)
}

