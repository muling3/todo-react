package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	DB, err := sql.Open("mysql", "root:password@/todo_db?parseTime=true")
	if err != nil {
		log.Println(err.Error())
	}

	testQueries = New(DB)
	os.Exit(m.Run())
	// DB.SetConnMaxLifetime(time.Minute * 3)
	// DB.SetMaxOpenConns(10)
	// DB.SetMaxIdleConns(10)

}