package services

import (
	"gin-funda-africa/internal/modules/users/models"
	"gin-funda-africa/internal/modules/users/repository"

	"github.com/google/uuid"
)

type UserSubscriptionService struct{}

func NewUserSubscriptionService() *UserSubscriptionService {
	return &UserSubscriptionService{}
}

func (s *UserSubscriptionService) CreateUserSubscription(sub *models.UserSubscription) (*models.UserSubscription, error) {
	return repository.SaveUserSubscription(sub)
}

func (s *UserSubscriptionService) GetUserSubscription(userID, planID uuid.UUID) (*models.UserSubscription, error) {
	return repository.GetUserSubscription(userID, planID)
}

func (s *UserSubscriptionService) ListUserSubscriptions(page, pageSize int) (*repository.PaginatedUserSubscriptions, error) {
	return repository.FindAllUserSubscriptions(page, pageSize)
}

func (s *UserSubscriptionService) UpdateUserSubscriptionPartial(userID, planID uuid.UUID, patch *repository.UserSubscriptionPatch) (*models.UserSubscription, error) {
	return repository.UpdateUserSubscriptionPartial(userID, planID, patch)
}
