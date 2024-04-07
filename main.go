package main

import (
	"github.com/themadman159/go-tutor/config"
	databases "github.com/themadman159/go-tutor/databases"
	"github.com/themadman159/go-tutor/server"
)

func main() {

	conf := config.ConfigGetting()

	db := databases.NewPostgresDB(conf.Database)

	server := server.NewEchoServer(conf, db.ConnectionGetting())

	server.Start()
}
