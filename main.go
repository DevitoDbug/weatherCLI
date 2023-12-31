package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Weather struct {
	Name    string `json:"name"`
	Weather []struct {
		Description string `json:"description"`
	} `json:"weather"`
	Main struct {
		Temp     float64 `json:"temp"`
		Humidity int     `json:"humidity"`
	} `json:"main"`
	Sys struct {
		Country string `json:"country"`
	} `json:"sys"`
}

func main() {
	city := "Nairobi"

	if len(os.Args) >= 2 {
		city = os.Args[1]
	}

	res, err := http.Get("https://api.openweathermap.org/data/2.5/weather?q=" + city + "&appid=73cbc8ffd789d41ed5535c3fdb8ca562")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		panic("Weather api not available ")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic("Could not read body of response")
	}

	var weather Weather
	err = json.Unmarshal(body, &weather)
	if err != nil {
		panic("Could not unmarshal data")
	}

	country, name, description, temp, humidity := weather.Sys.Country, weather.Name, weather.Weather[0].Description, weather.Main.Temp, weather.Main.Humidity
	fmt.Printf("Country: %v\nCity: %v\nWeather:%v, Temp:%.0f , Humidity:%v\n", country, name, description, temp, humidity)
}
