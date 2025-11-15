package router

import "github.com/gin-gonic/gin"

type RecordRouter struct{}

func (r *RecordRouter) InitRecordRouter(Router *gin.RouterGroup) {
	recordRouter := Router.Group("/record")

	recordRouter.POST("/create", recordApi.CreateRecord)
	recordRouter.PUT("/edit", recordApi.UpdateRecord)
	recordRouter.GET("/list", recordApi.ListRecords)
	recordRouter.DELETE("/delete", recordApi.DeleteRecord)
}
