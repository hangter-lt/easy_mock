package logic

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/hangter-lt/easy_mock/global"
	"github.com/hangter-lt/easy_mock/model"
	"github.com/hangter-lt/easy_mock/model/dto"
)

type requestLogic struct{}

var Request = requestLogic{}

func (r *requestLogic) RealTime(c *gin.Context) {

	// 增加chan
	rtc := model.RealTimeChan{
		C: make(chan string),
	}

	// 关闭连接时从关闭channel,并从global.ChanList移除
	defer func() {
		close(rtc.C)

		global.ChanListMutex.Lock()
		for i, ch := range global.ChanList {
			if ch == &rtc {
				global.ChanList = append(global.ChanList[:i], global.ChanList[i+1:]...)
				break
			}
		}
		global.ChanListMutex.Unlock()
	}()

	global.ChanList = append(global.ChanList, &rtc)
	fmt.Printf("global.ChanList: %v\n", global.ChanList)

	// set sse headers
	c.Header("Content-Type", "text/event-stream")
	c.Header("Cache-Control", "no-cache")
	c.Header("Connection", "keep-alive")

	for _, id := range global.CircularQueue.GetSlice() {
		reqData, err := r.Info(c, &dto.RequestInfoReq{Id: id})
		if err != nil {
			return
		}

		res := dto.RequestRealtimeRes{
			Id:        reqData.Id,
			ReqMethod: reqData.ReqMethod,
			ReqPath:   reqData.ReqPath,
			ReqTime:   reqData.ReqTime,
			IsMatch:   reqData.IsMatch,
		}
		resB, err := json.Marshal(res)
		if err != nil {
			return
		}

		c.Writer.WriteString(fmt.Sprintf("id: %s\ndata: %s\n\n", id, string(resB)))
		c.Writer.Flush()
	}

	for {
		select {
		case id, ok := <-rtc.C:
			if !ok {
				return
			}
			reqData, err := r.Info(c, &dto.RequestInfoReq{Id: id})
			if err != nil {
				return
			}

			res := dto.RequestRealtimeRes{
				Id:        reqData.Id,
				ReqMethod: reqData.ReqMethod,
				ReqPath:   reqData.ReqPath,
				ReqTime:   reqData.ReqTime,
				IsMatch:   reqData.IsMatch,
			}
			resB, err := json.Marshal(res)
			if err != nil {
				return
			}
			c.Writer.WriteString(fmt.Sprintf("id: %s\ndata: %s\n\n", id, string(resB)))
			c.Writer.Flush()
		case <-c.Request.Context().Done():
			return
		}
	}

	// TODO: 前端停止接收时处理逻辑
}

func (*requestLogic) Info(c *gin.Context, req *dto.RequestInfoReq) (*dto.RequestInfoRes, error) {

	// 请求表查询请求数据
	reqData, err := global.DB.Query(`
	SELECT "id", "api_id", "is_match", "req_data", "req_headers", "req_addr", 
	 "req_method", "req_path", "req_time", "res_data", "res_headers", "res_status"
	FROM requests WHERE id = ?;`, req.Id)
	if err != nil {
		return nil, err
	}
	defer reqData.Close()

	var res dto.RequestInfoRes
	for reqData.Next() {
		err = reqData.Scan(&res.Id, &res.ApiID, &res.IsMatch, &res.ReqData,
			&res.ReqHeaders, &res.ReqAddr, &res.ReqMethod, &res.ReqPath,
			&res.ReqTime, &res.ResData, &res.ResHeaders, &res.ResStatus)
		if err != nil {
			return nil, err
		}
	}

	return &res, nil
}
