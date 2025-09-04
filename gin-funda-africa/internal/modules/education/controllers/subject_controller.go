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



type SubjectController struct {
	service *services.SubjectService
}

func NewSubjectController(service *services.SubjectService) *SubjectController {
	return &SubjectController{service: service}
}

// @BasePath /api/v1/

// @Security CookieAuth
// @Summary Retrieve Subject
// @Description Get a single subject object by their unique UUID
// @Tags subject
// @Accept json
// @Produce json
// @Param id path string true "Subject UUID"
// @Success 200 {object} dto.SubjectResponse
// @Failure 400 {object} dto.BadRequestResponse "Bad request, invalid UUID"
// @Failure 401 {object} dto.UnauthorizedResponse "Unauthorized"
// @Failure 403 {object} dto.ForbiddenResponse "Forbidden"
// @Failure 404 {object} dto.NotFoundResponse "Subject not found"
// @Failure 500 {object} dto.InternalServerErrorResponse "Internal server error"
// @Router /education/subjects/{id} [get]
func (c *SubjectController) GetSubject(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": utils.InvalidUUID})
		return
	}

	subject, err := c.service.GetSubjectByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	var resp dto.SubjectResponse
	err = copier.Copy(&resp, subject)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

// @BasePath /api/v1

// Create godoc
// @Security CookieAuth
// @Summary Retrieve all subjects by country
// @Description Curricula management
// @Tags subjects
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param pageSize query int false "Number of items per page" default(10)
// @Param curriculum_id query string false "Country where the subject is from"
// @Param level query string false "Examining board for the subject"
// @Param name query string false "Query using subject name"
// @Param code query string false "Syllabus version"
// @Success 200 {array} dto.SubjectResponse
// @Failure 400 {object} dto.BadRequestResponse  "Bad request, validation error"
// @Failure 401 {object} dto.UnauthorizedResponse  "Unauthorized, authentication required"
// @Failure 403 {object} dto.ForbiddenResponse  "Forbidden"
// @Failure 404 {object} dto.NotFoundResponse "Resource not found"
// @Failure 500 {object} dto.InternalServerErrorResponse "Internal server error"
// @Router /education/subjects/ [get]
func (c *SubjectController) GetSubjects(ctx *gin.Context) {
	var params utils.Params
	err := ctx.ShouldBindQuery(&params)
	page := utils.GetQueryInt(ctx, "page", 1)
	pageSize := utils.GetQueryInt(ctx, "pageSize", 10)

	subjects, err := c.service.GetSubjects(params, page, pageSize)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, subjects)
}
