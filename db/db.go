package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func Init() {

	// username:password@protocol(address)/dbname?param=value
	//connectionString := conf.DB_USERNAME + ":" + conf.DB_PASSWORD + "@tcp(" + conf.DB_HOST + ":" + conf.DB_PORT + ")/" + conf.DB_NAME
	connectionString := "root:root@keuangan_service_mysql/keuangan"
	fmt.Println(connectionString)

	db, err = sql.Open("mysql", connectionString)
	if err != nil {
		panic("connection string error")
	}

	err = db.Ping()
	if err != nil {
		panic("DSN Invalid")
	}
}

func CreateCon() *sql.DB {
	return db
}
