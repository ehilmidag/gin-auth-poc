package utils

import (
	"fmt"
	"github.com/ehilmidag/models"
	"gorm.io/driver/postgres"
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
	_ "gorm.io/gorm"
	"os"
)

var DB *gorm.DB

func ConnectToDatabase() {
	var err error
	dsn := os.Getenv("DB")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect db")
	}
	fmt.Println("db connected")
}

func SyncDatabase() {
	err := DB.AutoMigrate(&models.User{})
	if err != nil {
		panic("initialize edemedim dayı")
	}
}
