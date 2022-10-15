package main

import (
	"encoding/json"
	"net/http"
	"os"
	"text/template"
	"time"
)

var path string = "G:\\Hacktiv8 - Scalable Web Services with Golang\\Assignment 3\\WeatherData\\"

type WeatherData struct {
	Wind  int
	Water int
}

type ClientData struct {
	WeatherData
	Date           string
	Time           string
	WindCondition  string
	WaterCondition string
}

var client ClientData

func main() {
	go func() {
		for range time.Tick(time.Second) {
			checkDirectory()

			data := []WeatherData{}
			getData(&data)
			if len(data) == 0 {
				panic("data is null")
			}

			client.Date = time.Now().Format("02/01/2006")
			client.Time = time.Now().Format("15:04")
			client.Wind = data[len(data)-1].Wind
			client.Water = data[len(data)-1].Water
			client.WindCondition = parsingWindCondition(client.Wind)
			client.WaterCondition = parsingWindCondition(client.Water)
		}
	}()

	http.HandleFunc("/", homePageHandler)
	http.ListenAndServe(":8080", nil)
}

func checkDirectory() {
	if _, err := os.Stat(path + "data.txt"); err != nil {
		panic("Data not found !")
	}
}

func getData(data *[]WeatherData) {
	jsonByte, err := os.ReadFile(path + "data.txt")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(jsonByte, data)
	if err != nil {
		panic(err)
	}
}

func homePageHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("G:\\Hacktiv8 - Scalable Web Services with Golang\\Assignment 3\\HTML Template\\index.html")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(w, client)
}

func parsingWaterCondition(waterValue int) string {
	switch {
	case waterValue > 8:
		return "Bahaya"
	case waterValue <= 8 && waterValue >= 6:
		return "Siaga"
	default:
		return "Aman"
	}
}

func parsingWindCondition(windValue int) string {
	switch {
	case windValue > 15:
		return "Bahaya"
	case windValue <= 15 && windValue >= 7:
		return "Siaga"
	default:
		return "Aman"
	}
}
