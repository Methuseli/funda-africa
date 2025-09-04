package controllers

import (
	"errors"
	"gin-funda-africa/internal/modules/organizations/dto"
	"gin-funda-africa/internal/modules/organizations/organization_models"
	"gin-funda-africa/internal/modules/organizations/repository"
	"gin-funda-africa/internal/modules/organizations/services"
	"gin-funda-africa/internal/pkg/middleware"
	"gin-funda-africa/internal/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
)

type OrganizationSubscriptionController struct {
	service *services.OrganizationSubscriptionService
}

func NewOrganizationSubscriptionController(s *services.OrganizationSubscriptionService) *OrganizationSubscriptionController {
	return &OrganizationSubscriptionController{service: s}
}

func checkOwnership(param uuid.UUID, ctx *gin.Context, service services.OrganizationService) (error) {
	org, err := service.GetOrganization(param)
	if err != nil {
		return errors.New("Organization does not exist")
	}

	userID, _ := middleware.GetCurrentUserID(ctx)
	userRole, _ := middleware.GetCurrentUserRole(ctx)

	if userID != &org.UserID {
		if userRole != nil && !(*userRole == "super_admin" || *userRole == "admin") {
			return errors.New(forbidden)
		}
	}

	return nil
}



// @BasePath /api/v1

// Create godoc
// @Security CookieAuth
// @Summary Add Organization Subscription Plan
// @Description Organization Subscription plan management
// @Tags organization-subscription
// @Accept json
// @Produce json
// @Param subscription_plan body dto.OrganizationSubscriptionResponse true "Organization Subscription Plan"
// @Success 201 {object} dto.OrganizationSubscriptionResponse
// @Failure 400 {object} dto.BadRequestResponse  "Bad request, validation error"
// @Failure 401 {object} dto.UnauthorizedResponse  "Unauthorized, authentication required"
// @Failure 403 {object} dto.ForbiddenResponse  "Forbidden"
// @Failure 500 {object} dto.InternalServerErrorResponse "Internal server error"
// @Router /organization-subscription/ [post]
func (c *OrganizationSubscriptionController) Create(ctx *gin.Context) {
	var sub organization_models.OrganizationSubscription
	if err := ctx.ShouldBindJSON(&sub); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	OrganizationSubscription, err := c.service.CreateOrganizationSubscription(&sub)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var resp dto.OrganizationSubscriptionResponse
	err = copier.Copy(&resp, OrganizationSubscription)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, resp)
}

// @BasePath /api/v1/

// @Security CookieAuth
// @Summary Retrieve Organization Subscription
// @Description Get a single Organization subscription object by their unique UUID
// @Tags organization-subscription
// @Accept json
// @Produce json
// @Param organization_id path string true "Organization UUID"
// @Param plan_id path string true "Subscription plan UUID"
// @Success 200 {object} dto.OrganizationSubscriptionResponse
// @Failure 400 {object} dto.BadRequestResponse "Bad request, invalid UUID"
// @Failure 401 {object} dto.UnauthorizedResponse "Unauthorized"
// @Failure 403 {object} dto.ForbiddenResponse "Forbidden"
// @Failure 404 {object} dto.NotFoundResponse "Organization Subscription not found"
// @Failure 500 {object} dto.InternalServerErrorResponse "Internal server error"
// @Router /organization-subscriptions/{organization_id}/{plan_id} [get]
func (c *OrganizationSubscriptionController) Get(ctx *gin.Context) {
	organizationID, err := uuid.Parse(ctx.Param("organization_id"))
	planID, err := uuid.Parse(ctx.Param("plan_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": utils.InvalidUUID})
		return
	}

	err = checkOwnership(organizationID, ctx, *services.NewOrganizationService())
	if err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	sub, err := c.service.GetOrganizationSubscription(organizationID, planID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	var resp dto.OrganizationSubscriptionResponse
	err = copier.Copy(&resp, sub)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, sub)
}

// @BasePath /api/v1

// Create godoc
// @Security CookieAuth
// @Summary Retrieve all Organization subscriptions
// @Description Organization Subscription
// @Tags organization-subscriptions
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param pageSize query int false "Number of items per page" default(20)
// @Success 200 {array} dto.OrganizationSubscriptionResponse
// @Failure 400 {object} dto.BadRequestResponse  "Bad request, validation error"
// @Failure 401 {object} dto.UnauthorizedResponse  "Unauthorized, authentication required"
// @Failure 403 {object} dto.ForbiddenResponse  "Forbidden"
// @Failure 404 {object} dto.NotFoundResponse "Resource not found"
// @Failure 500 {object} dto.InternalServerErrorResponse "Internal server error"
// @Router /organization-subscriptions/ [get]
func (c *OrganizationSubscriptionController) List(ctx *gin.Context) {
	page := utils.GetQueryInt(ctx, "page", 1)
	pageSize := utils.GetQueryInt(ctx, "pageSize", 20)

	subs, err := c.service.ListOrganizationSubscriptions(page, pageSize)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, subs)
}

// @BasePath /api/v1

// @Security CookieAuth
// @Summary Partially Update Organization Subscription
// @Description Update selected fields of a Organization subscription by their UUID
// @Tags organization-subscription
// @Accept json
// @Produce json
// @Param Organization_id path string true "Organization UUID"
// @Param plan_id path string true "Subscription plan UUID"
// @Param Organization_subscription_patch body repository.OrganizationSubscriptionPatch true "Organization fields to update"
// @Success 200 {object} dto.OrganizationSubscriptionResponse
// @Failure 400 {object} dto.BadRequestResponse "Bad request, invalid UUID or payload"
// @Failure 401 {object} dto.UnauthorizedResponse "Unauthorized"
// @Failure 403 {object} dto.ForbiddenResponse "Forbidden"
// @Failure 500 {object} dto.InternalServerErrorResponse "Internal server error"
// @Router /organization-subscriptions/{id} [patch]
func (c *OrganizationSubscriptionController) Update(ctx *gin.Context) {
	organizationID, err := uuid.Parse(ctx.Param("Organization_id"))
	planID, err := uuid.Parse(ctx.Param("plan_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": utils.InvalidUUID})
		return
	}

	err = checkOwnership(organizationID, ctx, *services.NewOrganizationService())
	if err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	var patch repository.OrganizationSubscriptionPatch
	if err := ctx.ShouldBindJSON(&patch); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	sub, err := c.service.UpdateOrganizationSubscriptionPartial(organizationID, planID, &patch)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var resp dto.OrganizationSubscriptionResponse
	err = copier.Copy(&resp, sub)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	ctx.JSON(http.StatusOK, sub)
}

