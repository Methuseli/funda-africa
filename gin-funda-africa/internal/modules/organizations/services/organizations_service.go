package services

import (
	"gin-funda-africa/internal/modules/organizations/organization_models"
	"gin-funda-africa/internal/modules/organizations/repository"

	"github.com/google/uuid"
)

type OrganizationService struct{}

func NewOrganizationService() *OrganizationService {
	return &OrganizationService{}
}

func (s *OrganizationService) CreateOrganization(organization *organization_models.Organization) (*organization_models.Organization, error) {
	return repository.SaveOrganization(organization)
}

func (s *OrganizationService) GetOrganization(id uuid.UUID) (*organization_models.Organization, error) {
	return repository.GetOrganization(id)
}

func (s *OrganizationService) ListOrganizations(page, pageSize int) (*repository.PaginatedOrganizations, error) {
	return repository.FindAllOrganizations(page, pageSize)
}

func (s *OrganizationService) ListUserOrganizations(page, pageSize int, id uuid.UUID) (*repository.PaginatedOrganizations, error) {
	return repository.FindOrganizationsByUserID(id, page, pageSize)
}

func (s *OrganizationService) UpdateOrganization(id uuid.UUID, patch *repository.OrganizationPatch) (*organization_models.Organization, error) {
	return repository.UpdateOrganization(id, patch)
}

func (s *OrganizationService) DeleteOrganization(id uuid.UUID) error {
	return repository.DeleteOrganization(id)
}
