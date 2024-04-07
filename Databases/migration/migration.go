package main

import (
	"fmt"

	"github.com/themadman159/go-tutor/config"
	databases "github.com/themadman159/go-tutor/databases"
)

func main() {
	conf := config.ConfigGetting()
	db := databases.NewPostgresDB(conf.Database)

	fmt.Println(db.ConnectionGetting())

}