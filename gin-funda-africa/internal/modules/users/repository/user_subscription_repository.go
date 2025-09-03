package repository

import (
	"gin-funda-africa/internal/modules/users"
	"gin-funda-africa/internal/modules/users/models"
	"gin-funda-africa/internal/pkg/utils"
	"time"

	"github.com/google/uuid"
)

const planQuery = "user_id = ? AND plan_id = ?"

// ---------- CREATE ----------
func SaveUserSubscription(sub *models.UserSubscription) (*models.UserSubscription, error) {
	if sub.UserID == uuid.Nil {
		sub.UserID = uuid.New()
	}
	sub.CreatedAt = time.Now()
	sub.UpdatedAt = time.Now()

	if err := users.Database.Create(&sub).Error; err != nil {
		return nil, err
	}
	return sub, nil
}

// ---------- RETRIEVE SINGLE ----------
func GetUserSubscription(userID uuid.UUID, planID uint) (*models.UserSubscription, error) {
	var sub models.UserSubscription
	if err := users.Database.
		Where(planQuery, userID, planID).
		First(&sub).Error; err != nil {
		return nil, err
	}
	return &sub, nil
}

// ---------- PAGINATED LIST ----------
type PaginatedUserSubscriptions struct {
	Data       []*models.UserSubscription `json:"data"`
	TotalCount int64                      `json:"total_count"`
	Page       int                        `json:"page"`
	PageSize   int                        `json:"page_size"`
}

func FindAllUserSubscriptions(page, pageSize int) (*PaginatedUserSubscriptions, error) {
	var list []*models.UserSubscription
	var totalCount int64

	if err := users.Database.Model(&models.UserSubscription{}).Count(&totalCount).Error; err != nil {
		return nil, err
	}

	if page < 1 {
		page = 1
	}
	offset := (page - 1) * pageSize

	if err := users.Database.Limit(pageSize).Offset(offset).Find(&list).Error; err != nil {
		return nil, err
	}

	return &PaginatedUserSubscriptions{
		Data:       list,
		TotalCount: totalCount,
		Page:       page,
		PageSize:   pageSize,
	}, nil
}

// ---------- PARTIAL UPDATE ----------
type UserSubscriptionPatch struct {
	StartDate *time.Time
	EndDate   **time.Time
	IsActive  *bool
}

func UpdateUserSubscriptionPartial(userID uuid.UUID, planID uint, patch *UserSubscriptionPatch) (*models.UserSubscription, error) {
	updates := utils.StructToMap(patch)

	if len(updates) == 0 {
		return nil, nil
	}

	updates["updated_at"] = time.Now()

	var sub models.UserSubscription
	if err := users.Database.Model(&sub).
		Where(planQuery, userID, planID).
		Updates(updates).Error; err != nil {
		return nil, err
	}

	if err := users.Database.
		Where(planQuery, userID, planID).
		First(&sub).Error; err != nil {
		return nil, err
	}

	return &sub, nil
}

// ---------- DELETE ----------
func DeleteUserSubscription(userID uuid.UUID, planID uint) error {
	var sub models.UserSubscription
	if err := users.Database.
		Where(planQuery, userID, planID).
		First(&sub).Error; err != nil {
		return err
	}
	return users.Database.Delete(&sub).Error
}
