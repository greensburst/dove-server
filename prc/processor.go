package prc

import (
	"dove-server/dbi"
	"dove-server/model"
)

func Processor(pkg *model.Package) {

	dbi := dbi.NewDataBaseInterface()
	dbi.Stream = []byte(pkg.Body)
	

	switch pkg.Header.Code {

	case model.SIGNUP_MESSAGE:
		dbi.SignUp()
	}
}
