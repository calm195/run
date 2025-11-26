package router

import "github.com/gin-gonic/gin"

type StandardRouter struct{}

func (s *StandardRouter) InitRouter(router *gin.RouterGroup) {
	standardRouterGroup := router.Group("/standard")

	standardRouterGroup.GET("/list", standardApi.ListAllStandard)
}
