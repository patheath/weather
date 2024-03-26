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

func (w Weather) Display() {
	for _, v := range w.Hourly {
		println(fmt.Sprintf("Hour %v: %v - %v", displayTime(v.Hour), v.Temp, v.Short))
	}
}

func displayTime(hour int) string {
	toAdd := time.Duration(hour) * time.Hour
	currentTime := time.Now().Add(toAdd)
	return fmt.Sprintf("The hour is %v", currentTime.Format("2006-1-2: 3pm"))
}
