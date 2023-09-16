package main

import (
	"fmt"
	"io"
	"net/http"
)

type Weather struct {
	Name    string `json:"name"`
	Weather []struct {
		Description string `json:"description"`
	} `json:"weather"`
	Main struct {
		Temp     int `json:"temp"`
		Pressure int `json:"pressure"`
		Humidity int `json:"humidity"`
	} `json:"main"`
	Sys struct {
		Country string `json:"country"`
	} `json:"sys"`
}

func main() {
	res, err := http.Get("https://api.openweathermap.org/data/2.5/weather?q=Nairobi&appid=73cbc8ffd789d41ed5535c3fdb8ca562")
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
	fmt.Println(string(body))
}
