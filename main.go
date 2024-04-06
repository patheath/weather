package main

import (
	"fmt"

	"github.com/patheath/weather/internal/client"
)

const NumberHours = 10  // Number of hours to display

func main() {

	wg := client.WeatherGov{Url: client.Url}
	w, err := client.FetchWeather(wg)
	if err != nil {
		fmt.Print("Error fetching weather: ", err)
		return
	}
	println(wg.DisplayName())
	w.Display()
}
