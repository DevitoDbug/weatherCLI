package main

import (
	"fmt"
	"io"
	"net/http"
)

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
