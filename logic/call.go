package logic

import (
	"encoding/json"
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

	isMatch := false
	defer c.Set("is_match", isMatch)

	reqParam := fmt.Sprintf("%s:%s:%s", c.Request.Method, c.ContentType(), c.Request.URL.Path)
	fmt.Printf("reqParam: %v\n", reqParam)

	paramIds := []string{}
	for k := range global.ReqParam {
		reg := regexp.MustCompile(k)
		if reg.Match([]byte(reqParam)) {
			paramIds = append(paramIds, global.ReqParam[k]...)
		}
	}

	// 未命中api
	if len(paramIds) == 0 {
		return
	}

	// TODO: 确认请求类型

	req := make(map[string]any)
	reqData, _ := c.Get("req_data")
	fmt.Printf("reqData: %v\n", reqData)
	err := json.Unmarshal(reqData.([]byte), &req)
	if err != nil {
		return
	}

	fmt.Printf("req: %v\n", req)

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

	isMatch = true

	rand.NewSource(time.Now().UnixNano())
	index := rand.Intn(len(targetParamIds))
	resContentType := global.ApiParam[targetParamIds[index]].ResContentType
	resData := global.ApiParam[targetParamIds[index]].ResData

	switch resContentType {
	case "application/json":
		res := map[string]any{}
		err := json.Unmarshal([]byte(resData), &res)
		if err != nil {
			c.String(200, string(resData))
			return
		}
		c.JSON(200, res)
	case "text/html":
		c.HTML(200, "index.html", resData)
	case "text/plain":
		c.String(200, resData)
	case "application/xml":
		c.XML(200, resData)
	}

}
