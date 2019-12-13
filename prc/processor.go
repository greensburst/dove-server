package prc

import (
	"database/sql"
	"dove-server/model"
	"encoding/json"
)

type Processor interface {
	Handler(*sql.DB) (model.ResponsePackage, error)
}

func messageFactory(code int) (processor Processor) {
	messageMap := make(map[int]Processor)
	messageMap[model.SignupMessage] = new(SignupProcessor)
	messageMap[model.SigninMessage] = new(SigninProcessor)

	return messageMap[code]
}

func RequestPackageFactory(requestPackage *model.RequestPackage) (processor Processor, err error) {

	processor = messageFactory(requestPackage.Code)
	err = json.Unmarshal([]byte(requestPackage.Body), processor)
	if err != nil {
		return nil, err
	}
	return processor, nil
}
