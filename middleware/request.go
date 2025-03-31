package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hangter-lt/easy_mock/global"
	"github.com/hangter-lt/easy_mock/utils"
)

type responseRecorder struct {
	gin.ResponseWriter
	body   *bytes.Buffer
	status int
}

func (r *responseRecorder) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}

func (r *responseRecorder) WriteHeader(statusCode int) {
	r.status = statusCode
	r.ResponseWriter.WriteHeader(statusCode)
}

func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now().UnixMilli()
		reqID := utils.UUID()

		reqData := make(map[string]any)
		contentType := c.ContentType()
		fmt.Printf("contentType: %v\n", contentType)

		// 根据内容类型解析 Body
		switch contentType {
		case "application/json":
			if err := c.ShouldBindJSON(&reqData); err != nil {
				fmt.Printf("err: %v\n", err)
				return
			}
		case "application/x-www-form-urlencoded":
			if err := c.Request.ParseForm(); err == nil {
				for k, v := range c.Request.PostForm {
					reqData[k] = v
				}
			}
		case "multipart/form-data":
			if err := c.Request.ParseMultipartForm(10 << 20); err == nil {
				for k, v := range c.Request.MultipartForm.Value {
					reqData[k] = v
				}
			}
		}

		// 合并 URL 查询参数
		// TODO: 优化数组参数
		for k, v := range c.Request.URL.Query() {
			if len(v) == 1 {
				reqData[k] = v[0]
			} else {
				reqData[k] = v
			}
		}

		reqDataB, err := json.Marshal(reqData)
		if err != nil {
			return
		}
		c.Set("req_data", reqDataB)

		// 包装响应记录器
		recorder := &responseRecorder{
			ResponseWriter: c.Writer,
			body:           bytes.NewBuffer(nil),
			status:         http.StatusOK,
		}
		c.Writer = recorder

		// 处理请求
		c.Next()

		// 准备插入数据
		stmt, _ := global.DB.Prepare(`
			INSERT INTO requests(
				id, req_time, is_match, req_headers, req_data, 
				req_path, req_method, res_headers, res_data, res_status, req_addr
			) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
		`)

		isMatch := c.GetBool("is_match")

		// 处理 headers
		reqHeaders, _ := json.Marshal(c.Request.Header)
		resHeaders, _ := json.Marshal(c.Writer.Header())

		// 执行插入
		_, err = stmt.Exec(
			reqID,
			start,
			isMatch,
			string(reqHeaders),
			string(reqDataB),
			c.Request.URL.Path,
			c.Request.Method,
			string(resHeaders),
			recorder.body.String(),
			recorder.status,
			c.Request.RemoteAddr,
		)

		if err != nil {
			fmt.Printf("Failed to save request: %v\n", err)
		}

		global.ChanListMutex.RLock()
		for _, cl := range global.ChanList {
			cl.C <- reqID
		}
		global.ChanListMutex.RUnlock()

		fmt.Printf("reqID: %v\n", reqID)
		global.CircularQueue.Push(reqID)
	}
}
