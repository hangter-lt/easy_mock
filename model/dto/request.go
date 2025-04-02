package dto

type RequestInfoReq struct {
	Id string `uri:"id"`
}

type RequestInfoRes struct {
	Id         *string `json:"id"`
	ReqTime    *int64  `json:"req_time"`
	IsMatch    *bool   `json:"is_match"`
	ReqAddr    *string `json:"req_addr"`
	ReqHeaders *string `json:"req_headers"`
	ReqData    *string `json:"req_data"`
	ReqPath    *string `json:"req_path"`
	ReqMethod  *string `json:"req_method"`
	ResHeaders *string `json:"res_headers"`
	ResData    *string `json:"res_data"`
	ResStatus  *int    `json:"res_status"`
	ApiID      *string `json:"api_id"`
}

type RequestRealtimeRes struct {
	Id        *string `json:"id"`
	IsMatch   *bool   `json:"is_match"`
	ReqMethod *string `json:"req_method"`
	ReqPath   *string `json:"req_path"`
	ReqTime   *int64  `json:"req_time"`
}
