package consts

// 建表
const (
	SqlCreateApiData = `
		CREATE TABLE IF NOT EXISTS "api_data" (
			"id" text(32) NOT NULL,
			"name" TEXT NOT NULL,
			"group" TEXT NOT NULL,
			"path" TEXT NOT NULL,
			"methods" TEXT NOT NULL,
			"req_content_type" TEXT,
			"description" TEXT,
			PRIMARY KEY ("id" DESC)
		);
	`

	SqlCreateApiParam = `
		CREATE TABLE IF NOT EXISTS "api_param" (
			"id" text(32) NOT NULL,
			"api_id" text(32) NOT NULL,
			"route" TEXT,
			"req_data" TEXT,
			"res_content_type" TEXT NOT NULL,
			"res_data" TEXT,
			"res_code" INTEGER NOT NULL,
			PRIMARY KEY ("id")
		);
	`
)

// sql预编译
const (
	PreSqlCreateApiData = `
INSERT INTO api_data ("id", "name", "group", "path", "methods", "req_content_type", "description")
VALUES (?, ?, ?, ?, ?, ?, ?)
`
	PreSqlCreateApiParam = `
INSERT INTO api_param ("id", "api_id", "route", "req_data", "res_data", "res_content_type", "res_code")
VALUES (?, ?, ?, ?, ?, ?, ?)
`
)
