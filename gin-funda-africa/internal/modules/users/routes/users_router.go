package routes

import (
	"gin-funda-africa/internal/modules/users/controllers"
	"gin-funda-africa/internal/modules/users/services"
	"github.com/gin-gonic/gin"
)

const usersRoute = "/users"

func RegisterUserRoutes(routerGroup *gin.RouterGroup, service *services.UserService) {
	controller := controllers.NewUserController(service)

	users := routerGroup.Group(usersRoute)
	{
		users.GET("/:id", controller.Get)
		users.PATCH("/:id", controller.Update)
		users.DELETE("/:id", controller.Delete)
	}
}

func RegisterAuthRoutes(routerGroup *gin.RouterGroup, service *services.UserService) {
	controller := controllers.NewAuthController(service)
	users := routerGroup.Group(usersRoute)
	{
		users.POST("/", controller.Create)
		users.POST("/login", controller.Login)
	}
}

func RegisterAdminRoutes(routerGroup *gin.RouterGroup, service *services.UserService) {
	controller := controllers.NewAdminUserController(service)

	users := routerGroup.Group(usersRoute)
	{
		users.GET("/", controller.List)
	}

}
