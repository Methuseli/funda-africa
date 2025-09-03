package controllers

import (
	"gin-funda-africa/internal/modules/users/models"
	"gin-funda-africa/internal/modules/users/repository"
	"gin-funda-africa/internal/modules/users/services"
	"gin-funda-africa/internal/pkg/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserSubscriptionController struct {
	service *services.UserSubscriptionService
}

func NewUserSubscriptionController(s *services.UserSubscriptionService) *UserSubscriptionController {
	return &UserSubscriptionController{service: s}
}

func (c *UserSubscriptionController) Create(ctx *gin.Context) {
	var sub models.UserSubscription
	if err := ctx.ShouldBindJSON(&sub); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := c.service.CreateUserSubscription(&sub)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, result)
}

func (c *UserSubscriptionController) Get(ctx *gin.Context) {
	userID, _ := uuid.Parse(ctx.Param("userID"))
	planID, _ := strconv.ParseUint(ctx.Param("planID"), 10, 64)
	sub, err := c.service.GetUserSubscription(userID, uint(planID))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, sub)
}

func (c *UserSubscriptionController) List(ctx *gin.Context) {
	page := utils.GetQueryInt(ctx, "page", 1)
	pageSize := utils.GetQueryInt(ctx, "pageSize", 10)
	subs, err := c.service.ListUserSubscriptions(page, pageSize)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, subs)
}

func (c *UserSubscriptionController) Update(ctx *gin.Context) {
	userID, _ := uuid.Parse(ctx.Param("userID"))
	planID, _ := strconv.ParseUint(ctx.Param("planID"), 10, 64)
	var patch repository.UserSubscriptionPatch
	if err := ctx.ShouldBindJSON(&patch); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	sub, err := c.service.UpdateUserSubscriptionPartial(userID, uint(planID), &patch)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, sub)
}

func (c *UserSubscriptionController) Delete(ctx *gin.Context) {
	userID, _ := uuid.Parse(ctx.Param("userID"))
	planID, _ := strconv.ParseUint(ctx.Param("planID"), 10, 64)
	if err := c.service.DeleteUserSubscription(userID, uint(planID)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.Status(http.StatusNoContent)
}
