package education_routes

import (
	"gin-funda-africa/internal/modules/education/controllers"
	"gin-funda-africa/internal/modules/education/services"

	"github.com/gin-gonic/gin"
)



func RegisterTopicRoutes(rg *gin.RouterGroup, service *services.TopicService) {
	controller := controllers.NewTopicController(service)

	orgs := rg.Group("/topics")
	{
		orgs.GET("/:id", controller.GetTopic)
		orgs.GET("/", controller.GetTopics)
	}
}
