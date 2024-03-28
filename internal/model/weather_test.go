package model

import (
	"strconv"
	"testing"
	"time"

	"github.com/patheath/weather/internal/utils"
	"github.com/stretchr/testify/assert"
)

var WeatherData = []Weather{}

func TestDisplay(t *testing.T) {

	d := Weather{
		[]Forecast{
			{1, 35, "Sunny"},
			{2, 46, "Getting hotter"},
			{3, 98, "Hot and humid"},
		},
	}

	output := utils.CaptureStdout(d.Display)

	for _, v := range d.Hourly {
		// ignore hour as its added to current time and converted
		// temp and short should be enough
		assert.Contains(t, output, strconv.Itoa(v.Temp))
		assert.Contains(t, output, v.Short)
	}
}

func TestDisplayTime(t *testing.T) {

	now = func() time.Time { return time.Date(2024, time.March, 10, 10, 0, 0, 0, time.UTC) }
	hour := 5
	result := "3pm"  // 5 hours past 10am
	assert.Contains(t, displayTime(hour), result)
}
