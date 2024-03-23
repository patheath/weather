package main

import (
	"log"

	"github.com/patheath/weather/internal/repository"
)

func main() {

	// Request weather data
	w := repository.Weather()

	for _, v := range w {
		log.Println(v)
	}
}
