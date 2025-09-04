package controllers

import (
	"gin-funda-africa/internal/modules/users/dto"
	"gin-funda-africa/internal/modules/users/models"
	"gin-funda-africa/internal/modules/users/repository"
	"gin-funda-africa/internal/modules/users/services"
	"gin-funda-africa/internal/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
)

type UserSubscriptionController struct {
	service *services.UserSubscriptionService
}

func NewUserSubscriptionController(s *services.UserSubscriptionService) *UserSubscriptionController {
	return &UserSubscriptionController{service: s}
}

// @BasePath /api/v1

// Create godoc
// @Security CookieAuth
// @Summary Add User Subscription Plan
// @Description User Subscription plan management
// @Tags user-subscription
// @Accept json
// @Produce json
// @Param subscription_plan body dto.UserSubscriptionResponse true "User Subscription Plan"
// @Success 201 {object} dto.UserSubscriptionResponse
// @Failure 400 {object} dto.BadRequestResponse  "Bad request, validation error"
// @Failure 401 {object} dto.UnauthorizedResponse  "Unauthorized, authentication required"
// @Failure 403 {object} dto.ForbiddenResponse  "Forbidden"
// @Failure 500 {object} dto.InternalServerErrorResponse "Internal server error"
// @Router /user-subscription/ [post]
func (c *UserSubscriptionController) Create(ctx *gin.Context) {
	var sub models.UserSubscription
	if err := ctx.ShouldBindJSON(&sub); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userSubscription, err := c.service.CreateUserSubscription(&sub)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var resp dto.UserSubscriptionResponse
	err = copier.Copy(&resp, userSubscription)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, resp)
}

// @BasePath /api/v1/

// @Security CookieAuth
// @Summary Retrieve User Subscription
// @Description Get a single user subscription object by their unique UUID
// @Tags user-subscription
// @Accept json
// @Produce json
// @Param user_id path string true "User UUID"
// @Param plan_id path string true "Subscription plan UUID"
// @Success 200 {object} dto.UserSubscriptionResponse
// @Failure 400 {object} dto.BadRequestResponse "Bad request, invalid UUID"
// @Failure 401 {object} dto.UnauthorizedResponse "Unauthorized"
// @Failure 403 {object} dto.ForbiddenResponse "Forbidden"
// @Failure 404 {object} dto.NotFoundResponse "User Subscription not found"
// @Failure 500 {object} dto.InternalServerErrorResponse "Internal server error"
// @Router /user-subscription/{user_id}/{plan_id} [get]
func (c *UserSubscriptionController) Get(ctx *gin.Context) {
	userID, err := uuid.Parse(ctx.Param("user_id"))
	planID, err := uuid.Parse(ctx.Param("plan_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": utils.InvalidUUID})
		return
	}

	sub, err := c.service.GetUserSubscription(userID, planID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	var resp dto.UserSubscriptionResponse
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
// @Summary Retrieve all user subscriptions
// @Description User Subscription
// @Tags user-subscriptions
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param pageSize query int false "Number of items per page" default(20)
// @Success 200 {array} dto.UserSubscriptionResponse
// @Failure 400 {object} dto.BadRequestResponse  "Bad request, validation error"
// @Failure 401 {object} dto.UnauthorizedResponse  "Unauthorized, authentication required"
// @Failure 403 {object} dto.ForbiddenResponse  "Forbidden"
// @Failure 404 {object} dto.NotFoundResponse "Resource not found"
// @Failure 500 {object} dto.InternalServerErrorResponse "Internal server error"
// @Router /user-subscriptions/ [get]
func (c *UserSubscriptionController) List(ctx *gin.Context) {
	page := utils.GetQueryInt(ctx, "page", 1)
	pageSize := utils.GetQueryInt(ctx, "pageSize", 20)

	subs, err := c.service.ListUserSubscriptions(page, pageSize)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, subs)
}

// @BasePath /api/v1

// @Security CookieAuth
// @Summary Partially Update User Subscription
// @Description Update selected fields of a user subscription by their UUID
// @Tags user-subscription
// @Accept json
// @Produce json
// @Param user_id path string true "User UUID"
// @Param plan_id path string true "Subscription plan UUID"
// @Param user_subscription_patch body repository.UserSubscriptionPatch true "User fields to update"
// @Success 200 {object} dto.UserSubscriptionResponse
// @Failure 400 {object} dto.BadRequestResponse "Bad request, invalid UUID or payload"
// @Failure 401 {object} dto.UnauthorizedResponse "Unauthorized"
// @Failure 403 {object} dto.ForbiddenResponse "Forbidden"
// @Failure 500 {object} dto.InternalServerErrorResponse "Internal server error"
// @Router /user-subscriptions/{id} [patch]
func (c *UserSubscriptionController) Update(ctx *gin.Context) {
	userID, err := uuid.Parse(ctx.Param("user_id"))
	planID, err := uuid.Parse(ctx.Param("plan_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": utils.InvalidUUID})
		return
	}

	var patch repository.UserSubscriptionPatch
	if err := ctx.ShouldBindJSON(&patch); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	sub, err := c.service.UpdateUserSubscriptionPartial(userID, planID, &patch)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var resp dto.UserSubscriptionResponse
	err = copier.Copy(&resp, sub)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	ctx.JSON(http.StatusOK, sub)
}

