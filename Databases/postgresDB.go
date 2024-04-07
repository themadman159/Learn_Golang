package databases

import (
	"fmt"
	"sync"
	"log"

	"github.com/themadman159/go-tutor/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresDB struct {
	*gorm.DB
}

var (
	postgresDBInstance *PostgresDB
	once               sync.Once
)

func NewPostgresDB(conf *config.Database) Database {

	once.Do(func() {
		dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s search_path=%s",
			conf.Host,
			conf.Port,
			conf.User,
			conf.Password,
			conf.DBName,
			conf.SSLMode,
			conf.Schema,
		)

		conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(err)
		}

		log.Printf("connect to database %s", conf.DBName)
		postgresDBInstance = &PostgresDB{conn}
	})
	return postgresDBInstance
}

func (db *PostgresDB) ConnectionGetting() *gorm.DB {
	return db.DB
}
