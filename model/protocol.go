package model

const (
	SignupMessage = 0
	SigninMessage = 1
)

const (
	ServerError    = 500 //服务器错误
	RequestSuccess = 200 //请求成功
	MailHasCreated = 401 //邮箱已被注册
	AccountError   = 402 //账户错误
	PasswdError    = 403 //密码错误
)

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
