package main

import "encoding/json"

type processor struct {
}

func (this *processor) signup(content string) {

	signup_dao := new(signup_dao)
	data := []byte(content)
	json.Unmarshal(data, signup_dao)
	signup_dao.add()
}
