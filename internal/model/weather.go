package model

import (
	"fmt"
	"time"
)

type Weather struct {
	Hourly []Forecast // "<number (hour from now)>: <tempature> <shortForecast>"
}

type Forecast struct {
	Hour  int    `json:"number"`
	Temp  int    `json:"temperature"`
	Short string `json:"shortForecast"`
}

var now = time.Now  // allows for testing

func (w Weather) Display() {
	for _, v := range w.Hourly {
		fmt.Printf("%v: %v - %v\n", displayTime(v.Hour), v.Temp, v.Short)
	}
}

func displayTime(hour int) string {
	toAdd := time.Duration(hour) * time.Hour
	currentTime := now().Add(toAdd)
	return fmt.Sprintf("The hour is %v", currentTime.Format("2006-1-2: 3pm"))
}
