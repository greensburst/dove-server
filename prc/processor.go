package prc

import (
	"dove-server/model"
	"encoding/json"
)

type Processor interface {
	Handler() ([]byte, error)
}

func RequestPackageFactory(requestPackage *model.RequestPackage) Processor {

	switch requestPackage.Code {
	case model.SignupMessage:
		processor := new(SignupProcessor)
		json.Unmarshal([]byte(requestPackage.Body), processor)
		return processor
	default:
		return nil
	}

}
