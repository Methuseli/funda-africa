package routes

import (
	"gin-funda-africa/internal/modules/users/services"
	"gin-funda-africa/internal/pkg/middleware"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RegisterAllUserServiceRoutes(router *gin.Engine, enforcer *casbin.Enforcer, jwtKey []byte) {
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	api := router.Group("/api/v1")

	RegisterAuthRoutes(api, services.NewUserService(jwtKey))
	RegisterRetrieveSubscriptionPlanRoutes(api, services.NewSubscriptionPlanService())

	protected := api.Group("/")
	protected.Use(middleware.AuthMiddleware(jwtKey))
	protected.Use(middleware.RequirePermission(enforcer))
	{
		RegisterAdminRoutes(protected, services.NewUserService(jwtKey))
		RegisterUserRoutes(protected, services.NewUserService(jwtKey))
		RegisterSubscriptionPlanRoutes(protected, services.NewSubscriptionPlanService())
		RegisterUserSubscriptionRoutes(protected, services.NewUserSubscriptionService())
	}

}
