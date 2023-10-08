package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"math/rand"
	"net/http"
	"time"

	"arsymedes.github.com/hacktiv-assignment3-05/controller"
	"arsymedes.github.com/hacktiv-assignment3-05/database"
	"arsymedes.github.com/hacktiv-assignment3-05/model"
	"github.com/gin-gonic/gin"
)

func main() {
	database.StartDB()
	router := gin.Default()

	router.GET("/weather", controller.GetWeather)
	router.PUT("/weather", controller.UpdateWeather)

	go updates()
	router.Run()
}

func updates() {
	for {
		data := model.Weather{
			Water: rand.Intn(101),
			Wind:  rand.Intn(101),
		}

		payload, err := json.Marshal(data)
		if err != nil {
			log.Fatal(err)
		}

		url := "http://localhost:8080/weather"
		request, err := http.NewRequest("PUT", url, bytes.NewBuffer(payload))
		if err != nil {
			log.Fatal(err)
		}
		request.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		response, err := client.Do(request)
		if err != nil {
			log.Fatal(err)
		}
		defer response.Body.Close()

		body, err := io.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}

		var bodyMap map[string]interface{}
		err = json.Unmarshal([]byte(body), &bodyMap)
		if err != nil {
			log.Fatal(err)
		}

		log.Println(bodyMap["weather"])
		log.Println("status water: ", bodyMap["status_water"])
		log.Println("status wind: ", bodyMap["status_wind"])

		time.Sleep(15 * time.Second)
	}
}
