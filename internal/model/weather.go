package model

import (
	"fmt"
)

type Weather struct {
	Hourly []Forecast // "<number (hour from now)>: <tempature> <shortForecast>"
}

type Forecast struct {
	Hour  int
	Temp  int
	Short string
}

func (w Weather) Display() {
	for _, v := range w.Hourly {
		println(fmt.Sprintf("Hour %v: %v %v", v.Hour, v.Temp, v.Short))
	}
}
