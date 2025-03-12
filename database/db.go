package database

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Database struct {
	*gorm.DB
}

var dbInstance *Database

func GetDatabaseInstance() *Database {
	databaseURL := os.Getenv("DATABASE_URL")

	if dbInstance == nil {
		db, err := gorm.Open(postgres.Open(databaseURL), &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				// SingularTable: true,
			},
		})
		if err != nil {
			panic("failed to connect to database")
		}

		dbInstance = &Database{DB: db}
	}

	return dbInstance
}
