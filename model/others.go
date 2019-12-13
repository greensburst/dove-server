package model

type User struct {
	Account int
	Mail    string
	Name    string
	Avatar  string //存头像在服务器的链接
	Gender  string
	Region  string
}

type Friend struct {
	User
	Star int
}

type Local struct {
	User
	Friends []Friend
}
