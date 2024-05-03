package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/thuanan7/SimpleBank-Golang/api"
	db "github.com/thuanan7/SimpleBank-Golang/db/sqlc"
	"github.com/thuanan7/SimpleBank-Golang/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server")
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start the server")
	}
}
