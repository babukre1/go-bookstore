package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	// "gorm.io/gorm/logger"
)

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open(postgres.Open("postgres""host=localhost port=5432 user=postgres dbname=bookstore password='00aa1122' sslmode=disable"))
	if err != nil {
		panic(err)
	}

	db = d
}

func GetDB() *gorm.DB {
	return db
}
