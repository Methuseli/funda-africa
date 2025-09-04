package education_routes

import (
	"gin-funda-africa/internal/modules/education/controllers"
	"gin-funda-africa/internal/modules/education/services"

	"github.com/gin-gonic/gin"
)

func RegisterSubjectRoutes(rg *gin.RouterGroup, service *services.SubjectService) {
	controller := controllers.NewSubjectController(service)

	orgs := rg.Group("/subjects")
	{
		orgs.GET("/:id", controller.GetSubject)
		orgs.GET("/", controller.GetSubjects)
	}
}
