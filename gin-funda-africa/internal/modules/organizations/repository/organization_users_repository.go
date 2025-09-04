package repository

import (
	"gin-funda-africa/internal/modules/organizations/dto"
	"gin-funda-africa/internal/modules/organizations/organization_models"
	"gin-funda-africa/internal/modules/users"
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/copier"
)

const queryIds = "user_id = ? AND organization_id = ?"

// ---------- CREATE ----------
func SaveOrganizationUser(uo *organization_models.OrganizationUser) (*organization_models.OrganizationUser, error) {
	if err := users.Database.Create(&uo).Error; err != nil {
		return nil, err
	}
	return uo, nil
}

// ---------- RETRIEVE SINGLE ----------
func GetOrganizationUser(userID, orgID uuid.UUID) (*organization_models.OrganizationUser, error) {
	var uo organization_models.OrganizationUser
	if err := users.Database.
		Where(queryIds, userID, orgID).
		First(&uo).Error; err != nil {
		return nil, err
	}
	return &uo, nil
}

// ---------- PAGINATED LIST ----------
type PaginatedOrganizationUsers struct {
	Data       []*dto.OrganizationUserResponse `json:"data"`
	TotalCount int64                           `json:"total_count"`
	Page       int                             `json:"page"`
	PageSize   int                             `json:"page_size"`
}

func FindByOrganization(id uuid.UUID, page, pageSize int) (*PaginatedOrganizationUsers, error) {
	var list []*organization_models.OrganizationUser
	var totalCount int64

	if err := users.Database.Model(&organization_models.OrganizationUser{}).Where("organization_id = ?", id).Count(&totalCount).Error; err != nil {
		return nil, err
	}

	if page < 1 {
		page = 1
	}
	offset := (page - 1) * pageSize

	if err := users.Database.Model(&organization_models.OrganizationUser{}).Where("organization_id = ?", id).Limit(pageSize).Offset(offset).Find(&list).Error; err != nil {
		return nil, err
	}

	var responses []*dto.OrganizationUserResponse
	if err := copier.Copy(&responses, &list); err != nil {
		return nil, err
	}

	return &PaginatedOrganizationUsers{
		Data:       responses,
		TotalCount: totalCount,
		Page:       page,
		PageSize:   pageSize,
	}, nil
}

// ---------- PARTIAL UPDATE ----------
type OrganizationUserPatch struct {
	Position *string
	IsActive bool
}

func UpdateOrganizationUser(userID, orgID uuid.UUID, patch *OrganizationUserPatch) (*organization_models.OrganizationUser, error) {
	updates := make(map[string]interface{})
	if patch.Position != nil {
		updates["position"] = *patch.Position
	}
	if len(updates) == 0 {
		return nil, nil
	}

	updates["updated_at"] = time.Now()

	var uo organization_models.OrganizationUser
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
