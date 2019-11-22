package pool

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var Pool *sql.DB

func InitPool() { //构造函数
	Pool, _ = sql.Open("mysql", "root:love@tcp(49.233.188.145:3306)/dove")
	Pool.SetMaxOpenConns(10)
}
