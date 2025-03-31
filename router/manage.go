package router

import (
	"github.com/hangter-lt/easy_mock/api"

	"github.com/gin-gonic/gin"
)

func RouterManage() *gin.Engine {
	r := gin.Default()

	manage := r.Group("/manages")
	{
		manage.POST("/create", api.Manage.Create)
		manage.POST("/update", api.Manage.Update)
		manage.DELETE("/:id", api.Manage.Delete)
		manage.GET("", api.Manage.List)
		manage.GET("/:id", api.Manage.Info)
		manage.GET("/groups", api.Manage.Groups)
	}
	request := r.Group("/request")
	{
		request.GET("/realtime", api.Request.RealTime)
		request.GET("/:id")
		request.GET("historys")
	}

	return r
}
