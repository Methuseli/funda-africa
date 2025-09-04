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



type CurriculaController struct {
	service *services.CurriculaService
}

func NewCurriculaController(service *services.CurriculaService) *CurriculaController {
	return &CurriculaController{service: service}
}

// @BasePath /api/v1/

// @Security CookieAuth
// @Summary Retrieve Curriculum
// @Description Get a single curriculum object by their unique UUID
// @Tags curriculum
// @Accept json
// @Produce json
// @Param id path string true "Curriculum UUID"
// @Success 200 {object} dto.CurriculumResponse
// @Failure 400 {object} dto.BadRequestResponse "Bad request, invalid UUID"
// @Failure 401 {object} dto.UnauthorizedResponse "Unauthorized"
// @Failure 403 {object} dto.ForbiddenResponse "Forbidden"
// @Failure 404 {object} dto.NotFoundResponse "Curriculum not found"
// @Failure 500 {object} dto.InternalServerErrorResponse "Internal server error"
// @Router /education/curricula/{id} [get]
func (c *CurriculaController) Get(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": utils.InvalidUUID})
		return
	}

	curriculum, err := c.service.GetCurriculumById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	var resp dto.CurriculumResponse
	err = copier.Copy(&resp, curriculum)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

// @BasePath /api/v1

// Create godoc
// @Security CookieAuth
// @Summary Retrieve all curricula by country
// @Description Curricula management
// @Tags curricula
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param pageSize query int false "Number of items per page" default(10)
// @Param country query string false "Country where the curriculum is from"
// @Param board query string false "Examining board for the curriculum"
// @Param version query string false "Syllabus version"
// @Success 200 {array} dto.CurriculumResponse
// @Failure 400 {object} dto.BadRequestResponse  "Bad request, validation error"
// @Failure 401 {object} dto.UnauthorizedResponse  "Unauthorized, authentication required"
// @Failure 403 {object} dto.ForbiddenResponse  "Forbidden"
// @Failure 404 {object} dto.NotFoundResponse "Resource not found"
// @Failure 500 {object} dto.InternalServerErrorResponse "Internal server error"
// @Router /education/curricula/ [get]
func (c *CurriculaController) GetCurricula(ctx *gin.Context) {
	var params utils.Params
	err := ctx.ShouldBindQuery(&params)
	page := utils.GetQueryInt(ctx, "page", 1)
	pageSize := utils.GetQueryInt(ctx, "pageSize", 10)

	curricula, err := c.service.GetCurricula(params, page, pageSize)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, curricula)
}
