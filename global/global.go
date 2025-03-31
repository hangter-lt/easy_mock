package global

import (
	"database/sql"
	"sync"

	"github.com/hangter-lt/easy_mock/model"
)

var (
	// 数据库
	DB *sql.DB
	// 存放所有的api数据信息
	ApiData map[string]model.ApiData
	// 存正则编码后的请求参数和参数对的关系
	ReqParam map[string][]string
	// 存放参数id和内容的对应关系
	ApiParam map[string]model.ApiParam
	// 存放chan结构体指针
	ChanList      []*model.RealTimeChan
	ChanListMutex sync.RWMutex
	// 循环队列
	CircularQueue *model.CircularQueue
)
