package model

const (
	SIGNUP_MESSAGE = iota
	SIGNIN_MESSAGE
	CHAT_MESSAGE
)

type Package struct {
	Code   int    `json:"code"`
	Target string `json:"target"`
	Source string `json:"source"`
	Body   string `json:"body"`
}
