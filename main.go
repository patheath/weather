package main

import (
	"github.com/patheath/weather/internal/client"
)

const NumberHours = 10  // Number of hours to display

func main() {

	w := client.FetchWeather()
	w.Display()
}
