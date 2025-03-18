package logic

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hangter-lt/easy_mock/consts"
	"github.com/hangter-lt/easy_mock/global"
	"github.com/hangter-lt/easy_mock/initialize"
	"github.com/hangter-lt/easy_mock/model"
	"github.com/hangter-lt/easy_mock/model/dto"
	"github.com/hangter-lt/easy_mock/utils"

	"github.com/gin-gonic/gin"
)

type manageLogic struct{}

var Manage = manageLogic{}

func (*manageLogic) Create(c *gin.Context, req *dto.ManageCreateReq) error {

	// 验证group和name是否已经存在
	checkExist, err := global.DB.Query(`SELECT * FROM api_data WHERE "name" = ? AND "group" = ?;`,
		req.Name, req.Group)
	if err != nil {
		return err
	}
	defer checkExist.Close()
	if checkExist.Next() {
		return fmt.Errorf("already exists")
	}

	apiId := utils.UUID()
	paramsIds := []string{}
	for _, param := range req.Params {
		id := utils.UUID()
		global.ApiParam[id] = model.ApiParam{
			ApiId:          apiId,
			Route:          param.Route,
			ReqData:        param.ReqData,
			ResData:        param.ResData,
			ResContentType: param.ResContentType,
		}
		paramsIds = append(paramsIds, id)

		route := []any{}
		for _, pR := range param.Route {
			route = append(route, pR)
		}
		path := req.Path
		if len(param.Route) != 0 {
			path = fmt.Sprintf(req.Path, route...)
		}

		// 根据参数构建正则表达式
		requestRe := fmt.Sprintf(`^(%s):%s:%s`,
			strings.Join(req.Methods, "|"),
			consts.ContentTypeRe[req.ReqContentType],
			path,
		)

		if v, ok := global.ReqParam[requestRe]; !ok {
			global.ReqParam[requestRe] = []string{id}
		} else {
			global.ReqParam[requestRe] = append(v, id)
		}
	}

	global.ApiData[apiId] = model.ApiData{
		Name:           req.Name,
		Group:          req.Group,
		Path:           req.Path,
		Methods:        strings.Join(req.Methods, "|"),
		ReqContentType: req.ReqContentType,
	}

	// 数据写入数据库
	tx, err := global.DB.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec(consts.PreSqlCreateApiData,
		apiId, req.Name, req.Group, req.Path,
		strings.Join(req.Methods, "|"), req.ReqContentType)
	if err != nil {
		return err
	}

	for _, paramId := range paramsIds {
		reqData, err := json.Marshal(global.ApiParam[paramId].ReqData)
		if err != nil {
			return err
		}
		_, err = tx.Exec(consts.PreSqlCreateApiParam,
			paramId, apiId,
			strings.Join(global.ApiParam[paramId].Route, ","),
			string(reqData),
			global.ApiParam[paramId].ResData,
			global.ApiParam[paramId].ResContentType,
		)
		if err != nil {
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (*manageLogic) Update(c *gin.Context, req *dto.ManageUpdateReq) error {

	// 检查数据是否存在
	checkExist, err := global.DB.Query("SELECT * FROM api_data WHERE id = ?;", req.Id)
	if err != nil {
		return err
	}
	defer checkExist.Close()
	if !checkExist.Next() {
		return fmt.Errorf("not exists")
	}

	tx, err := global.DB.Begin()
	if err != nil {
		return err
	}

	// 更新数据api_data
	_, err = tx.Exec(`
		UPDATE "api_data" 
		SET "name" = ?, "group" = ?, "path" = ?, "methods" = ?, "req_content_type" = ? 
		WHERE id = ?;
		`, req.Name, req.Group, req.Path, strings.Join(req.Methods, "|"), req.ReqContentType, req.Id)
	if err != nil {
		return err
	}

	// 删除原有api_params中旧数据
	_, err = tx.Exec(`
		DELETE FROM "api_param" WHERE "api_id" = ?;
		`, req.Id)
	if err != nil {
		return err
	}

	// 创建数据api_params
	for _, param := range req.Params {
		id := utils.UUID()
		reqData, err := json.Marshal(param.ReqData)
		if err != nil {
			return err
		}
		_, err = tx.Exec(consts.PreSqlCreateApiParam,
			id, req.Id,
			strings.Join(param.Route, ","),
			string(reqData),
			param.ResData,
			param.ResContentType,
		)
		if err != nil {
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return initialize.LoadData()

}

func (*manageLogic) List(c *gin.Context) (dto.ManageListRes, error) {

	// 查询所有api_data
	apiDatas, err := global.DB.Query(`
	SELECT "id", "name", "group", "path", "methods", "req_content_type"
	FROM api_data;`)
	if err != nil {
		return nil, err
	}
	defer apiDatas.Close()

	list := make(dto.ManageListRes)
	for apiDatas.Next() {
		var apiData model.ApiData
		err = apiDatas.Scan(&apiData.Id, &apiData.Name,
			&apiData.Group, &apiData.Path, &apiData.Methods,
			&apiData.ReqContentType)
		if err != nil {
			return nil, err
		}

		if v, ok := list[apiData.Group]; !ok {
			list[apiData.Group] = []dto.ManageList{
				{
					Group:          apiData.Group,
					Id:             apiData.Id,
					Methods:        strings.Split(apiData.Methods, "|"),
					Name:           apiData.Name,
					Path:           apiData.Path,
					ReqContentType: apiData.ReqContentType,
				},
			}
		} else {
			list[apiData.Group] = append(v, dto.ManageList{
				Group:          apiData.Group,
				Id:             apiData.Id,
				Methods:        strings.Split(apiData.Methods, "|"),
				Name:           apiData.Name,
				Path:           apiData.Path,
				ReqContentType: apiData.ReqContentType,
			})
		}
	}

	return list, nil
}

func (*manageLogic) Info(c *gin.Context, req *dto.ManageInfoReq) (*dto.ManageInfoRes, error) {

	// 查询api_params
	apiParams, err := global.DB.Query(`
	SELECT "id", "api_id", "route", "req_data", "res_data", "res_content_type"
	FROM api_param WHERE api_id = ?;`, req.Id)
	if err != nil {
		return nil, err
	}
	defer apiParams.Close()

	params := []dto.ManageInfoParam{}
	for apiParams.Next() {
		var apiParam model.ApiParam
		var routeS string
		var reqDataS string
		err := apiParams.Scan(&apiParam.Id, &apiParam.ApiId, &routeS,
			&reqDataS, &apiParam.ResData, &apiParam.ResContentType)
		if err != nil {
			return nil, err
		}
		reqData := map[string]any{}
		err = json.Unmarshal([]byte(reqDataS), &reqData)
		if err != nil {
			return nil, err
		}

		params = append(params, dto.ManageInfoParam{
			Id:             apiParam.Id,
			ReqData:        reqData,
			ResContentType: apiParam.ResContentType,
			ResData:        apiParam.ResData,
			Route:          strings.Split(routeS, ","),
		})
	}

	// 查询api_data
	apiDatas, err := global.DB.Query(`
	SELECT "id", "name", "group", "path", "methods", "req_content_type"
	FROM api_data WHERE id = ?;`, req.Id)
	if err != nil {
		return nil, err
	}
	defer apiDatas.Close()

	var apiData model.ApiData
	apiDatas.Next()
	err = apiDatas.Scan(&apiData.Id, &apiData.Name,
		&apiData.Group, &apiData.Path, &apiData.Methods,
		&apiData.ReqContentType)
	if err != nil {
		return nil, err
	}

	res := dto.ManageInfoRes{
		Id:             apiData.Id,
		Name:           apiData.Name,
		Group:          apiData.Group,
		Path:           apiData.Path,
		Methods:        strings.Split(apiData.Methods, "|"),
		ReqContentType: apiData.ReqContentType,
		Params:         params,
	}

	return &res, nil
}

func (*manageLogic) Delete(c *gin.Context, req *dto.ManageDeleteReq) error {
	tx, err := global.DB.Begin()
	if err != nil {
		return err
	}

	// 删除api_data
	_, err = global.DB.Exec(`
	DELETE FROM "api_data" WHERE id = ?;`, req.Id)
	if err != nil {
		return err
	}

	// 删除api_params
	_, err = global.DB.Exec(`
	DELETE FROM"api_param" WHERE "api_id" = ?;`, req.Id)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return initialize.LoadData()
}
