package initialize

import (
	"database/sql"
	"os"

	"github.com/hangter-lt/easy_mock/consts"
	"github.com/hangter-lt/easy_mock/global"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB() error {
	// 创建data文件夹
	os.Mkdir("data", os.ModePerm)

	db, err := sql.Open("sqlite3", "data/db.db")
	if err != nil {
		return err
	}
	global.DB = db

	_, err = db.Exec(consts.SqlCreateApiData)
	if err != nil {
		return err
	}
	_, err = db.Exec(consts.SqlCreateApiParam)
	if err != nil {
		return err
	}

	_, err = db.Exec(consts.SqlCreateRequest)
	if err != nil {
		return err
	}

	return nil
}
