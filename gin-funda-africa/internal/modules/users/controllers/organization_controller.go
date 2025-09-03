package controllers

import (
	"gin-funda-africa/internal/modules/users/models"
	"gin-funda-africa/internal/modules/users/repository"
	"gin-funda-africa/internal/modules/users/services"
	"gin-funda-africa/internal/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type OrganizationController struct {
	service *services.OrganizationService
}

func NewOrganizationController(s *services.OrganizationService) *OrganizationController {
	return &OrganizationController{service: s}
}

func (c *OrganizationController) Create(ctx *gin.Context) {
	var org models.Organization
	if err := ctx.ShouldBindJSON(&org); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := c.service.CreateOrganization(&org)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, result)
}

func (c *OrganizationController) Get(ctx *gin.Context) {
	id, _ := uuid.Parse(ctx.Param("id"))
	org, err := c.service.GetOrganization(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, org)
}

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
	ctx.JSON(http.StatusOK, org)
}

func (c *OrganizationController) Delete(ctx *gin.Context) {
	id, _ := uuid.Parse(ctx.Param("id"))
	if err := c.service.DeleteOrganization(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.Status(http.StatusNoContent)
}


