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

func (this *SignupProcessor) Handler(db *sql.DB) (code int, err error) { //注册事件方法，添加注册信息到mysql_user表
	// defer db.Close()

	data, err := db.Query("SELECT mail FROM users WHERE mail = ?;", this.Mail)
	defer data.Close()

	if err != nil {
		code = model.ServerError
		return
	} else if data.Next() {
		code = model.MailHasCreated
		return
	}

	_, err = db.Exec("INSERT INTO users(mail, passwd, name) VALUES (?, ?, ?);", this.Mail, this.Passwd, this.Name)

	if err != nil {
		code = model.ServerError
	} else {
		code = model.RequestSuccess
	}
	return
}
