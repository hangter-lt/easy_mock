package api

import (
	"github.com/gin-gonic/gin"
	"github.com/hangter-lt/easy_mock/logic"
)

type requestApi struct{}

var Request = requestApi{}

func (*requestApi) RealTime(c *gin.Context) {
	logic.Request.RealTime(c)
}
