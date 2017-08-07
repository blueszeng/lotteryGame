package db

import (
	"database/sql"

	"log"

	_ "github.com/go-sql-driver/mysql"

	"lotteryGame/config"
)

var db *sql.DB

func init() {
	var err error
	var mysqlConnectString = config.Mysqlconnstring
	mysqlConnectString += "/" + config.Dbname + "?charset=utf8"
	db, err = sql.Open("mysql", mysqlConnectString)
	if err != nil {
		log.Fatal("connect mysql error :", err)
	}
	log.Println("connect mysql success..")
}
