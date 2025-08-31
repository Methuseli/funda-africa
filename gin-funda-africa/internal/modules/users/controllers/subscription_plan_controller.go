package controllers

import (
	"gin-funda-africa/internal/modules/users/models"
	"gin-funda-africa/internal/modules/users/repository"
	"gin-funda-africa/internal/modules/users/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type SubscriptionPlanController struct {
	service *services.SubscriptionPlanService
}

func NewSubscriptionPlanController(s *services.SubscriptionPlanService) *SubscriptionPlanController {
	return &SubscriptionPlanController{service: s}
}

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
	ctx.JSON(http.StatusCreated, result)
}

func (c *SubscriptionPlanController) Get(ctx *gin.Context) {
	id, _ := uuid.Parse(ctx.Param("id"))
	plan, err := c.service.GetSubscriptionPlan(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, plan)
}

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
	ctx.JSON(http.StatusOK, plan)
}

func (c *SubscriptionPlanController) Delete(ctx *gin.Context) {
	id, _ := uuid.Parse(ctx.Param("id"))
	if err := c.service.DeleteSubscriptionPlan(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.Status(http.StatusNoContent)
}
