package models

import (
	"github.com/jinzhu/gorm"
	// _ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
)

var DB gorm.DB

func Init(database string, connectionString string) error {
	var err error
	DB, err = gorm.Open(database, connectionString)

	if err != nil {
		return err

	}

	DB.CreateTable(Location{})
	DB.CreateTable(Thumbnail{})
	DB.CreateTable(User{})
	DB.CreateTable(UserProfile{})
	DB.CreateTable(Video{})

	return nil
}
