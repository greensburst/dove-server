package dbi

import "database/sql"

/*
	database interface 数据库接口
*/

type SignupDbi struct { //注册事件接口，该事件需要邮箱和密码两个字段
	mail   string
	passwd string
}

func (this *SignupDbi) Handler(db *sql.DB) { //注册事件方法，添加注册信息到mysql_user表

	db.Exec("INSERT INTO users(mail, passwd) VALUES (?, ?)", this.mail, this.passwd)
	defer db.Close()
}
