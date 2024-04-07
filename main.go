package main

import (
	"fmt"

	"github.com/patheath/weather/internal/client"
	"github.com/patheath/weather/internal/client/weatherapi"
	"github.com/patheath/weather/internal/client/weathergov"
)

const NumberHours = 10 // Number of hours to display

func main() {

	wa := weatherapi.WeatherApi{Url: weatherapi.Url}
	w, err := client.FetchWeather(wa)
	if err != nil {
		fmt.Print("Error fetching weather: ", err)
		return
	}
	println(wa.DisplayName())
	w.Display()

	wg := weathergov.WeatherGov{Url: weathergov.Url}
	w, err = client.FetchWeather(wg)
	if err != nil {
		fmt.Print("Error fetching weather: ", err)
		return
	}
	println(wg.DisplayName())
	w.Display()
}
