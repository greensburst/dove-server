package dbi

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

/*
	database interface 数据库接口
*/

type SignupDbi struct { //注册事件接口，该事件需要邮箱和密码两个字段
	Name   string `json:"name"`
	Mail   string `json:"mail"`
	Passwd string `json:"passwd"`
}

func (this *SignupDbi) Handler() (err error) { //注册事件方法，添加注册信息到mysql_user表

	db, err := sql.Open("mysql", "root:Lizhan@521@tcp(49.233.188.145:3306)/dove")
	if err != nil {
		fmt.Println("连接错误！", err)
	}
	_, err = db.Exec("INSERT INTO users(mail, passwd, name) VALUES (?, ?, ?);", this.Mail, this.Passwd, this.Name)
	if err != nil {
		fmt.Println("出错！", err)
	}
	defer db.Close()
	return err
}
