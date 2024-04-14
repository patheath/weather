package weatherapi

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/patheath/weather/internal/model"
	"github.com/stretchr/testify/assert"
)

type Condition struct {
	Text string `json:"text"`
}

type Hour struct {
	TempF     float64   `json:"temp_f"`
	Condition Condition `json:"condition"`
}

type Forecastday struct {
	Hour []Hour `json:"hour"`
}

type Forecast struct {
	Forecastday []Forecastday `json:"forecastday"`
}

type Response struct {
	Forecast Forecast `json:"forecast"`
}

func initHours(hours int) []Hour {
	s := 30.0 // starting temp
	c := Condition{Text: "Cold and cloudy"}
	h := make([]Hour, hours)
	for i := 0; i < hours; i++ {
		h[i] = Hour{TempF: s + float64(i), Condition: c}
	}
	return h
}

func TestReadResponseBadStatusCode(t *testing.T) {

	ts := httptest.NewServer(http.NotFoundHandler())
	defer ts.Close()

	wa := WeatherApi{Url: ts.URL}

	_, err := wa.GetWeather()
	assert.Contains(t, err.Error(), http.StatusText(http.StatusNotFound))
}

func TestReadResponseIncompleteForecast(t *testing.T) {

	h := initHours(3)

	resp := &Response{
		Forecast: Forecast{
			Forecastday: []Forecastday{{Hour: h}},
		},
	}

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatal("Error converting response to JSON")
	}

	wa := WeatherApi{Url: "example.com"}
	_, err = wa.ReadResponse(jsonResp)

	assert.Contains(t, err.Error(), "less than 24")
}

func TestReadResponseOK(t *testing.T) {

	h := initHours(24)
	hour := time.Now().Hour()

	resp := &Response{
		Forecast: Forecast{
			Forecastday: []Forecastday{{Hour: h}, {Hour: h}},
		},
	}

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatal("Error converting response to JSON")
	}

	wa := WeatherApi{Url: "example.com"}
	w, _ := wa.ReadResponse(jsonResp)

	if hour == 23 {
		assert.Len(t, w.Hourly, 0)
		return
	} else {
		i := min(24-hour, model.HOURS)

		check := w.Hourly[0]
		assert.Contains(t, check.Short, "Cold and cloudy")
		assert.Equal(t, 30+hour, int(check.Temp))
		assert.Len(t, w.Hourly, i)
	}
}

func TestReadResponseWrongContent(t *testing.T) {

	res := []byte("Bad response")
	wa := WeatherApi{Url: "example.com"}
	_, err := wa.ReadResponse(res)
	assert.Contains(t, err.Error(), "Error unmarshalling response")
}

func TestReadResponseEmpty(t *testing.T) {

	resp := &Response{
		Forecast: Forecast{},
	}

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatal("Error converting response to JSON")
	}

	wa := WeatherApi{Url: "example.com"}
	_, err = wa.ReadResponse(jsonResp)
	assert.Equal(t, "No results from weatherapi.com provider", err.Error())

}

func TestDisplayName(t *testing.T) {
	wa := WeatherApi{Url: "example.com"}
	assert.Contains(t, wa.DisplayName(), "weatherapi.com")

}

func TestWeatherApiIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping weather api integration test")
	}

	wa := WeatherApi{Url: Url}
	resp, err := wa.GetWeather()
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	w, err := wa.ReadResponse(resp)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	assert.Equal(t, model.HOURS, len(w.Hourly))
}
