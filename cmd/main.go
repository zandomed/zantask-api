package main

import (
	"github.com/zandomed/zantask-api/config"
	"github.com/zandomed/zantask-api/database"
	"github.com/zandomed/zantask-api/server"
)

func main() {
	conf := config.GetConfig()
	db := database.NewPostgresDatabase(conf)
	server.NewFiberServer(conf, db).Start()
}
