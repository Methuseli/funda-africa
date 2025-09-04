package organization_routes

import (
	"gin-funda-africa/internal/modules/organizations/controllers"
	"gin-funda-africa/internal/modules/organizations/services"

	"github.com/gin-gonic/gin"
)

func RegisterOrganizationRoutes(rg *gin.RouterGroup, service *services.OrganizationService) {
	controller := controllers.NewOrganizationController(service)

	orgs := rg.Group("/organizations")
	{
		orgs.POST("/", controller.Create)
		orgs.GET("/:id", controller.Get)
		orgs.GET("/", controller.List)
		orgs.GET("/user/:userID", controller.ListByUser)
		orgs.PATCH("/:id", controller.Update)
		orgs.DELETE("/:id", controller.Delete)
	}
}
