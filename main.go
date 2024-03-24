package main

import (
	"github.com/patheath/weather/internal/client"
)

func main() {

	w := client.FetchWeather()
	w.Display()
}
