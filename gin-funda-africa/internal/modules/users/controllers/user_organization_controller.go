package controllers

import (
	"gin-funda-africa/internal/modules/users/models"
	"gin-funda-africa/internal/modules/users/repository"
	"gin-funda-africa/internal/modules/users/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserOrganizationController struct {
	service *services.UserOrganizationService
}

func NewUserOrganizationController(s *services.UserOrganizationService) *UserOrganizationController {
	return &UserOrganizationController{service: s}
}

func (c *UserOrganizationController) Create(ctx *gin.Context) {
	var uo models.UserOrganization
	if err := ctx.ShouldBindJSON(&uo); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := c.service.CreateUserOrganization(&uo)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, result)
}

func (c *UserOrganizationController) Get(ctx *gin.Context) {
	userID, _ := uuid.Parse(ctx.Param("userID"))
	orgID, _ := uuid.Parse(ctx.Param("orgID"))
	uo, err := c.service.GetUserOrganization(userID, orgID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, uo)
}

func (c *UserOrganizationController) List(ctx *gin.Context) {
	page := getQueryInt(ctx, "page", 1)
	pageSize := getQueryInt(ctx, "pageSize", 10)
	uos, err := c.service.ListUserOrganizations(page, pageSize)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, uos)
}

func (c *UserOrganizationController) Update(ctx *gin.Context) {
	userID, _ := uuid.Parse(ctx.Param("userID"))
	orgID, _ := uuid.Parse(ctx.Param("orgID"))
	var patch repository.UserOrganizationPatch
	if err := ctx.ShouldBindJSON(&patch); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := c.service.UpdateUserOrganizationPartial(userID, orgID, &patch)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, result)
}

func (c *UserOrganizationController) Delete(ctx *gin.Context) {
	userID, _ := uuid.Parse(ctx.Param("userID"))
	orgID, _ := uuid.Parse(ctx.Param("orgID"))
	if err := c.service.DeleteUserOrganization(userID, orgID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.Status(http.StatusNoContent)
}
