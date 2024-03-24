package model

import (
	"fmt"
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
		println(fmt.Sprintf("Hour %v: %v %v", v.Hour, v.Temp, v.Short))
	}
}
