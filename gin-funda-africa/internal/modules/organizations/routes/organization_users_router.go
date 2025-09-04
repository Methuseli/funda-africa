package organization_routes

import (
	"gin-funda-africa/internal/modules/organizations/controllers"
	"gin-funda-africa/internal/modules/organizations/services"

	"github.com/gin-gonic/gin"
)

const orgParam = "/:userID/:orgID"

func RegisterOrganizationUserRoutes(rg *gin.RouterGroup, service *services.OrganizationUserService) {
	controller := controllers.NewOrganizationUserController(service)

	uos := rg.Group("/organization-users")
	{
		uos.POST("/", controller.Create)
		uos.GET(orgParam, controller.Get)
		uos.GET("/", controller.List)
		uos.PATCH(orgParam, controller.Update)
	}
}
