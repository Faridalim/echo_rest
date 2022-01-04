package models

import (
	"database/sql"
	"echo_rest/db"
	"echo_rest/helpers"
	"fmt"
)

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
}

func CheckLogin(username, password string) (bool, error) {
	var obj User
	var pwd string

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM users WHERE username = ?"
	err := con.QueryRow(sqlStatement, username).Scan(&obj.Id, &obj.Username, &pwd)

	if err == sql.ErrNoRows {
		fmt.Println("Username not found")
		return false, err
	}

	if err != nil {
		return false, err
	}

	match, err := helpers.CheckPasswordHash(password, pwd)
	if !match {
		fmt.Println("Wrong password")
		return false, err
	}

	return true, nil
}
