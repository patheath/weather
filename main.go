package main

import (
	"fmt"

	"github.com/patheath/weather/internal/client"
)

const NumberHours = 10  // Number of hours to display

func main() {

	w, err := client.FetchWeather()
	if err != nil {
		fmt.Print("Error fetching weather: ", err)
		return
	}
	w.Display()
}
