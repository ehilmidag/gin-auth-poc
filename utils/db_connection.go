package utils

import (
	"fmt"
	"gorm.io/driver/postgres"
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
	_ "gorm.io/gorm"
	"os"
)

type DB struct {
	Database *gorm.DB
}

func (db *DB) ConnectToDatabase() {
	var err error
	dsn := os.Getenv("DB")
	db.Database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect db")
	}
	fmt.Println("db connected")
}
