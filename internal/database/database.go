package database

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func MakeDatabase() Database {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&User{})

	fmt.Println("database connected successfully")

	return Database{
		db: db,
	}
}

func (db *Database) CreateUsers(name string, email string, password string) {
	user := User{Name: name, Email: email, Password: password}
	result := db.db.Create(&user)
	fmt.Println(result.RowsAffected)
}
