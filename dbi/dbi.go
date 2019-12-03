package dbi

import (
	"database/sql"
	"encoding/json"
)

type DataBaseInterface struct {
	Pool   *sql.DB
	Stream []byte
}

func NewDataBaseInterface() *DataBaseInterface {
	dbi := new(DataBaseInterface)
	// dbi.Pool, _ = sql.Open("mysql", "root:love@tcp(49.233.188.145:3306)/dove")
	// dbi.Pool.SetMaxOpenConns(10)
	return dbi
}

func (this *DataBaseInterface) SignUp() (err error) {
	su := new(SignupDbi)
	json.Unmarshal(this.Stream, su)
	err = su.Handler()
	return err
}
