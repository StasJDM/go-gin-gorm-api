package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "host=127.0.0.1 port=5432 user=root dbname=api password=root sslmode=disable"

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Не удалось подключиться к БД")
	}

	err = database.AutoMigrate(&User{})
	if err != nil {
		return
	}

	DB = database
}
