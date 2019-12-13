package model

type RequestPackage struct {
	/*
		请求码：
			0 - 注册信息
	*/
	Code   int    `json:"code"`
	Target string `json:"target"`
	Source string `json:"source"`
	Body   string `json:"body"`
}

type ResponsePackage struct {
	/*
		响应码：
			200 - 成功
			201 - 请求成功但用户名已存在
			501 - 数据库错误
	*/
	Code int    `json:"code"`
	Body string `json:"Body"`
}
