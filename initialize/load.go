package initialize

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hangter-lt/easy_mock/consts"
	"github.com/hangter-lt/easy_mock/global"
	"github.com/hangter-lt/easy_mock/model"
)

func LoadData() error {
	// 初始化全局变量
	InitGlobal()

	// 查询所有api_data表数据, 写入全局变量
	apiDatas, err := global.DB.Query(`
	SELECT "id", "name", "group", "path", "methods", "req_content_type", "description"
	FROM api_data;`)
	if err != nil {
		return err
	}
	defer apiDatas.Close()
	for apiDatas.Next() {
		var apiData model.ApiData
		err = apiDatas.Scan(&apiData.Id, &apiData.Name,
			&apiData.Group, &apiData.Path, &apiData.Methods,
			&apiData.ReqContentType, &apiData.Description)
		if err != nil {
			return err
		}
		global.ApiData[apiData.Id] = apiData
	}

	// 查询所有api_param表数据, 写入全局变量
	apiParams, err := global.DB.Query(`
	SELECT "id", "api_id", "route", "req_data", "res_data", "res_content_type"
	FROM api_param;`)
	if err != nil {
		return err
	}
	defer apiParams.Close()
	for apiParams.Next() {
		var apiParam model.ApiParam
		var routeS string
		var reqDataS string
		err := apiParams.Scan(&apiParam.Id, &apiParam.ApiId, &routeS,
			&reqDataS, &apiParam.ResData, &apiParam.ResContentType)
		if err != nil {
			return err
		}
		route := strings.Split(routeS, ",")
		apiParam.Route = route
		mapReaData := map[string]any{}
		err = json.Unmarshal([]byte(reqDataS), &mapReaData)
		if err != nil {
			return err
		}
		apiParam.ReqData = mapReaData

		global.ApiParam[apiParam.Id] = apiParam

		routeAny := []any{}
		for _, r := range route {
			routeAny = append(routeAny, r)
		}
		path := global.ApiData[apiParam.ApiId].Path
		if len(route) > 1 {
			path = fmt.Sprintf(path, routeAny...)
		}

		// 根据参数构建正则表达式, 写入全局变量
		requestRe := fmt.Sprintf(`^(%s):%s:%s`,
			global.ApiData[apiParam.ApiId].Methods,
			consts.ContentTypeRe[global.ApiData[apiParam.ApiId].ReqContentType],
			path,
		)

		if v, ok := global.ReqParam[requestRe]; !ok {
			global.ReqParam[requestRe] = []string{apiParam.Id}
		} else {
			global.ReqParam[requestRe] = append(v, apiParam.Id)
		}
	}

	return nil
}
