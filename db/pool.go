package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var Pool *sql.DB

func init() {
	Pool, _ = sql.Open("mysql", "root:Lizhan@521@tcp(49.233.188.145:3306)/dove")
	Pool.SetMaxOpenConns(2000)
	Pool.SetMaxIdleConns(1000)
}
