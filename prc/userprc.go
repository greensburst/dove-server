package prc

/*
	processor 处理器
*/

import (
	"dove-server/dbi"
	"dove-server/model"
	"encoding/json"
)

type UserPrc struct { //对user_model的处理器
	model.User
}

func (this *UserPrc) Signup(body string) { //user的注册事件

	sd := new(dbi.SignupDbi) //实例化一个注册接口
	stream := []byte(body)
	json.Unmarshal(stream, sd) //为注册接口赋值

	sd.Add() //执行数据库添加操作
}

func (this *UserPrc) Signin(body string) {

	
}
