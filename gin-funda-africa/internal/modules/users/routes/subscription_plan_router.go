package routes

import (
	"gin-funda-africa/internal/modules/users/controllers"
	"gin-funda-africa/internal/modules/users/services"
	"github.com/gin-gonic/gin"
)

func RegisterSubscriptionPlanRoutes(rg *gin.RouterGroup, service *services.SubscriptionPlanService) {
	controller := controllers.NewSubscriptionPlanController(service)

	plans := rg.Group("/subscription-plans")
	{
		plans.POST("/", controller.Create)
		plans.PATCH("/:id", controller.Update)
		plans.DELETE("/:id", controller.Delete)
	}
}

func RegisterRetrieveSubscriptionPlanRoutes(rg *gin.RouterGroup, service *services.SubscriptionPlanService) {
	controller := controllers.NewSubscriptionPlanController(service)
	plans := rg.Group("/subscription-plans")
	{
		plans.GET("/:id", controller.Get)
		plans.GET("/", controller.List)
	}
}
