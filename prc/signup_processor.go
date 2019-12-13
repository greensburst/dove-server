package prc

import (
	"database/sql"
	"dove-server/model"
)

type SignupProcessor struct { //注册事件接口，该事件需要昵称邮箱和密码三个字段
	Name   string `json:"name"`
	Mail   string `json:"mail"`
	Passwd string `json:"passwd"`
}

func (this *SignupProcessor) Handler(db *sql.DB) (responsePackage model.ResponsePackage, err error) { //注册事件方法，添加注册信息到mysql_user表

	data, err := db.Query("SELECT mail FROM users WHERE mail = ?;", this.Mail)
	defer data.Close()

	if err != nil {
		responsePackage.Code = model.ServerError
		return
	} else if data.Next() {
		responsePackage.Code = model.MailHasCreated
		return
	}

	_, err = db.Exec("INSERT INTO users(mail, passwd) VALUES (?, ?);", this.Mail, this.Passwd)
	_, err = db.Exec("INSERT INTO user_info(account,name) VALUES((SELECT account FROM users WHERE mail = ?),?);", this.Mail, this.Name)

	if err != nil {
		responsePackage.Code = model.ServerError
	} else {
		responsePackage.Code = model.RequestSuccess
	}
	return
}
