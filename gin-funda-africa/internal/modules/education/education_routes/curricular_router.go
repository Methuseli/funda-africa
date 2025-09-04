package education_routes

import (
	"gin-funda-africa/internal/modules/education/controllers"
	"gin-funda-africa/internal/modules/education/services"

	"github.com/gin-gonic/gin"
)



func RegisterCurriculaRoutes(rg *gin.RouterGroup, service *services.CurriculaService) {
	controller := controllers.NewCurriculaController(service)

	orgs := rg.Group("/curricula")
	{
		orgs.GET("/:id", controller.Get)
		orgs.GET("/", controller.GetCurricula)
	}
}
