package controllers

import (
	"gin-funda-africa/internal/modules/users/dto"
	"gin-funda-africa/internal/modules/users/models"
	"gin-funda-africa/internal/modules/users/services"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	Log "github.com/sirupsen/logrus"
)

type AuthController struct {
	service *services.UserService
}

func NewAuthController(service *services.UserService) *AuthController {
	return &AuthController{service: service}
}

// @BasePath /api/v1

// Create godoc
// @Summary Register user
// @Description Create account with profile image
// @Tags user
// @Accept mpfd
// @Produce json
// @Param username formData string true "Username"
// @Param email formData string false "Email"
// @Param password formData string true "Password"
// @Param first_name formData string true "First Name"
// @Param middle_name formData string false "Middle Name"
// @Param last_name formData string true "Last Name"
// @Param role formData string true "Role"
// @Param date_of_birth formData string true "Date Of Birth"
// @Param gender formData string false "gender"
// @Param phone_number formData string false "Phone Number"
// @Param country_code formData string true "Country Code"
// @Param profile_picture formData file false "Profile Picture"
// @Success 201 {object} dto.UserResponse
// @Failure 400 {object} dto.BadRequestResponse  "Bad request, validation error"
// @Failure 500 {object} dto.InternalServerErrorResponse "Internal server error"
// @Router /users [post]
func (c *AuthController) Create(ctx *gin.Context) {
	if err := ctx.Request.ParseMultipartForm(10 << 20); err != nil { // 10MB max
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "failed to parse form data"})
		return
	}

	dobStr := ctx.PostForm("date_of_birth")
	var dob *time.Time
	if dobStr != "" {
		parsedDob, err := time.Parse("2006-01-02", dobStr) // expected format: YYYY-MM-DD
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid date_of_birth format, use YYYY-MM-DD"})
			return
		}
		dob = &parsedDob
	}

	user := models.User{
		Username:    ctx.PostForm("username"),
		Email:       ctx.PostForm("email"),
		Password:    ctx.PostForm("password"),
		FirstName:   ctx.PostForm("first_name"),
		LastName:    ctx.PostForm("last_name"),
		MiddleName:  ctx.PostForm("middle_name"),
		PhoneNumber: ctx.PostForm("phone_number"),
		Role:        ctx.PostForm("role"),
		CountryCode: ctx.PostForm("country_code"),
		DateOfBirth: dob,
		Gender:      ctx.PostForm("gender"),
	}

	file, header, err := ctx.Request.FormFile("profile_picture")
	if err == nil {
		defer file.Close()

		profileImage, err := c.service.SaveProfileImage(file, header)
		if err != nil {
			Log.Debug("Error saving image ", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save profile image"})
			return
		}

		user.ProfilePicture = profileImage
	}

	createdUser, err := c.service.CreateUser(&user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var resp dto.UserResponse
	err = copier.Copy(&resp, createdUser)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, resp)
}

// @BasePath /api/v1

// Create godoc
// @Summary Login
// @Description User Login
// @Tags user
// @Accept json
// @Produce json
// @Param credentials body dto.LoginRequest true "User login credentials"
// @Success 200 {object} dto.LoginResponse
// @Failure 400 {object} dto.BadRequestResponse  "Bad request, validation error"
// @Failure 401 {object} dto.UnauthorizedResponse  "Invalid Credentials"
// @Failure 500 {object} dto.InternalServerErrorResponse "Internal server error"
// @Router /users/login [post]
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

	var resp dto.UserResponse
	err = copier.Copy(&resp, user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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
		ctx.JSON(http.StatusOK, gin.H{"user": resp})
	} else {
		// Mobile: return bearer token
		ctx.JSON(http.StatusOK, gin.H{"user": resp, "token": token})
	}
}
