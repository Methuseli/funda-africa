package controllers

import (
	"gin-funda-africa/internal/modules/users/dto"
	"gin-funda-africa/internal/modules/users/models"
	"gin-funda-africa/internal/modules/users/repository"
	"gin-funda-africa/internal/modules/users/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
)

type SubscriptionPlanController struct {
	service *services.SubscriptionPlanService
}

func NewSubscriptionPlanController(s *services.SubscriptionPlanService) *SubscriptionPlanController {
	return &SubscriptionPlanController{service: s}
}


// @BasePath /api/v1

// Create godoc
// @Security CookieAuth
// @Summary Add Subscription Plan
// @Description Subscription plan management
// @Tags subscription plan
// @Accept json
// @Produce json
// @Param subscription_plan body dto.SubscriptionPlanResponse true "Subscription Plan"
// @Success 201 {object} dto.SubscriptionPlanResponse
// @Failure 400 {object} dto.BadRequestResponse  "Bad request, validation error"
// @Failure 401 {object} dto.UnauthorizedResponse  "Unauthorized, authentication required"
// @Failure 403 {object} dto.ForbiddenResponse  "Forbidden"
// @Failure 500 {object} dto.InternalServerErrorResponse "Internal server error"
// @Router /subscription-plans/ [post]
func (c *SubscriptionPlanController) Create(ctx *gin.Context) {
	var plan models.SubscriptionPlan
	if err := ctx.ShouldBindJSON(&plan); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := c.service.CreateSubscriptionPlan(&plan)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var resp dto.SubscriptionPlanResponse
	err = copier.Copy(&resp, result)
	ctx.JSON(http.StatusCreated, resp)
}

// @BasePath /api/v1

// Create godoc
// @Security CookieAuth
// @Summary Retrieve Single Subscription Plan
// @Description Subscription plan management
// @Tags subscription plan
// @Accept json
// @Produce json
// @Param id path string true "Subscription Plan ID"
// @Success 200 {object} dto.SubscriptionPlanResponse
// @Failure 400 {object} dto.BadRequestResponse  "Bad request, validation error"
// @Failure 401 {object} dto.UnauthorizedResponse  "Unauthorized, authentication required"
// @Failure 403 {object} dto.ForbiddenResponse  "Forbidden"
// @Failure 404 {object} dto.NotFoundResponse "Resource not found"
// @Failure 500 {object} dto.InternalServerErrorResponse "Internal server error"
// @Router /subscription-plans/{id} [get]
func (c *SubscriptionPlanController) Get(ctx *gin.Context) {
	id, _ := uuid.Parse(ctx.Param("id"))
	plan, err := c.service.GetSubscriptionPlan(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	var resp dto.SubscriptionPlanResponse
	err = copier.Copy(&resp, plan)
	ctx.JSON(http.StatusOK, resp)
}


// @BasePath /api/v1

// Create godoc
// @Security CookieAuth
// @Summary Retrieve All Subscription Plans
// @Description Subscription plan management
// @Tags subscription plans
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param pageSize query int false "Number of items per page" default(10)
// @Success 200 {array} dto.SubscriptionPlanResponse
// @Failure 400 {object} dto.BadRequestResponse  "Bad request, validation error"
// @Failure 401 {object} dto.UnauthorizedResponse  "Unauthorized, authentication required"
// @Failure 403 {object} dto.ForbiddenResponse  "Forbidden"
// @Failure 404 {object} dto.NotFoundResponse "Resource not found"
// @Failure 500 {object} dto.InternalServerErrorResponse "Internal server error"
// @Router /subscription-plans/ [get]
func (c *SubscriptionPlanController) List(ctx *gin.Context) {
	page := getQueryInt(ctx, "page", 1)
	pageSize := getQueryInt(ctx, "pageSize", 10)
	plans, err := c.service.ListSubscriptionPlans(page, pageSize)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, plans)
}

// @BasePath /api/v1

// UpdateSubscriptionPlan godoc
// @Summary Edit subscription plan
// @Description Update an existing subscription plan by ID
// @Tags subscription plan
// @Security CookieAuth
// @Accept json
// @Produce json
// @Param id path string true "Subscription Plan ID"
// @Param subscription_plan body repository.SubscriptionPlanPatch true "Partial subscription plan update"
// @Success 200 {object} dto.SubscriptionPlanResponse
// @Failure 400 {object} dto.BadRequestResponse "Bad request, validation error"
// @Failure 401 {object} dto.UnauthorizedResponse "Unauthorized, authentication required"
// @Failure 403 {object} dto.ForbiddenResponse "Forbidden"
// @Failure 404 {object} dto.NotFoundResponse "Subscription plan not found"
// @Failure 500 {object} dto.InternalServerErrorResponse "Internal server error"
// @Router /subscription-plans/{id} [patch]
func (c *SubscriptionPlanController) Update(ctx *gin.Context) {
	id, _ := uuid.Parse(ctx.Param("id"))
	var patch repository.SubscriptionPlanPatch
	if err := ctx.ShouldBindJSON(&patch); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	plan, err := c.service.UpdateSubscriptionPlan(id, &patch)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var resp dto.SubscriptionPlanResponse
	err = copier.Copy(&resp, plan)
	
	ctx.JSON(http.StatusOK, resp)
}

// DeleteSubscriptionPlan godoc
// @Summary Delete subscription plan
// @Description Permanently remove a subscription plan by ID
// @Tags subscription plan
// @Security CookieAuth
// @Accept json
// @Produce json
// @Param id path string true "Subscription Plan ID"
// @Success 204 "No Content"
// @Failure 400 {object} dto.BadRequestResponse "Bad request, validation error"
// @Failure 401 {object} dto.UnauthorizedResponse "Unauthorized, authentication required"
// @Failure 403 {object} dto.ForbiddenResponse "Forbidden"
// @Failure 404 {object} dto.NotFoundResponse "Subscription plan not found"
// @Failure 500 {object} dto.InternalServerErrorResponse "Internal server error"
// @Router /subscription-plans/{id} [delete]
func (c *SubscriptionPlanController) Delete(ctx *gin.Context) {
	id, _ := uuid.Parse(ctx.Param("id"))
	if err := c.service.DeleteSubscriptionPlan(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.Status(http.StatusNoContent)
}

