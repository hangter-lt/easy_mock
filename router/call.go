package router

import (
	"github.com/hangter-lt/easy_mock/api"

	"github.com/gin-gonic/gin"
)

func RouterCall() *gin.Engine {
	r := gin.Default()

	r.Any("*path", api.Call.Handler)

	return r
}
