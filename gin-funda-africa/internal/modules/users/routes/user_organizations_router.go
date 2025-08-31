package routes

import (
	"gin-funda-africa/internal/modules/users/controllers"
	"gin-funda-africa/internal/modules/users/services"
	"github.com/gin-gonic/gin"
)

const orgParam = "/:userID/:orgID"

func RegisterUserOrganizationRoutes(rg *gin.RouterGroup, service *services.UserOrganizationService) {
	controller := controllers.NewUserOrganizationController(service)

	uos := rg.Group("/user-organizations")
	{
		uos.POST("/", controller.Create)
		uos.GET(orgParam, controller.Get)
		uos.GET("/", controller.List)
		uos.PATCH(orgParam, controller.Update)
		uos.DELETE(orgParam, controller.Delete)
	}
}
