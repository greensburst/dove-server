package model

const (
	SIGNUP_MESSAGE = iota
	SIGNIN_MESSAGE
	CHAT_MESSAGE
)

type Header struct {
	Code   int
	Target string
	Source string
}

type Package struct {
	Header
	Body string
}
