package services

import (
	"errors"
	"gin-funda-africa/internal/modules/users/models"
	"gin-funda-africa/internal/modules/users/repository"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type UserService struct{
	jwtKey []byte
}

type Claims struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}


func NewUserService(jwtKey []byte) *UserService {
	return &UserService{jwtKey: jwtKey}
}

// Create a new user
func (s *UserService) CreateUser(user *models.User) (*models.User, error) {
	// Add business logic if needed (e.g., hashing password, validation)
	return repository.SaveUser(user)
}

// Get a single user by ID
func (s *UserService) GetUserByID(userID uuid.UUID) (*models.User, error) {
	return repository.GetUser(userID)
}

// List users with pagination
func (s *UserService) ListUsers(page, pageSize int) (*repository.PaginatedUsers, error) {
	return repository.FindAllUsers(page, pageSize)
}

// Partial update
func (s *UserService) UpdateUserPartial(userID uuid.UUID, patch *repository.UserPatch) (*models.User, error) {
	return repository.UpdateUser(userID, patch)
}

// Delete user
func (s *UserService) DeleteUser(userID uuid.UUID) error {
	return repository.DeleteUser(userID)
}

func (s *UserService) Login(username, password string) (*models.User, string, error) {
	user, err := repository.GetUserForLogin(username, password)
	if err != nil {
		return nil, "", errors.New("Invalid Credentials")
	}

	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		UserID: user.ID.String(),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(s.jwtKey)
	if err != nil {
		return nil, "", err
	}

	return user, tokenString, nil
}

