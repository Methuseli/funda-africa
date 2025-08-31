package repository

import (
	"fmt"
	"gin-funda-africa/internal/modules/users"
	"gin-funda-africa/internal/modules/users/models"
	"gin-funda-africa/internal/pkg/utils"
	"time"

	"github.com/google/uuid"
)

// ---------- CREATE ----------
func SaveUser(user *models.User) (*models.User, error) {
	return user.Save()
}

// ---------- RETRIEVE SINGLE ----------
func GetUser(id uuid.UUID) (*models.User, error) {
	var user models.User
	if err := users.Database.First(&user, utils.IdQuery, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	if err := users.Database.First(&user, "username = ?", username).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUserForLogin(username, password string) (*models.User, error) {
	user, err := GetUserByUsername(username)
	if err != nil {
		return nil, err
	}

	if !user.ValidatePassword(password) {
		return nil, fmt.Errorf("invalid credentials")
	}

	return user, nil
}

// ---------- PAGINATED LIST ----------
type PaginatedUsers struct {
	Data       []*models.User
	TotalCount int64
	Page       int
	PageSize   int
}

func FindAllUsers(page, pageSize int) (*PaginatedUsers, error) {
	var usersList []*models.User
	var totalCount int64

	if err := users.Database.Model(&models.User{}).Count(&totalCount).Error; err != nil {
		return nil, err
	}

	if page < 1 {
		page = 1
	}
	offset := (page - 1) * pageSize

	if err := users.Database.Limit(pageSize).Offset(offset).Find(&usersList).Error; err != nil {
		return nil, err
	}

	return &PaginatedUsers{
		Data:       usersList,
		TotalCount: totalCount,
		Page:       page,
		PageSize:   pageSize,
	}, nil
}

// ---------- PARTIAL UPDATE ----------
type UserPatch struct {
	Username       *string
	Email          *string
	Password       *string
	PhoneNumber    *string
	FirstName      *string
	LastName       *string
	MiddleName     *string
	ProfilePicture *string
	DateOfBirth    **time.Time
	Gender         **string
	CountryCode    *string
}

func UpdateUser(id uuid.UUID, patch *UserPatch) (*models.User, error) {
	updates := utils.StructToMap(patch)
	
	if len(updates) == 0 {
		return nil, nil
	}

	updates["updated_at"] = time.Now()

	var user models.User
	if err := users.Database.Model(&user).Where(utils.IdQuery, id).Updates(updates).Error; err != nil {
		return nil, err
	}

	if err := users.Database.First(&user, utils.IdQuery, id).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

// ---------- DELETE ----------
func DeleteUser(id uuid.UUID) error {
	var user models.User
	if err := users.Database.First(&user, utils.IdQuery, id).Error; err != nil {
		return err
	}
	return users.Database.Delete(&user).Error
}
