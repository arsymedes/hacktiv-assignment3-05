package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"arsymedes.github.com/hacktiv-assignment3-05/model"
)

var (
	host     = "localhost"
	user     = "postgres"
	dbPort   = "5432"
	dbName   = "postgres"
	password = "password"
	db       *gorm.DB
	err      error
)

func StartDB() {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbName, dbPort)

	db, err = gorm.Open(postgres.Open(config), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	db.Exec("CREATE  SCHEMA IF NOT EXISTS public;")

	if err != nil {
		log.Fatal("error connecting to database :", err)
	}

	db.Debug().AutoMigrate(model.Weather{})
}

func GetDB() *gorm.DB {
	return db
}
