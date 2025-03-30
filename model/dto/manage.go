package dto

type ManageCreateReq struct {
	Name           string              `json:"name" binding:"required"`
	Group          string              `json:"group" binding:"required"`
	Path           string              `json:"path" binding:"required"`
	Methods        []string            `json:"methods" binding:"required"`
	Params         []ManageCreateParam `json:"params" binding:"required"`
	ReqContentType string              `json:"req_content_type"`
	Description    string              `json:"description"`
}

type ManageCreateRes struct {
	Id string `json:"id"`
}

type ManageParamReqData struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type ManageCreateParam struct {
	Route          []string             `json:"route"`
	ReqDatas       []ManageParamReqData `json:"req_datas"`
	ResData        string               `json:"res_data" binding:"required"`
	ResCode        int                  `json:"res_code" binding:"required"`
	ResContentType string               `json:"res_content_type" binding:"required"`
}

type ManageUpdateReq struct {
	Id string `json:"id" binding:"required"`
	ManageCreateReq
}

type ManageListRes map[string][]ManageList

type ManageList struct {
	Id             string   `json:"id"`
	Name           string   `json:"name"`
	Group          string   `json:"group"`
	Methods        []string `json:"methods"`
	Path           string   `json:"path"`
	ReqContentType string   `json:"req_content_type"`
}

type ManageInfoReq struct {
	Id string `uri:"id"`
}
type ManageInfoRes struct {
	Id             string            `json:"id"`
	Name           string            `json:"name"`
	Group          string            `json:"group"`
	Path           string            `json:"path"`
	Methods        []string          `json:"methods"`
	ReqContentType string            `json:"req_content_type"`
	Params         []ManageInfoParam `json:"params"`
	Description    string            `json:"description"`
}

type ManageInfoParam struct {
	Id             string               `json:"id"`
	Route          []string             `json:"route"`
	ReqDatas       []ManageParamReqData `json:"req_datas"`
	ResCode        int                  `json:"res_code"`
	ResData        string               `json:"res_data"`
	ResContentType string               `json:"res_content_type"`
}

type ManageDeleteReq struct {
	Id string `uri:"id"`
}
