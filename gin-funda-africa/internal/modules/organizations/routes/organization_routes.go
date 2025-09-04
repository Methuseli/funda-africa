package organization_routes

import (
	"gin-funda-africa/internal/modules/organizations/services"
	"gin-funda-africa/internal/pkg/middleware"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

func RegisterAllOrganizationServiceRoutes(router *gin.Engine, enforcer *casbin.Enforcer, jwtKey []byte) {
	protected := router.Group("/api/v1")
	protected.Use(middleware.AuthMiddleware(jwtKey))
	protected.Use(middleware.RequirePermission(enforcer))
	{
		RegisterOrganizationRoutes(protected, services.NewOrganizationService())
		RegisterOrganizationSubscriptionRoutes(protected, services.NewOrganizationSubscriptionService())
		RegisterOrganizationUserRoutes(protected, services.NewOrganizationUserService())
	}
}