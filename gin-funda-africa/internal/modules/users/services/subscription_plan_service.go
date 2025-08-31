package services

import (
	"gin-funda-africa/internal/modules/users/models"
	"gin-funda-africa/internal/modules/users/repository"

	"github.com/google/uuid"
)

type SubscriptionPlanService struct{}

func NewSubscriptionPlanService() *SubscriptionPlanService {
	return &SubscriptionPlanService{}
}

func (s *SubscriptionPlanService) CreateSubscriptionPlan(plan *models.SubscriptionPlan) (*models.SubscriptionPlan, error) {
	return repository.SaveSubscriptionPlan(plan)
}

func (s *SubscriptionPlanService) GetSubscriptionPlan(planID uuid.UUID) (*models.SubscriptionPlan, error) {
	return repository.GetSubscriptionPlan(planID)
}

func (s *SubscriptionPlanService) ListSubscriptionPlans(page, pageSize int) (*repository.PaginatedSubscriptionPlans, error) {
	return repository.FindAllSubscriptionPlans(page, pageSize)
}

func (s *SubscriptionPlanService) UpdateSubscriptionPlan(planID uuid.UUID, patch *repository.SubscriptionPlanPatch) (*models.SubscriptionPlan, error) {
	return repository.UpdateSubscriptionPlanPartial(planID, patch)
}

func (s *SubscriptionPlanService) DeleteSubscriptionPlan(planID uuid.UUID) error {
	return repository.DeleteSubscriptionPlanByID(planID)
}
