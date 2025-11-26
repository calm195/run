package router

import "github.com/gin-gonic/gin"

type EventRouter struct{}

func (s *EventRouter) InitRouter(router *gin.RouterGroup) {
	eventRouterGroup := router.Group("/event")

	eventRouterGroup.GET("/list", eventApi.ListAllEvents)
}
