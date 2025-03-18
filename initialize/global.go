package initialize

import (
	"github.com/hangter-lt/easy_mock/global"
	"github.com/hangter-lt/easy_mock/model"
)

func InitGlobal() {
	apiData := map[string]model.ApiData{}
	global.ApiData = apiData

	apiParam := map[string]model.ApiParam{}
	global.ApiParam = apiParam

	reqParam := map[string][]string{}
	global.ReqParam = reqParam
}
