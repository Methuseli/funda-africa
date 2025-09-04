package organization_routes

import (
	"gin-funda-africa/internal/modules/organizations/controllers"
	"gin-funda-africa/internal/modules/organizations/services"

	"github.com/gin-gonic/gin"
)

const planParam = "/:orgID/:planID"

func RegisterOrganizationSubscriptionRoutes(rg *gin.RouterGroup, service *services.OrganizationSubscriptionService) {
	controller := controllers.NewOrganizationSubscriptionController(service)

	subs := rg.Group("/organization-subscriptions")
	{
		subs.POST("/", controller.Create)
		subs.GET(planParam, controller.Get)
		subs.GET("/", controller.List)
		subs.PATCH(planParam, controller.Update)
	}
}
