package prc

import (
	"database/sql"
	"dove-server/model"
	"encoding/json"

	_ "github.com/go-sql-driver/mysql"
)

type SignupProcessor struct { //注册事件接口，该事件需要昵称邮箱和密码三个字段
	Name   string `json:"name"`
	Mail   string `json:"mail"`
	Passwd string `json:"passwd"`
}

func (this *SignupProcessor) Handler() (res []byte, err error) { //注册事件方法，添加注册信息到mysql_user表

	responsePackage := new(model.ResponsePackage)
	db, err := sql.Open("mysql", "root:Lizhan@521@tcp(49.233.188.145:3306)/dove")
	defer db.Close()
	if err != nil {
		responsePackage.Code = 501
	}
	_, err = db.Exec("INSERT INTO users(mail, passwd, name) VALUES (?, ?, ?);", this.Mail, this.Passwd, this.Name)
	if err != nil {
		responsePackage.Code = 201
	} else {
		responsePackage.Code = 200
	}
	res, _ = json.Marshal(responsePackage)

	return
}
