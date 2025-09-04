package routes

import (
	"gin-funda-africa/internal/modules/users/controllers"
	"gin-funda-africa/internal/modules/users/services"
	"github.com/gin-gonic/gin"
)

const planParam = "/:userID/:planID"

func RegisterUserSubscriptionRoutes(rg *gin.RouterGroup, service *services.UserSubscriptionService) {
	controller := controllers.NewUserSubscriptionController(service)

	subs := rg.Group("/user-subscriptions")
	{
		subs.POST("/", controller.Create)
		subs.GET(planParam, controller.Get)
		subs.GET("/", controller.List)
		subs.PATCH(planParam, controller.Update)
	}
}
