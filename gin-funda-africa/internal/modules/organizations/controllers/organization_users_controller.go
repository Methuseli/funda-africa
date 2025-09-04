package controllers

import (
	"gin-funda-africa/internal/modules/organizations/dto"
	"gin-funda-africa/internal/modules/organizations/organization_models"
	"gin-funda-africa/internal/modules/organizations/repository"
	"gin-funda-africa/internal/modules/organizations/services"
	"gin-funda-africa/internal/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
)

type OrganizationUserController struct {
	service *services.OrganizationUserService
}

func NewOrganizationUserController(s *services.OrganizationUserService) *OrganizationUserController {
	return &OrganizationUserController{service: s}
}

// @BasePath /api/v1

// Create godoc
// @Security CookieAuth
// @Summary Add Organizations User
// @Description Organizations User management
// @Tags organization-user
// @Accept json
// @Produce json
// @Param subscription_plan body dto.OrganizationUserResponse true "Organizations User"
// @Success 201 {object} dto.OrganizationUserResponse
// @Failure 400 {object} dto.BadRequestResponse  "Bad request, validation error"
// @Failure 401 {object} dto.UnauthorizedResponse  "Unauthorized, authentication required"
// @Failure 403 {object} dto.ForbiddenResponse  "Forbidden"
// @Failure 500 {object} dto.InternalServerErrorResponse "Internal server error"
// @Router /organization-users/ [post]
func (c *OrganizationUserController) Create(ctx *gin.Context) {
	var uo organization_models.OrganizationUser
	if err := ctx.ShouldBindJSON(&uo); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := c.service.CreateOrganizationUser(&uo)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var resp dto.OrganizationUserResponse
	err = copier.Copy(&resp, result)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, result)
}

// @BasePath /api/v1/

// @Security CookieAuth
// @Summary Retrieve organization user
// @Description Get a single organization object by their unique UUID
// @Tags organization-user
// @Accept json
// @Produce json
// @Param id path string true "Organization UUID"
// @Success 200 {object} dto.OrganizationResponse
// @Failure 400 {object} dto.BadRequestResponse "Bad request, invalid UUID"
// @Failure 401 {object} dto.UnauthorizedResponse "Unauthorized"
// @Failure 403 {object} dto.ForbiddenResponse "Forbidden"
// @Failure 404 {object} dto.NotFoundResponse "organization not found"
// @Failure 500 {object} dto.InternalServerErrorResponse "Internal server error"
// @Router /organization-user/{id} [get]
func (c *OrganizationUserController) Get(ctx *gin.Context) {
	userID, _ := uuid.Parse(ctx.Param("user_"))
	orgID, _ := uuid.Parse(ctx.Param("orgID"))
	uo, err := c.service.GetOrganizationUser(userID, orgID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	var resp dto.OrganizationUserResponse
	err = copier.Copy(&resp, uo)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, uo)
}

// @BasePath /api/v1

// Create godoc
// @Security CookieAuth
// @Summary Retrieve users in an organization
// @Description Organization Users
// @Tags organization-users
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param pageSize query int false "Number of items per page" default(20)
// @Param organization_id query string true "Organization UUID"
// @Success 200 {array} dto.OrganizationUserResponse
// @Failure 400 {object} dto.BadRequestResponse  "Bad request, validation error"
// @Failure 401 {object} dto.UnauthorizedResponse  "Unauthorized, authentication required"
// @Failure 403 {object} dto.ForbiddenResponse  "Forbidden"
// @Failure 404 {object} dto.NotFoundResponse "Resource not found"
// @Failure 500 {object} dto.InternalServerErrorResponse "Internal server error"
// @Router /organization-users/ [get]
func (c *OrganizationUserController) List(ctx *gin.Context) {
	page := utils.GetQueryInt(ctx, "page", 1)
	pageSize := utils.GetQueryInt(ctx, "pageSize", 10)
	id, err := uuid.Parse(ctx.Param("organization_id"))
	uos, err := c.service.ListUserByOrganization(id, page, pageSize)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, uos)
}

// @BasePath /api/v1

// @Security CookieAuth
// @Summary Partially Update Organization User 
// @Description Update selected fields of a user subscription by their UUID
// @Tags organization-user
// @Accept json
// @Produce json
// @Param user_id path string true "User UUID"
// @Param org_id path string true "Organization UUID"
// @Param organization_user_patch body repository.OrganizationUserPatch true "User fields to update"
// @Success 200 {object} dto.OrganizationUserResponse
// @Failure 400 {object} dto.BadRequestResponse "Bad request, invalid UUID or payload"
// @Failure 401 {object} dto.UnauthorizedResponse "Unauthorized"
// @Failure 403 {object} dto.ForbiddenResponse "Forbidden"
// @Failure 500 {object} dto.InternalServerErrorResponse "Internal server error"
// @Router /organization-users/{id} [patch]
func (c *OrganizationUserController) Update(ctx *gin.Context) {
	userID, _ := uuid.Parse(ctx.Param("user_id"))
	orgID, _ := uuid.Parse(ctx.Param("org_id"))
	var patch repository.OrganizationUserPatch
	if err := ctx.ShouldBindJSON(&patch); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := c.service.UpdateOrganizationUserPartial(userID, orgID, &patch)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, result)
}

