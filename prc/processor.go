package prc

import (
	"dove-server/dbi"
	"dove-server/model"
	"net"
)

func Processor(pkg *model.Package, conn net.Conn) {

	defer conn.Close()
	dbi := dbi.NewDataBaseInterface()
	dbi.Stream = []byte(pkg.Body)

	switch pkg.Code {

	case model.SIGNUP_MESSAGE:
		err := dbi.SignUp()
		if err != nil {
			conn.Write([]byte("0"))
		} else {
			conn.Write([]byte("1"))
		}
	}
}
