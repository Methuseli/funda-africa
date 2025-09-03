package controllers

import (
	"gin-funda-africa/internal/modules/users/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AdminUserController struct {
	service *services.UserService
}

func NewAdminUserController(service *services.UserService) *AdminUserController {
	return &AdminUserController{service: service}
}

// @BasePath /api/v1

// Create godoc
// @Security CookieAuth
// @Summary Retrieve All Users
// @Description User management
// @Tags users
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param pageSize query int false "Number of items per page" default(10)
// @Success 200 {array} dto.UserResponse
// @Failure 400 {object} dto.BadRequestResponse  "Bad request, validation error"
// @Failure 401 {object} dto.UnauthorizedResponse  "Unauthorized, authentication required"
// @Failure 403 {object} dto.ForbiddenResponse  "Forbidden"
// @Failure 404 {object} dto.NotFoundResponse "Resource not found"
// @Failure 500 {object} dto.InternalServerErrorResponse "Internal server error"
// @Router /users/ [get]
func (c *AdminUserController) List(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "10"))

	users, err := c.service.ListUsers(page, pageSize)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, users)
}
