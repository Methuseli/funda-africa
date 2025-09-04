package controllers

import (
	"gin-funda-africa/internal/modules/education/dto"
	"gin-funda-africa/internal/modules/education/services"
	"gin-funda-africa/internal/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
)



type TopicController struct {
	service *services.TopicService
}

func NewTopicController(service *services.TopicService) *TopicController {
	return &TopicController{service: service}
}

// @BasePath /api/v1/

// @Security CookieAuth
// @Summary Retrieve Topic
// @Description Get a single topic object by their unique UUID
// @Tags topic
// @Accept json
// @Produce json
// @Param id path string true "Topic UUID"
// @Success 200 {object} dto.TopicResponse
// @Failure 400 {object} dto.BadRequestResponse "Bad request, invalid UUID"
// @Failure 401 {object} dto.UnauthorizedResponse "Unauthorized"
// @Failure 403 {object} dto.ForbiddenResponse "Forbidden"
// @Failure 404 {object} dto.NotFoundResponse "Topic not found"
// @Failure 500 {object} dto.InternalServerErrorResponse "Internal server error"
// @Router /education/topics/{id} [get]
func (c *TopicController) GetTopic(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": utils.InvalidUUID})
		return
	}

	topic, err := c.service.GetTopicByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	var resp dto.TopicResponse
	err = copier.Copy(&resp, topic)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

// @BasePath /api/v1

// Create godoc
// @Security CookieAuth
// @Summary Retrieve all topics by country
// @Description Curricula management
// @Tags topics
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param pageSize query int false "Number of items per page" default(10)
// @Param subject_id query string false "Subject in which the topics belongs to"
// @Param name query string false "Topic name"
// @Success 200 {array} dto.TopicResponse
// @Failure 400 {object} dto.BadRequestResponse  "Bad request, validation error"
// @Failure 401 {object} dto.UnauthorizedResponse  "Unauthorized, authentication required"
// @Failure 403 {object} dto.ForbiddenResponse  "Forbidden"
// @Failure 404 {object} dto.NotFoundResponse "Resource not found"
// @Failure 500 {object} dto.InternalServerErrorResponse "Internal server error"
// @Router /education/topics/ [get]
func (c *TopicController) GetTopics(ctx *gin.Context) {
	var params utils.Params
	err := ctx.ShouldBindQuery(&params)
	page := utils.GetQueryInt(ctx, "page", 1)
	pageSize := utils.GetQueryInt(ctx, "pageSize", 10)

	topics, err := c.service.GetTopics(params, page, pageSize)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, topics)
}
