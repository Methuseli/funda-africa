package repository

import (
	"gin-funda-africa/internal/modules/users"
	"gin-funda-africa/internal/modules/users/models"
	"gin-funda-africa/internal/pkg/utils"

	"github.com/google/uuid"
)

// ---------- CREATE ----------
func SaveSubscriptionPlan(plan *models.SubscriptionPlan) (*models.SubscriptionPlan, error) {
	if err := users.Database.Create(&plan).Error; err != nil {
		return nil, err
	}

	return plan, nil
}

// ---------- RETRIEVE SINGLE ----------
func GetSubscriptionPlan(planID uuid.UUID) (*models.SubscriptionPlan, error) {
	var plan models.SubscriptionPlan
	if err := users.Database.First(&plan, utils.IdQuery, planID).Error; err != nil {
		return nil, err
	}
	return &plan, nil
}

// ---------- PAGINATED LIST ----------
type PaginatedSubscriptionPlans struct {
	Data       []*models.SubscriptionPlan
	TotalCount int64
	Page       int
	PageSize   int
}

func FindAllSubscriptionPlans(page, pageSize int) (*PaginatedSubscriptionPlans, error) {
	var plans []*models.SubscriptionPlan
	var totalCount int64

	if err := users.Database.Model(&models.SubscriptionPlan{}).Count(&totalCount).Error; err != nil {
		return nil, err
	}

	if page < 1 {
		page = 1
	}
	offset := (page - 1) * pageSize

	if err := users.Database.
		Limit(pageSize).
		Offset(offset).
		Find(&plans).Error; err != nil {
		return nil, err
	}

	return &PaginatedSubscriptionPlans{
		Data:       plans,
		TotalCount: totalCount,
		Page:       page,
		PageSize:   pageSize,
	}, nil
}

// ---------- PARTIAL UPDATE ----------
type SubscriptionPlanPatch struct {
	Name         *string
	Description  *string
	PriceUSD     *float64
	FeaturesJSON *[]byte
	IsActive     *bool
}

func UpdateSubscriptionPlanPartial(planID uuid.UUID, patch *SubscriptionPlanPatch) (*models.SubscriptionPlan, error) {
	updates := utils.StructToMap(patch)
	
	if len(updates) == 0 {
		return nil, nil
	}

	var plan models.SubscriptionPlan
	if err := users.Database.Model(&plan).Where(utils.IdQuery, planID).Updates(updates).Error; err != nil {
		return nil, err
	}

	if err := users.Database.First(&plan, utils.IdQuery, planID).Error; err != nil {
		return nil, err
	}

	return &plan, nil
}

// ---------- DELETE ----------
func DeleteSubscriptionPlanByID(planID uuid.UUID) error {
	var plan models.SubscriptionPlan
	if err := users.Database.First(&plan, utils.IdQuery, planID).Error; err != nil {
		return err
	}

	if err := users.Database.Delete(&plan).Error; err != nil {
		return err
	}
	return nil
}
