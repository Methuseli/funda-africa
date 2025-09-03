package repository

import (
	"gin-funda-africa/internal/modules/users"
	"gin-funda-africa/internal/modules/users/models"
	"time"

	"github.com/google/uuid"
)

const queryIds = "user_id = ? AND organization_id = ?"

// ---------- CREATE ----------
func SaveUserOrganization(uo *models.UserOrganization) (*models.UserOrganization, error) {
	if err := users.Database.Create(&uo).Error; err != nil {
		return nil, err
	}
	return uo, nil
}

// ---------- RETRIEVE SINGLE ----------
func GetUserOrganization(userID, orgID uuid.UUID) (*models.UserOrganization, error) {
	var uo models.UserOrganization
	if err := users.Database.
		Where(queryIds, userID, orgID).
		First(&uo).Error; err != nil {
		return nil, err
	}
	return &uo, nil
}

// ---------- PAGINATED LIST ----------
type PaginatedUserOrganizations struct {
	Data       []*models.UserOrganization `json:"data"`
	TotalCount int64                      `json:"total_count"`
	Page       int                        `json:"page"`
	PageSize   int                        `json:"page_size"`
}

func FindAllUserOrganizations(page, pageSize int) (*PaginatedUserOrganizations, error) {
	var list []*models.UserOrganization
	var totalCount int64

	if err := users.Database.Model(&models.UserOrganization{}).Count(&totalCount).Error; err != nil {
		return nil, err
	}

	if page < 1 {
		page = 1
	}
	offset := (page - 1) * pageSize

	if err := users.Database.Limit(pageSize).Offset(offset).Find(&list).Error; err != nil {
		return nil, err
	}

	return &PaginatedUserOrganizations{
		Data:       list,
		TotalCount: totalCount,
		Page:       page,
		PageSize:   pageSize,
	}, nil
}

// ---------- PARTIAL UPDATE ----------
type UserOrganizationPatch struct {
	Position *string
}

func UpdateUserOrganization(userID, orgID uuid.UUID, patch *UserOrganizationPatch) (*models.UserOrganization, error) {
	updates := make(map[string]interface{})
	if patch.Position != nil {
		updates["position"] = *patch.Position
	}
	if len(updates) == 0 {
		return nil, nil
	}

	updates["updated_at"] = time.Now()

	var uo models.UserOrganization
	if err := users.Database.Model(&uo).
		Where(queryIds, userID, orgID).
		Updates(updates).Error; err != nil {
		return nil, err
	}

	if err := users.Database.
		Where(queryIds, userID, orgID).
		First(&uo).Error; err != nil {
		return nil, err
	}

	return &uo, nil
}

// ---------- DELETE ----------
func DeleteUserOrganization(userID, orgID uuid.UUID) error {
	var uo models.UserOrganization
	if err := users.Database.
		Where(queryIds, userID, orgID).
		First(&uo).Error; err != nil {
		return err
	}
	return users.Database.Delete(&uo).Error
}
