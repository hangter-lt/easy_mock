package model

type ApiData struct {
	Id             string `json:"id"`
	Name           string `json:"name"`
	Group          string `json:"group"`
	Path           string `json:"path"`
	Methods        string `json:"methods"`
	ReqContentType string `json:"req_content_type"`
	Description    string `json:"description"`
}

type ApiParam struct {
	Id             string         `json:"id"`
	ApiId          string         `json:"api_id"`
	Route          []string       `json:"route"`
	ReqData        map[string]any `json:"req_data"`
	ResData        string         `json:"res_data"`
	ResCode        int            `json:"res_code"`
	ResContentType string         `json:"res_content_type"`
}
