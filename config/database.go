package config

import (
	"os"

	"github.com/MatBas09/API-in-golang/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var dbPath string = "../data/db.db"
var dbFolder string = "../data/"
var DB *gorm.DB

func creatDbFolder() {
	err := os.MkdirAll(dbFolder, os.ModePerm)
	if err != nil {
		panic("can´t creat the db folder")
	}
}

func ConnectDatabase() {
	_, err := os.Stat(dbFolder)
	if os.IsNotExist(err) {
		creatDbFolder()
	} else {
		panic("one erro has occorrude trying to get info about db folder")
	} 

	DB, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		panic("can´t open the database")
	}

	DB.AutoMigrate(&models.MenuItem{})
}
