package controllers

import (
	"gin-funda-africa/internal/modules/users/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AdminUserController struct {
	service *services.UserService
}

func NewAdminUserController(service *services.UserService) *AdminUserController {
	return &AdminUserController{service: service}
}

// List users with pagination
func (c *AdminUserController) List(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "10"))

	users, err := c.service.ListUsers(page, pageSize)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, users)
}
