package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/muling3/go-todos-api/api"
	db "github.com/muling3/go-todos-api/db/sqlc"
	"github.com/muling3/go-todos-api/util"
	"moul.io/banner"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Config not found " + err.Error())
	}

	fmt.Println(banner.Inline("Todos"))

	DB, err := sql.Open(config.DBDriver, config.DBUrl)
	if err != nil {
		log.Println(err.Error())
	}

	queries := db.New(DB)
	server := api.NewServer(queries)

	if err := server.StartServer(config.DBAddress); err != nil {
		log.Fatal("Server Failed to start " + err.Error())
	}
}
