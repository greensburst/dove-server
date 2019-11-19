package prc

import (
	"dove-server/dbi"
	"dove-server/model"
	"encoding/json"
)

type Processor interface {
	Handler()
}

func ProcessorFactory(pkg *model.Package) Processor {

	switch pkg.Header.Code {
	case model.SIGNUP_MESSAGE:
		sudbi := new(dbi.SignupDbi)
		stream := []byte(pkg.Body)
		json.Unmarshal(stream, sudbi)
		return sudbi
	default:
		return nil
	}
}
