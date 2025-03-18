package database

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Init(mysqlHost, mysqlDb, mysqlUser, mysqlPassword string) {
	mysqlUrl := fmt.Sprintf("%v:%v@tcp(%v)/%v?parseTime=true", mysqlUser, mysqlPassword, mysqlHost, mysqlDb)
	var err error
	DB, err = sql.Open("mysql", mysqlUrl)
	if err != nil {
		panic(err)
	}
	DB.SetConnMaxLifetime(time.Minute * 3)
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(10)
	fmt.Println("connect database successfully")
}
