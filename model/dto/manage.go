package dto

type ManageCreateReq struct {
	Name           string              `json:"name"`
	Group          string              `json:"group"`
	Path           string              `json:"path"`
	Methods        []string            `json:"methods"`
	Params         []ManageCreateParam `json:"params"`
	ReqContentType string              `json:"req_content_type"`
}

type ManageCreateParam struct {
	Route          []string       `json:"route"`
	ReqData        map[string]any `json:"req_data"`
	ResData        string         `json:"res_data"`
	ResContentType string         `json:"res_content_type"`
}

type ManageUpdateReq struct {
	Id string `json:"id"`
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
}

type ManageInfoParam struct {
	Id             string         `json:"id"`
	Route          []string       `json:"route"`
	ReqData        map[string]any `json:"req_data"`
	ResData        string         `json:"res_data"`
	ResContentType string         `json:"res_content_type"`
}

type ManageDeleteReq struct {
	Id string `uri:"id"`
}
