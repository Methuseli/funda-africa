package repository

import (
	"gin-funda-africa/internal/modules/organizations/dto"
	"gin-funda-africa/internal/modules/organizations/organization_models"
	"gin-funda-africa/internal/modules/users"
	"gin-funda-africa/internal/pkg/utils"

	"github.com/google/uuid"
	"github.com/jinzhu/copier"
)

type PaginatedOrganizations struct {
    Data       []*dto.OrganizationResponse `json:"data"`
    TotalCount int64 `json:"total_count"`
    Page       int `json:"page"`
    PageSize   int `json:"page_size"`
}

type OrganizationPatch struct {
	Name        *string `json:"name"`
    Type        *string `json:"type"`
    CountryCode *string `json:"country_code"`
    City        *string `json:"city"`
}

func SaveOrganization(organization *organization_models.Organization) (*organization_models.Organization, error) {
	err := users.Database.Create(&organization). Error
	if err != nil {
		return &organization_models.Organization{}, err
	}

	return organization, nil
}

func FindOrganizationsByUserID(userID uuid.UUID, page, pageSize int) (*PaginatedOrganizations, error) {
    var organizations []*organization_models.Organization
    var totalCount int64

    // First, get the total count (for frontend pagination)
    countQuery := users.Database.
        Model(&organization_models.Organization{}).
        Joins("JOIN user_organizations uo ON uo.organization_id = organizations.id").
        Where("uo.user_id = ?", userID)

    if err := countQuery.Count(&totalCount).Error; err != nil {
        return nil, err
    }

    // Pagination calculation
    if page < 1 {
        page = 1
    }
    offset := (page - 1) * pageSize

    // Fetch the actual page of data
    query := users.Database.
        Model(&organization_models.Organization{}).
        Joins("JOIN user_organizations uo ON uo.organization_id = organizations.id").
        Where("uo.user_id = ?", userID).
        Limit(pageSize).
        Offset(offset)

    if err := query.Find(&organizations).Error; err != nil {
        return nil, err
    }

    var responses []*dto.OrganizationResponse
	if err := copier.Copy(&responses, &organizations); err != nil {
		return nil, err
	}

    return &PaginatedOrganizations{
        Data:       responses,
        TotalCount: totalCount,
        Page:       page,
        PageSize:   pageSize,
    }, nil
}

func FindAllOrganizations(page, pageSize int) (*PaginatedOrganizations, error) {
    var organizations []*organization_models.Organization
    var totalCount int64

    // Count total organizations
    if err := users.Database.
        Model(&organization_models.Organization{}).
        Count(&totalCount).Error; err != nil {
        return nil, err
    }

    // Default page handling
    if page < 1 {
        page = 1
    }
    offset := (page - 1) * pageSize

    // Fetch paginated results
    if err := users.Database.
        Model(&organization_models.Organization{}).
        Limit(pageSize).
        Offset(offset).
        Find(&organizations).Error; err != nil {
        return nil, err
    }

    var responses []*dto.OrganizationResponse
	if err := copier.Copy(&responses, &organizations); err != nil {
		return nil, err
	}

    return &PaginatedOrganizations{
        Data:       responses,
        TotalCount: totalCount,
        Page:       page,
        PageSize:   pageSize,
    }, nil
}

func UpdateOrganization(id uuid.UUID, patch *OrganizationPatch) (*organization_models.Organization, error) {
	updates := utils.StructToMap(patch)
    if len(updates) == 0 {
        return nil, nil // nothing to update
    }

    var org organization_models.Organization
    if err := users.Database.Model(&org).Where(utils.IdQuery, id).Updates(updates).Error; err != nil {
        return nil, err
    }

    if err := users.Database.First(&org, utils.IdQuery, id).Error; err != nil {
        return nil, err
    }

    return &org, nil
}

func DeleteOrganization(id uuid.UUID) error {
	var org organization_models.Organization
    if err := users.Database.First(&org, utils.IdQuery, id).Error; err != nil {
        return err // organization not found or DB error
    }

    if err := users.Database.Delete(&org).Error; err != nil {
        return err
    }

    return nil

}

func GetOrganization(id uuid.UUID) (*organization_models.Organization, error) {
    var org organization_models.Organization
    if err := users.Database.First(&org, utils.IdQuery, id).Error; err != nil {
        return nil, err
    }
    return &org, nil
}
