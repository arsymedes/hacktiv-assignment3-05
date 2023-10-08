package repository

import (
	"arsymedes.github.com/hacktiv-assignment3-05/database"
	"arsymedes.github.com/hacktiv-assignment3-05/model"
)

func GetWeatherDB(weather *model.Weather) error {
	db := database.GetDB()

	return db.First(&weather, "id = ?", 1).Error
}

func PutWeatherDB(weather *model.Weather) error {
	db := database.GetDB()

	var count int64
	err := db.Model(&weather).Count(&count).Error
	if err != nil {
		return err
	}

	if count == 0 {
		return db.Create(&weather).Error
	}

	return db.Model(&weather).Where("id = ?", 1).Updates(&weather).Error
}