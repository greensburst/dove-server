package prc

import (
	"database/sql"
	"dove-server/model"
	"encoding/json"
	"fmt"
	"strings"
)

type SigninProcessor struct {
	Account string `json:"account"`
	Passwd  string `json:"passwd"`
}

func (this *SigninProcessor) Handler(db *sql.DB) (responsePackage model.ResponsePackage, err error) {
	isMail := strings.Contains(this.Account, "@")

	var rows *sql.Rows
	var user model.User
	var friends []model.Friend
	var local model.Local

	if isMail {
		rows, err = db.Query("SELECT account FROM users WHERE mail = ?;", this.Account)
	} else {
		rows, err = db.Query("SELECT account FROM users WHERE account = ?;", this.Account)
	}

	if err != nil {
		responsePackage.Code = model.ServerError
		return
	}

	hasAccount := 0
	for rows.Next() {
		hasAccount = hasAccount + 1
		err = rows.Scan(&this.Account)
		if err != nil {
			responsePackage.Code = model.ServerError
			return
		}
	}

	if hasAccount == 0 {
		responsePackage.Code = model.AccountError
		return
	}
	rows.Close()

	rows, err = db.Query("SELECT passwd FROM users WHERE account = ? AND passwd = ?;", this.Account, this.Passwd)
	fmt.Println(this.Account)
	if err != nil {
		responsePackage.Code = model.ServerError
		return
	} else if !rows.Next() {
		responsePackage.Code = model.PasswdError
		return
	}
	rows.Close()

	rows, err = db.Query("SELECT users.account,mail,name,avatar,gender,region FROM users,user_info WHERE users.account = user_info.account AND users.account = ?;", this.Account)
	if err != nil {
		responsePackage.Code = model.ServerError
		return
	}
	for rows.Next() {
		err = rows.Scan(&user.Account, &user.Mail, &user.Name, &user.Avatar, &user.Gender, &user.Region)
		if err != nil {
			fmt.Println("111")
			responsePackage.Code = model.ServerError
			return
		}
	}
	rows.Close()

	rows, err = db.Query(`SELECT users.account,mail,name,avatar,gender,region,star FROM users,user_info,friendship
						WHERE users.account = user_info.account
						AND friendship.friend_id = users.account AND friendship.user_id = ?
						AND users.account IN (
							SELECT friend_id FROM friendship WHERE user_id = ?
						);`, this.Account, this.Account)
	if err != nil {
		fmt.Println("222")
		responsePackage.Code = model.ServerError
		return
	}
	for rows.Next() {
		var friend model.Friend
		_ = rows.Scan(&friend.Account, &friend.Mail, &friend.Name, &friend.Avatar, &friend.Gender, &friend.Region, &friend.Star)
		if err != nil {
			responsePackage.Code = model.ServerError
			return
		}
		friends = append(friends, friend)
	}
	rows.Close()

	local.User = user
	local.Friends = friends

	res, err := json.Marshal(local)
	if err != nil {
		responsePackage.Code = model.ServerError
		return
	}

	responsePackage.Code = model.RequestSuccess
	responsePackage.Body = string(res)
	fmt.Println(local)
	return
}
