package model

const ( //消息类型
	SignupMessage = 0
	SigninMessage = 1
)

const ( //状态码
	ServerError    = 500 //服务器错误
	RequestSuccess = 200 //请求成功
	MailHasCreated = 401 //邮箱已被注册
	AccountError   = 402 //账户错误
	PasswdError    = 403 //密码错误
)
