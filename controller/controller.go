package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"arsymedes.github.com/hacktiv-assignment3-05/helper"
	"arsymedes.github.com/hacktiv-assignment3-05/model"
	"arsymedes.github.com/hacktiv-assignment3-05/repository"
)

func GetWeather(ctx *gin.Context) {
	var weather model.Weather

	if err := repository.GetWeatherDB(&weather); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"weather":      weather,
		"status_water": helper.WindStatus(weather.Wind),
		"status_wind":  helper.WindStatus(weather.Water),
	})
}

func UpdateWeather(ctx *gin.Context) {
	var weather model.Weather

	err := ctx.ShouldBindJSON(&weather)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := repository.PutWeatherDB(&weather); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"weather":      weather,
		"status_water": helper.WindStatus(weather.Wind),
		"status_wind":  helper.WindStatus(weather.Water),
	})
}
