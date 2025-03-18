package database

import (
	"database/sql"
	"fmt"
	"log"
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
	fmt.Println("connect mysql successfully")
	createTable()
}

func createTable() {

	createUsersTable := `
		create table if not exists tasks(
			id int primary key auto_increment,
			description varchar(255),
			is_done bool default false
		)
	`
	_, err := DB.Exec(createUsersTable)
	if err != nil {
		log.Fatal(err)
		panic("could not create tasks table")
	}
	fmt.Println("create table tasks successfully")
}
