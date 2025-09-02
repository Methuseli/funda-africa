package controllers

import (
	"gin-funda-africa/internal/modules/users/dto"
	"gin-funda-africa/internal/modules/users/repository"
	"gin-funda-africa/internal/modules/users/services"
	"gin-funda-africa/internal/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
)

type UserController struct {
	service *services.UserService
	jwtKey []byte
}

func NewUserController(service *services.UserService) *UserController {
	return &UserController{service: service}
}

// @BasePath /api/v1

// @Security CookieAuth
// @Summary Retrieve Single User
// @Description Get a single user by their unique UUID
// @Tags user
// @Accept json
// @Produce json
// @Param id path string true "User UUID"
// @Success 200 {object} dto.UserResponse
// @Failure 400 {object} dto.BadRequestResponse "Bad request, invalid UUID"
// @Failure 401 {object} dto.UnauthorizedResponse "Unauthorized"
// @Failure 403 {object} dto.ForbiddenResponse "Forbidden"
// @Failure 404 {object} dto.NotFoundResponse "User not found"
// @Failure 500 {object} dto.InternalServerErrorResponse "Internal server error"
// @Router /users/{id} [get]
func (c *UserController) Get(ctx *gin.Context) {
	idParam := ctx.Param("id")
	userID, err := uuid.Parse(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": utils.InvalidUUID})
		return
	}

	user, err := c.service.GetUserByID(userID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	var resp dto.UserResponse
	err = copier.Copy(&resp, user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}


// @BasePath /api/v1

// @Security CookieAuth
// @Summary Partially Update User
// @Description Update selected fields of a user by their UUID
// @Tags user
// @Accept json
// @Produce json
// @Param id path string true "User UUID"
// @Param user_patch body repository.UserPatch true "User fields to update"
// @Success 200 {object} dto.UserResponse
// @Failure 400 {object} dto.BadRequestResponse "Bad request, invalid UUID or payload"
// @Failure 401 {object} dto.UnauthorizedResponse "Unauthorized"
// @Failure 403 {object} dto.ForbiddenResponse "Forbidden"
// @Failure 500 {object} dto.InternalServerErrorResponse "Internal server error"
// @Router /users/{id} [patch]
func (c *UserController) Update(ctx *gin.Context) {
	idParam := ctx.Param("id")
	userID, err := uuid.Parse(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": utils.InvalidUUID})
		return
	}

	var patch repository.UserPatch
	if err := ctx.ShouldBindJSON(&patch); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedUser, err := c.service.UpdateUserPartial(userID, &patch)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var resp dto.UserResponse
	err = copier.Copy(&resp, updatedUser)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

// @BasePath /api/v1

// @Security CookieAuth
// @Summary Delete User
// @Description Permanently delete a user by their UUID
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User UUID"
// @Success 204 "No Content"
// @Failure 400 {object} dto.BadRequestResponse "Bad request, invalid UUID"
// @Failure 401 {object} dto.UnauthorizedResponse "Unauthorized"
// @Failure 403 {object} dto.ForbiddenResponse "Forbidden"
// @Failure 500 {object} dto.InternalServerErrorResponse "Internal server error"
// @Router /users/{id} [delete]
func (c *UserController) Delete(ctx *gin.Context) {
	idParam := ctx.Param("id")
	userID, err := uuid.Parse(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": utils.InvalidUUID})
		return
	}

	if err := c.service.DeleteUser(userID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}


