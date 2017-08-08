package db

import (
	"time"

	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	"lotteryGame/config"
)

var db *sqlx.DB

func init() {
	var err error
	var mysqlConnectString = config.Mysqlconnstring
	mysqlConnectString += "/" + config.Dbname + "?charset=utf8"
	db, err = sqlx.Connect("mysql", mysqlConnectString)
	if err != nil {
		log.Fatal("connect mysql error :", err)
	}
	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(10)
	db.SetConnMaxLifetime(5 * time.Minute)
	log.Println("connect mysql success..")
}

func GetDb() *sqlx.DB {
	return db
}
