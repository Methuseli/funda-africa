package education_routes

import (
	"gin-funda-africa/internal/modules/education/controllers"
	"gin-funda-africa/internal/modules/education/services"

	"github.com/gin-gonic/gin"
)

func RegisterSubTopicsRoutes(rg *gin.RouterGroup, service *services.SubTopicService) {
	controller := controllers.NewSubTopicController(service)

	orgs := rg.Group("/sub-topics")
	{
		orgs.GET("/:id", controller.GetSubTopic)
		orgs.GET("/", controller.GetSubTopics)
	}
}
