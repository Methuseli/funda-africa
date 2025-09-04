package repository

import (
	"gin-funda-africa/internal/modules/organizations/dto"
	"gin-funda-africa/internal/modules/organizations/organization_models"
	"gin-funda-africa/internal/modules/users"
	"gin-funda-africa/internal/pkg/utils"
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/copier"
)

const planQuery = "organization_id = ? AND plan_id = ?"

// ---------- CREATE ----------
func SaveOrganizationSubscription(sub *organization_models.OrganizationSubscription) (*organization_models.OrganizationSubscription, error) {
	if sub.OrganizationID == uuid.Nil {
		sub.OrganizationID = uuid.New()
	}
	sub.CreatedAt = time.Now()
	sub.UpdatedAt = time.Now()

	if err := users.Database.Create(&sub).Error; err != nil {
		return nil, err
	}
	return sub, nil
}

// ---------- RETRIEVE SINGLE ----------
func GetOrganizationSubscription(orgID, planID uuid.UUID) (*organization_models.OrganizationSubscription, error) {
	var sub organization_models.OrganizationSubscription
	if err := users.Database.
		Where(planQuery, orgID, planID).
		First(&sub).Error; err != nil {
		return nil, err
	}
	return &sub, nil
}

// ---------- PAGINATED LIST ----------
type PaginatedOrganizationSubscriptions struct {
	Data       []*dto.OrganizationSubscriptionResponse `json:"data"`
	TotalCount int64                              `json:"total_count"`
	Page       int                                `json:"page"`
	PageSize   int                                `json:"page_size"`
}

func FindAllOrganizationSubscriptions(page, pageSize int) (*PaginatedOrganizationSubscriptions, error) {
	var list []*organization_models.OrganizationSubscription
	var totalCount int64

	if err := users.Database.Model(&organization_models.OrganizationSubscription{}).Count(&totalCount).Error; err != nil {
		return nil, err
	}

	if page < 1 {
		page = 1
	}
	offset := (page - 1) * pageSize

	if err := users.Database.Limit(pageSize).Offset(offset).Find(&list).Error; err != nil {
		return nil, err
	}

	var organizationSubscriptionResponses []*dto.OrganizationSubscriptionResponse
	if err := copier.Copy(&organizationSubscriptionResponses, &list); err != nil {
		return nil, err
	}

	return &PaginatedOrganizationSubscriptions{
		Data:       organizationSubscriptionResponses,
		TotalCount: totalCount,
		Page:       page,
		PageSize:   pageSize,
	}, nil
}

// ---------- PARTIAL UPDATE ----------
type OrganizationSubscriptionPatch struct {
	StartDate *time.Time
	EndDate   *time.Time
	IsActive  *bool
}

func UpdateOrganizationSubscriptionPartial(orgID, planID uuid.UUID, patch *OrganizationSubscriptionPatch) (*organization_models.OrganizationSubscription, error) {
	updates := utils.StructToMap(patch)

	if len(updates) == 0 {
		return nil, nil
	}

	updates["updated_at"] = time.Now()

	var sub organization_models.OrganizationSubscription
	if err := users.Database.Model(&sub).
		Where(planQuery, orgID, planID).
		Updates(updates).Error; err != nil {
		return nil, err
	}

	if err := users.Database.
		Where(planQuery, orgID, planID).
		First(&sub).Error; err != nil {
		return nil, err
	}

	return &sub, nil
}
