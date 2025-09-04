package services

import (
	"gin-funda-africa/internal/modules/organizations/organization_models"
	"gin-funda-africa/internal/modules/organizations/repository"

	"github.com/google/uuid"
)

type OrganizationSubscriptionService struct{}

func NewOrganizationSubscriptionService() *OrganizationSubscriptionService {
	return &OrganizationSubscriptionService{}
}

func (s *OrganizationSubscriptionService) CreateOrganizationSubscription(sub *organization_models.OrganizationSubscription) (*organization_models.OrganizationSubscription, error) {
	return repository.SaveOrganizationSubscription(sub)
}

func (s *OrganizationSubscriptionService) GetOrganizationSubscription(userID, planID uuid.UUID) (*organization_models.OrganizationSubscription, error) {
	return repository.GetOrganizationSubscription(userID, planID)
}

func (s *OrganizationSubscriptionService) ListOrganizationSubscriptions(page, pageSize int) (*repository.PaginatedOrganizationSubscriptions, error) {
	return repository.FindAllOrganizationSubscriptions(page, pageSize)
}

func (s *OrganizationSubscriptionService) UpdateOrganizationSubscriptionPartial(userID, planID uuid.UUID, patch *repository.OrganizationSubscriptionPatch) (*organization_models.OrganizationSubscription, error) {
	return repository.UpdateOrganizationSubscriptionPartial(userID, planID, patch)
}
