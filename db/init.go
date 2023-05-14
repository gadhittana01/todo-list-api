package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gadhittana01/todolist/config"
	"github.com/gadhittana01/todolist/helper"
	_ "github.com/go-sql-driver/mysql"
)

func InitDB() *sql.DB {
	var db *sql.DB
	var err error

	config := &config.GlobalConfig{}
	helper.LoadConfig(config)

	dbConn := config.DB

	connString := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s",
		dbConn.User, dbConn.Password, dbConn.Host, dbConn.Port, dbConn.Name)

	db, err = sql.Open("mysql", connString)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("DB " + dbConn.Name + " connected Successfully!")

	return db
}
