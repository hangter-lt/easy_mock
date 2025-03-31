package logic

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/hangter-lt/easy_mock/global"
	"github.com/hangter-lt/easy_mock/model"
)

type requestLogic struct{}

var Request = requestLogic{}

func (*requestLogic) RealTime(c *gin.Context) {

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
		c.Writer.WriteString(fmt.Sprintf("data: %s\n\n", id))
		c.Writer.Flush()
	}

	for {
		select {
		case id, ok := <-rtc.C:
			if !ok {
				return
			}
			c.Writer.WriteString(fmt.Sprintf("data: %s\n\n", id))
			c.Writer.Flush()
		case <-c.Request.Context().Done():
			fmt.Printf("\"a\": %v\n", "a")
			return
		}
	}

}
