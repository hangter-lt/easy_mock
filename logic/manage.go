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

func (*manageLogic) Create(c *gin.Context, req *dto.ManageCreateReq) (*dto.ManageCreateRes, error) {

	// 验证group和name是否已经存在
	checkExist, err := global.DB.Query(`SELECT * FROM api_data WHERE "name" = ? AND "group" = ?;`,
		req.Name, req.Group)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return nil, err
	}
	if checkExist.Next() {
		return nil, fmt.Errorf("already exists")
	}
	checkExist.Close()

	// TODO: 检查path和route数量是否匹配

	// 数据写入数据库
	tx, err := global.DB.Begin()
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return nil, err
	}

	apiId := utils.UUID()
	_, err = tx.Exec(consts.PreSqlCreateApiData,
		apiId, req.Name, req.Group, req.Path,
		strings.Join(req.Methods, "|"), req.ReqContentType, req.Description)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return nil, err
	}

	// 创建数据api_params
	for _, param := range req.Params {
		id := utils.UUID()

		mapReqDatas := map[string]any{}
		for _, reqData := range param.ReqDatas {
			mapReqDatas[reqData.Key] = reqData.Value
		}
		reqDataS, err := json.Marshal(mapReqDatas)
		if err != nil {
			return nil, err
		}
		_, err = tx.Exec(consts.PreSqlCreateApiParam,
			id, apiId,
			strings.Join(param.Route, ","),
			string(reqDataS),
			param.ResData,
			param.ResContentType,
			param.ResCode,
		)
		if err != nil {
			fmt.Printf("err: %v\n", err)
			return nil, err
		}
	}

	err = tx.Commit()
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return nil, err
	}

	// TODO: 失败后删除数据
	return &dto.ManageCreateRes{Id: apiId}, initialize.LoadData()
}

func (*manageLogic) Update(c *gin.Context, req *dto.ManageUpdateReq) error {

	// 检查数据是否存在
	checkExist, err := global.DB.Query("SELECT * FROM api_data WHERE id = ?;", req.Id)
	if err != nil {
		return err
	}
	if !checkExist.Next() {
		return fmt.Errorf("not exists")
	}
	checkExist.Close()

	tx, err := global.DB.Begin()
	if err != nil {
		fmt.Printf("\"1\": %v\n", "1")
		return err
	}

	// 更新数据api_data
	_, err = tx.Exec(`
		UPDATE "api_data"
		SET "name" = ?, "group" = ?, "path" = ?, "methods" = ?, "req_content_type" = ?, "description" = ?
		WHERE id = ?;
		`, req.Name, req.Group, req.Path, strings.Join(req.Methods, "|"), req.ReqContentType, req.Description, req.Id)
	if err != nil {
		fmt.Printf("\"2\": %v\n", "2")
		return err
	}

	// 删除原有api_params中旧数据
	_, err = tx.Exec(`
		DELETE FROM "api_param" WHERE "api_id" = ?;
		`, req.Id)
	if err != nil {
		fmt.Printf("\"3\": %v\n", "3")
		return err
	}

	// 创建数据api_params
	for _, param := range req.Params {
		id := utils.UUID()
		mapReqDatas := map[string]any{}
		for _, reqData := range param.ReqDatas {
			mapReqDatas[reqData.Key] = reqData.Value
		}
		reqDataS, err := json.Marshal(mapReqDatas)
		if err != nil {
			return err
		}
		_, err = tx.Exec(consts.PreSqlCreateApiParam,
			id, req.Id,
			strings.Join(param.Route, ","),
			string(reqDataS),
			param.ResData,
			param.ResContentType,
			param.ResCode,
		)
		if err != nil {
			fmt.Printf("\"4\": %v\n", "4")
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		fmt.Printf("\"5\": %v\n", "5")
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
	SELECT "id", "api_id", "route", "req_data", "res_data", "res_content_type", "res_code"
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
			&reqDataS, &apiParam.ResData, &apiParam.ResContentType, &apiParam.ResCode)
		if err != nil {
			return nil, err
		}
		reqDataMap := map[string]any{}
		err = json.Unmarshal([]byte(reqDataS), &reqDataMap)
		if err != nil {
			return nil, err
		}
		reqDatas := []dto.ManageParamReqData{}
		for k, v := range reqDataMap {
			reqDatas = append(reqDatas, dto.ManageParamReqData{
				Key:   k,
				Value: fmt.Sprint(v),
			})
		}

		route := []string{}
		if len(routeS) > 0 {
			route = strings.Split(routeS, ",")
		}

		params = append(params, dto.ManageInfoParam{
			Id:             apiParam.Id,
			ReqDatas:       reqDatas,
			ResCode:        apiParam.ResCode,
			ResContentType: apiParam.ResContentType,
			ResData:        apiParam.ResData,
			Route:          route,
		})
	}

	// 查询api_data
	apiDatas, err := global.DB.Query(`
	SELECT "id", "name", "group", "path", "methods", "req_content_type", "description"
	FROM api_data WHERE id = ?;`, req.Id)
	if err != nil {
		return nil, err
	}
	defer apiDatas.Close()

	var apiData model.ApiData
	apiDatas.Next()
	err = apiDatas.Scan(&apiData.Id, &apiData.Name,
		&apiData.Group, &apiData.Path, &apiData.Methods,
		&apiData.ReqContentType, &apiData.Description)
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
		Description:    apiData.Description,
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

func (*manageLogic) Groups(c *gin.Context) ([]string, error) {
	groups, err := global.DB.Query(`
	SELECT distinct "group" FROM "api_data";`)
	if err != nil {
		return nil, err
	}
	defer groups.Close()

	res := []string{}
	for groups.Next() {
		var group string
		err = groups.Scan(&group)
		if err != nil {
			return nil, err
		}
		res = append(res, group)
	}

	return res, nil
}
