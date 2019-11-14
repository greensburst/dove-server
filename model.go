package main

type user struct {
	doveid string
	mail   string
	passwd string
	name   string
	avatar string
}

type signup_dao struct {
	mail   string
	passwd string
}

func (this *signup_dao) add() {
	
}

const (
	SIGNUP_INFORMATION = iota
	CHAT_INFORMATION
)

type data_header struct {
	message_code int
	target_dove  string
	source_dove  string
}

type data_package struct {
	data_header
	content string
}
