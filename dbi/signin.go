package dbi

import "database/sql"

type SigninDbi struct {
	Account string
	Passwd  string
}

func (this *SigninDbi) Handler() {

	db, _ := sql.Open("mysql", "root:love@tcp(49.233.188.145:3306)/dove")
	defer db.Close()
}
