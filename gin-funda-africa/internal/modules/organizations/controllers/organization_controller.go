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

type OrganizationController struct {
	service *services.OrganizationService
}

func NewOrganizationController(s *services.OrganizationService) *OrganizationController {
	return &OrganizationController{service: s}
}

// @BasePath /api/v1

// Create godoc
// @Security CookieAuth
// @Summary Add Organizations
// @Description Organizations management
// @Tags organization
// @Accept json
// @Produce json
// @Param subscription_plan body dto.OrganizationResponse true "Organizations"
// @Success 201 {object} dto.OrganizationResponse
// @Failure 400 {object} dto.BadRequestResponse  "Bad request, validation error"
// @Failure 401 {object} dto.UnauthorizedResponse  "Unauthorized, authentication required"
// @Failure 403 {object} dto.ForbiddenResponse  "Forbidden"
// @Failure 500 {object} dto.InternalServerErrorResponse "Internal server error"
// @Router /organizations/ [post]
func (c *OrganizationController) Create(ctx *gin.Context) {
	var org organization_models.Organization
	if err := ctx.ShouldBindJSON(&org); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := c.service.CreateOrganization(&org)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var resp dto.OrganizationResponse
	err = copier.Copy(&resp, result)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, result)
}

// @BasePath /api/v1/

// @Security CookieAuth
// @Summary Retrieve organization
// @Description Get a single organization object by their unique UUID
// @Tags organizations
// @Accept json
// @Produce json
// @Param id path string true "Organization UUID"
// @Success 200 {object} dto.OrganizationResponse
// @Failure 400 {object} dto.BadRequestResponse "Bad request, invalid UUID"
// @Failure 401 {object} dto.UnauthorizedResponse "Unauthorized"
// @Failure 403 {object} dto.ForbiddenResponse "Forbidden"
// @Failure 404 {object} dto.NotFoundResponse "organization not found"
// @Failure 500 {object} dto.InternalServerErrorResponse "Internal server error"
// @Router /organizations/{id} [get]
func (c *OrganizationController) Get(ctx *gin.Context) {
	id, _ := uuid.Parse(ctx.Param("id"))
	org, err := c.service.GetOrganization(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	var resp dto.OrganizationResponse
	err = copier.Copy(&resp, org)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, org)
}

// @BasePath /api/v1

// Create godoc
// @Security CookieAuth
// @Summary Retrieve all organizations
// @Description Organizations
// @Tags organizations
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param pageSize query int false "Number of items per page" default(20)
// @Success 200 {array} dto.OrganizationResponse
// @Failure 400 {object} dto.BadRequestResponse  "Bad request, validation error"
// @Failure 401 {object} dto.UnauthorizedResponse  "Unauthorized, authentication required"
// @Failure 403 {object} dto.ForbiddenResponse  "Forbidden"
// @Failure 404 {object} dto.NotFoundResponse "Resource not found"
// @Failure 500 {object} dto.InternalServerErrorResponse "Internal server error"
// @Router /organizations/ [get]
func (c *OrganizationController) List(ctx *gin.Context) {
	page := utils.GetQueryInt(ctx, "page", 1)
	pageSize := utils.GetQueryInt(ctx, "pageSize", 10)
	orgs, err := c.service.ListOrganizations(page, pageSize)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, orgs)
}

// @BasePath /api/v1

// Create godoc
// @Security CookieAuth
// @Summary Retrieve all organizations by user
// @Description Organizations
// @Tags organizations
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param pageSize query int false "Number of items per page" default(20)
// @Param user_id query string true "User UUID"
// @Success 200 {array} dto.OrganizationResponse
// @Failure 400 {object} dto.BadRequestResponse  "Bad request, validation error"
// @Failure 401 {object} dto.UnauthorizedResponse  "Unauthorized, authentication required"
// @Failure 403 {object} dto.ForbiddenResponse  "Forbidden"
// @Failure 404 {object} dto.NotFoundResponse "Resource not found"
// @Failure 500 {object} dto.InternalServerErrorResponse "Internal server error"
// @Router /organizations/{user_id} [get]
func (c *OrganizationController) ListByUser(ctx *gin.Context) {
	page := utils.GetQueryInt(ctx, "page", 1)
	pageSize := utils.GetQueryInt(ctx, "pageSize", 10)
	userID, err := uuid.Parse(ctx.Param("user_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": utils.InvalidUUID})
		return
	}

	orgs, err := c.service.ListUserOrganizations(page, pageSize, userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, orgs)
}

// @BasePath /api/v1

// @Security CookieAuth
// @Summary Partially Update Organization
// @Description Update selected fields of a organization by their UUID
// @Tags organization
// @Accept json
// @Produce json
// @Param id path string true "Organization UUID"
// @Param organization_patch body repository.OrganizationPatch true "User fields to update"
// @Success 200 {object} dto.OrganizationResponse
// @Failure 400 {object} dto.BadRequestResponse "Bad request, invalid UUID or payload"
// @Failure 401 {object} dto.UnauthorizedResponse "Unauthorized"
// @Failure 403 {object} dto.ForbiddenResponse "Forbidden"
// @Failure 500 {object} dto.InternalServerErrorResponse "Internal server error"
// @Router /organizations/{id} [patch]
func (c *OrganizationController) Update(ctx *gin.Context) {
	id, _ := uuid.Parse(ctx.Param("id"))
	var patch repository.OrganizationPatch
	if err := ctx.ShouldBindJSON(&patch); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	org, err := c.service.UpdateOrganization(id, &patch)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var resp dto.OrganizationResponse
	err = copier.Copy(&resp, org)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, org)
}

// @BasePath /api/v1

// @Security CookieAuth
// @Summary Delete Organization
// @Description Permanently delete an organization by their UUID
// @Tags organization
// @Accept json
// @Produce json
// @Param id path string true "Organization UUID"
// @Success 204 "No Content"
// @Failure 400 {object} dto.BadRequestResponse "Bad request, invalid UUID"
// @Failure 401 {object} dto.UnauthorizedResponse "Unauthorized"
// @Failure 403 {object} dto.ForbiddenResponse "Forbidden"
// @Failure 500 {object} dto.InternalServerErrorResponse "Internal server error"
// @Router /organizations/{id} [delete]
func (c *OrganizationController) Delete(ctx *gin.Context) {
	id, _ := uuid.Parse(ctx.Param("id"))
	if err := c.service.DeleteOrganization(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.Status(http.StatusNoContent)
}


