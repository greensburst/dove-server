package dbi

/*
	database interface 数据库接口
*/

import (
	"database/sql"
)

type SignupDbi struct { //注册事件接口，该事件需要邮箱和密码两个字段
	mail   string
	passwd string
}

func (this *SignupDbi) Handler() { //注册事件方法，添加注册信息到mysql_user表

	db, _ := sql.Open("mysql", "root:love@tcp(49.233.188.145:3306)/dove")
	defer db.Close()
	db.Exec("INSERT INTO users(mail, passwd) VALUES (?, ?)", this.mail, this.passwd)
}
