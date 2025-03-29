package api

import (
	"log"

	"github.com/hangter-lt/easy_mock/logic"
	"github.com/hangter-lt/easy_mock/model/dto"

	"github.com/gin-gonic/gin"
)

type manageApi struct{}

var Manage = manageApi{}

func (*manageApi) Create(c *gin.Context) {
	var req dto.ManageCreateReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Println(err)
		c.Status(400)
		return
	}

	res, err := logic.Manage.Create(c, &req)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"data":    res,
	})
}

func (*manageApi) Update(c *gin.Context) {
	var req dto.ManageUpdateReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.Status(400)
		return
	}

	err = logic.Manage.Update(c, &req)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
	})
}

func (*manageApi) List(c *gin.Context) {
	res, err := logic.Manage.List(c)
	if err != nil {
		c.JSON(500, gin.H{
			"success": false,
			"message": err.Error(),
		})
	}

	c.JSON(200, gin.H{
		"success": true,
		"data":    res,
	})
}

func (*manageApi) Info(c *gin.Context) {
	var req dto.ManageInfoReq
	err := c.ShouldBindUri(&req)
	if err != nil {
		c.Status(400)
		return
	}

	res, err := logic.Manage.Info(c, &req)
	if err != nil {
		c.JSON(500, gin.H{
			"success": false,
			"message": err.Error(),
		})
	}

	c.JSON(200, gin.H{
		"success": true,
		"data":    res,
	})
}

func (*manageApi) Delete(c *gin.Context) {
	var req dto.ManageDeleteReq
	err := c.ShouldBindUri(&req)
	if err != nil {
		c.Status(400)
		return
	}

	err = logic.Manage.Delete(c, &req)
	if err != nil {
		c.JSON(500, gin.H{
			"success": false,
			"message": err.Error(),
		})
	}

	c.JSON(200, gin.H{
		"success": true,
	})
}

func (*manageApi) Groups(c *gin.Context) {

	data, err := logic.Manage.Groups(c)
	if err != nil {
		c.JSON(500, gin.H{
			"success": false,
			"message": err.Error(),
		})
	}

	c.JSON(200, gin.H{
		"success": true,
		"data":    data,
	})

}
