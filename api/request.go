package api

import (
	"github.com/gin-gonic/gin"
	"github.com/hangter-lt/easy_mock/logic"
	"github.com/hangter-lt/easy_mock/model/dto"
)

type requestApi struct{}

var Request = requestApi{}

func (*requestApi) RealTime(c *gin.Context) {
	logic.Request.RealTime(c)
}

func (*requestApi) Info(c *gin.Context) {
	var req dto.RequestInfoReq
	err := c.ShouldBindUri(&req)
	if err != nil {
		c.Status(400)
		return
	}

	data, err := logic.Request.Info(c, &req)
	if err != nil {
		c.JSON(500, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"data":    data,
	})

}
