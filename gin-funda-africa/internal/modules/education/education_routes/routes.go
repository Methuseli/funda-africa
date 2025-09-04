package education_routes

import (
	"gin-funda-africa/internal/modules/education/services"
	"gin-funda-africa/internal/pkg/middleware"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

func RegisterAllEducationServiceRoutes(router *gin.Engine, enforcer *casbin.Enforcer, jwtKey []byte) {
	protected := router.Group("/api/v1/education")
	protected.Use(middleware.AuthMiddleware(jwtKey))
	protected.Use(middleware.RequirePermission(enforcer))
	{
		RegisterCurriculaRoutes(protected, services.NewCurriculaService())
		RegisterSubTopicsRoutes(protected, services.NewSubTopicService())
		RegisterSubjectRoutes(protected, services.NewSubjectService())
		RegisterTopicRoutes(protected, services.NewTopicService())
	}
}