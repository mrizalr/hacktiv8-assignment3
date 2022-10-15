package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"
)

var path string = "G:\\Hacktiv8 - Scalable Web Services with Golang\\Assignment 3\\WeatherData\\"
var interval int = 15

type Weather struct {
	Wind  int
	Water int
}

func generateRandom(min, max int) int {
	return rand.Intn((max+1)-min) + min
}

func main() {
	for range time.Tick(time.Duration(interval) * time.Second) {
		checkDirectory()

		var weatherData []Weather
		readPreviousData(&weatherData)

		newWeatherData := createNewWeatherData()
		weatherData = append(weatherData, *newWeatherData)

		writeWeatherData(&weatherData)

		fmt.Printf("Success append data, current amount of data : %d\n", len(weatherData))
	}
}

func createNewWeatherData() *Weather {
	weather := &Weather{
		Wind:  generateRandom(1, 15),
		Water: generateRandom(1, 10),
	}
	return weather
}

func checkDirectory() {
	if _, err := os.Stat(path + "data.txt"); err != nil {
		os.Create(path + "data.txt")
	}
}

func readPreviousData(w *[]Weather) {
	content, err := os.ReadFile(path + "data.txt")
	if err != nil {
		panic(err)
	}
	json.Unmarshal(content, w)
}

func writeWeatherData(w *[]Weather) {
	jsonByte, _ := json.Marshal(*w)
	err := os.WriteFile(path+"data.txt", jsonByte, 0666)
	if err != nil {
		panic(err)
	}
}
