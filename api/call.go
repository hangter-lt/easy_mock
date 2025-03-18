package api

import (
	"github.com/hangter-lt/easy_mock/logic"

	"github.com/gin-gonic/gin"
)

type callApi struct{}

var Call = callApi{}

func (*callApi) Handler(c *gin.Context) {

	logic.Call.Handler(c)

}
