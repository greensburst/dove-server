package prc

import (
	"database/sql"
	"dove-server/model"
	"strings"
)

type SigninProcessor struct {
	Account string `json:"account"`
	Passwd  string `json:"passwd"`
}

func (this *SigninProcessor) Handler(db *sql.DB) (code int, err error) {
	// defer db.Close()
	isMail := strings.Contains(this.Account, "@")

	var account *sql.Rows
	var passwd *sql.Rows

	if isMail {
		account, err = db.Query("SELECT mail FROM users WHERE mail = ?;", this.Account)
	} else {
		account, err = db.Query("SELECT account FROM users WHERE account = ?;", this.Account)
	}
	defer account.Close()

	if err != nil {
		code = model.ServerError
		return
	} else if !account.Next() {
		code = model.AccountError
		return
	}

	if isMail {
		passwd, err = db.Query("SELECT passwd FROM users WHERE mail = ? AND passwd = ?;", this.Account, this.Passwd)
	} else {
		passwd, err = db.Query("SELECT passwd FROM users WHERE account = ? AND passwd = ?;", this.Account, this.Passwd)
	}

	if err != nil {
		code = model.ServerError
		return
	} else if !passwd.Next() {
		code = model.PasswdError
		return
	}

	code = model.RequestSuccess
	return
}
