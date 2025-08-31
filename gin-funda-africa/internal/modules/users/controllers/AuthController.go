package controllers

import (
	"gin-funda-africa/internal/modules/users/models"
	"gin-funda-africa/internal/modules/users/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	service *services.UserService
}

func NewAuthController(service *services.UserService) *AuthController {
	return &AuthController{service: service}
}

// @BasePath /api/v1

// PingExample godoc
// @Summary Register user
// @Schemes
// @Description Create account
// @Tags user
// @Accept mpfd
// @Produce json
// @Param user formData models.User true "User Data"
// @Success 201 {object} models.User
// @Router /users [post]
func (c *AuthController) Create(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdUser, err := c.service.CreateUser(&user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, createdUser)
}

// Login user
func (c *AuthController) Login(ctx *gin.Context) {
	var credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Client   string `json:"client"`
	}

	if err := ctx.ShouldBindJSON(&credentials); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, token, err := c.service.Login(credentials.Username, credentials.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	if credentials.Client == "web" {
		// Set HttpOnly cookie
		ctx.SetCookie(
			"session_token",
			token,
			3600*24, // 24 hours
			"/",
			"",
			false, // secure = false for dev, true in prod with HTTPS
			true,  // HttpOnly
		)
		ctx.JSON(http.StatusOK, gin.H{"user": user})
	} else {
		// Mobile: return bearer token
		ctx.JSON(http.StatusOK, gin.H{"user": user, "token": token})
	}
}
