package logic

import (
	"fmt"
	"math/rand"
	"regexp"
	"time"

	"github.com/hangter-lt/easy_mock/global"
	"github.com/hangter-lt/easy_mock/utils"

	"github.com/gin-gonic/gin"
)

type callLogic struct{}

var Call = callLogic{}

func (*callLogic) Handler(c *gin.Context) {
	reqParam := fmt.Sprintf("%s:%s:%s", c.Request.Method, c.ContentType(), c.Request.URL.Path)

	paramIds := []string{}
	for k := range global.ReqParam {
		reg := regexp.MustCompile(k)
		if reg.Match([]byte(reqParam)) {
			paramIds = append(paramIds, global.ReqParam[k]...)
		}
	}

	if len(paramIds) == 0 {
		return
	}

	// 确认请求类型
	reqContentType := global.ApiData[global.ApiParam[paramIds[0]].ApiId].ReqContentType
	// reqContentType := "application/json"

	req := make(map[string]any)
	switch reqContentType {
	case "application/json":
		err := c.ShouldBindJSON(&req)
		if err != nil {
			fmt.Printf("err: %v\n", err)
			return
		}
	case "application/x-www-form-urlencoded":
		err := c.Request.ParseForm()
		if err != nil {
			fmt.Printf("err: %v\n", err)
			return
		}
		for k, v := range c.Request.Form {
			if len(v) == 1 {
				req[k] = v[0]
			} else {
				req[k] = v
			}
		}
	case "":
		err := c.Request.ParseForm()
		if err != nil {
			fmt.Printf("err: %v\n", err)
			return
		}
		for k, v := range c.Request.Form {
			if len(v) == 1 {
				req[k] = v[0]
			} else {
				req[k] = v
			}
		}
	case "multipart/form-data":
		err := c.Request.ParseMultipartForm(10 << 20)
		if err != nil {
			fmt.Printf("err: %v\n", err)
			return
		}
		for k, v := range c.Request.MultipartForm.Value {
			if len(v) == 1 {
				req[k] = v[0]
			} else {
				req[k] = v
			}
		}
	}

	fmt.Printf("req: %+v\n", req)

	// 参数匹配
	targetParamIds := []string{}
	for _, paramId := range paramIds {
		if utils.IsAInB(req, global.ApiParam[paramId].ReqData) {
			targetParamIds = append(targetParamIds, paramId)
		}
	}

	// 未命中参数
	if len(targetParamIds) == 0 {
		return
	}

	rand.NewSource(time.Now().UnixNano())
	index := rand.Intn(len(targetParamIds))
	resContentType := global.ApiParam[targetParamIds[index]].ResContentType
	resData := global.ApiParam[targetParamIds[index]].ResData

	switch resContentType {
	case "application/json":
		c.JSON(200, resData)
	case "text/html":
		c.HTML(200, "index.html", resData)
	case "text/plain":
		c.String(200, resData)
	case "application/xml":
		c.XML(200, resData)
	}

}
