package main

import (
	"github.com/hangter-lt/easy_mock/cmd"
	"github.com/hangter-lt/easy_mock/initialize"
)

func main() {

	// 初始化数据库
	initialize.InitDB()

	// 加载数据
	initialize.LoadData()

	cmd.Execute()
}
