package prc

import (
	"dove-server/model"
	"encoding/json"
	"errors"
)

type Processor interface {
	Handler() ([]byte, error)
}

func RequestPackageFactory(requestPackage *model.RequestPackage) (processor Processor, err error) {

	hasError := false
	switch requestPackage.Code {
	case model.SignupMessage:
		processor = new(SignupProcessor)
		err = json.Unmarshal([]byte(requestPackage.Body), processor)
		if err != nil {
			hasError = true
		}
	default:
		return nil, errors.New("response package type is not founded.")
	}

	if hasError {
		return nil, err
	} else {
		return processor, nil
	}
}
