package model

import "time"

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

type Request struct {
	Id         string    `json:"id"`
	ReqTime    time.Time `json:"req_time"`
	IsMatch    bool      `json:"is_match"`
	ReqHeaders string    `json:"req_headers"`
	ReqData    string    `json:"req_data"`
	ReqPath    string    `json:"req_path"`
	ReqMethod  string    `json:"req_method"`
	ResHeaders string    `json:"res_headers"`
	ResData    string    `json:"res_data"`
	ResStatus  int       `json:"res_status"`
	ApiID      string    `json:"api_id"`
}
