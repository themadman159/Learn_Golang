package main

import (
	"fmt"

	"github.com/themadman159/go-tutor/config"
	databases "github.com/themadman159/go-tutor/databases"
	"github.com/themadman159/go-tutor/entities"
)

func main() {
	conf := config.ConfigGetting()
	db := databases.NewPostgresDB(conf.Database)

	fmt.Println(db.ConnectionGetting())

	playerMigation(db)
	adminMigation(db)
	itemMigation(db)
	playerCoinMigation(db)
	inventoryMigation(db)
	purchaseHistoryMigation(db)
}

func playerMigation(db databases.Database) {
	db.ConnectionGetting().Migrator().CreateTable(&entities.Player{})
}

func adminMigation(db databases.Database) {
	db.ConnectionGetting().Migrator().CreateTable(&entities.Admin{})
}

func itemMigation(db databases.Database) {
	db.ConnectionGetting().Migrator().CreateTable(&entities.Item{})
}

func playerCoinMigation(db databases.Database) {
	db.ConnectionGetting().Migrator().CreateTable(&entities.PlayerCoin{})
}

func inventoryMigation(db databases.Database) {
	db.ConnectionGetting().Migrator().CreateTable(&entities.Inventory{})
}

func purchaseHistoryMigation(db databases.Database) {
	db.ConnectionGetting().Migrator().CreateTable(&entities.PurchaseHistory{})
}
