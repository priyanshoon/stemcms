package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

func MakeDatabase() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{
		QueryFields: true,
	})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&User{})

	log.Println("database connected successfully")

	return db
}
