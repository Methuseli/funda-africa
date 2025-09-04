package controllers

import (
	"gin-funda-africa/internal/modules/education/services"
	"gin-funda-africa/internal/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type SubTopicController struct {
	service *services.SubTopicService
}

func NewSubTopicController(service *services.SubTopicService) *SubTopicController {
	return &SubTopicController{service: service}
}

// @BasePath /api/v1/

// @Security CookieAuth
// @Summary Retrieve Sub Topic
// @Description Get a single Sub Topic object by their unique UUID
// @Tags sub-topic
// @Accept json
// @Produce json
// @Param id path string true "Sub Topic UUID"
// @Success 200 {object} dto.SubTopicResponse
// @Failure 400 {object} dto.BadRequestResponse "Bad request, invalid UUID"
// @Failure 401 {object} dto.UnauthorizedResponse "Unauthorized"
// @Failure 403 {object} dto.ForbiddenResponse "Forbidden"
// @Failure 404 {object} dto.NotFoundResponse "User not found"
// @Failure 500 {object} dto.InternalServerErrorResponse "Internal server error"
// @Router /education/sub-topics/{id} [get]
func (c *SubTopicController) GetSubTopic(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": utils.InvalidUUID})
		return
	}

	subTopic, err := c.service.GetSubTopicByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, subTopic)
}

// @BasePath /api/v1

// Create godoc
// @Security CookieAuth
// @Summary Sub topics for an individual topic
// @Description Sub Topic management
// @Tags sub-topics
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param pageSize query int false "Number of items per page" default(10)
// @Param topic_id query string false "Topic related to the sub topic"
// @Success 200 {array} dto.CurriculumResponse
// @Failure 400 {object} dto.BadRequestResponse  "Bad request, validation error"
// @Failure 401 {object} dto.UnauthorizedResponse  "Unauthorized, authentication required"
// @Failure 403 {object} dto.ForbiddenResponse  "Forbidden"
// @Failure 404 {object} dto.NotFoundResponse "Resource not found"
// @Failure 500 {object} dto.InternalServerErrorResponse "Internal server error"
// @Router /education/sub-topics/ [get]
func (c *SubTopicController) GetSubTopics(ctx *gin.Context) {
	var params utils.Params
	err := ctx.ShouldBindQuery(&params)
	page := utils.GetQueryInt(ctx, "page", 1)
	pageSize := utils.GetQueryInt(ctx, "pageSize", 10)

	curricula, err := c.service.GetSubTopics(params, page, pageSize)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, curricula)
}
