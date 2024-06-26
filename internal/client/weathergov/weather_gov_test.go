package weathergov

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Period struct {
	Number        int    `json:"number"`
	Temperature   int    `json:"temperature"`
	ShortForecast string `json:"shortForecast"`
}

type Properties struct {
	Periods []Period `json:"periods"`
}

type Response struct {
	Properties Properties `json:"properties"`
}

func TestReadResponseBadStatusCode(t *testing.T) {

	ts := httptest.NewServer(http.NotFoundHandler())
	defer ts.Close()

	wg := WeatherGov{Url: ts.URL}

	_, err := wg.GetWeather()
	assert.Contains(t, err.Error(), http.StatusText(http.StatusNotFound))
}

func TestReadResponseOK(t *testing.T) {

	resp := &Response{
		Properties: Properties{
			Periods: []Period{
				{Number: 1, Temperature: 36, ShortForecast: "Sunny"},
				{Number: 2, Temperature: 42, ShortForecast: "Rainy and cold"},
				{Number: 9, Temperature: 103, ShortForecast: "Very hot and humid"},
			},
		},
	}

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatal("Error converting response to JSON")
	}

	wg := WeatherGov{Url: "example.com"}
	w, _ := wg.ReadResponse(jsonResp)

	check := w.Hourly[len(resp.Properties.Periods)-1]

	assert.Contains(t, check.Short, "Very hot and humid")
	assert.Equal(t, check.Temp, 103)
	assert.Equal(t, check.Hour, 9)

}

func TestReadResponseWrongContent(t *testing.T) {

	res := []byte("Bad response")
	wg := WeatherGov{Url: "example.com"}
	_, err := wg.ReadResponse(res)
	assert.Contains(t, err.Error(), "Error unmarshalling response")
}

func TestReadResponseEmpty(t *testing.T) {

	resp := &Response{
		Properties: Properties{},
	}

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatal("Error converting response to JSON")
	}

	wg := WeatherGov{Url: "example.com"}
	_, err = wg.ReadResponse(jsonResp)
	assert.Equal(t, err.Error(), "No results from weather.gov provider")

}

func TestDisplayName(t *testing.T) {
	wg := WeatherGov{Url: "example.com"}
	assert.Contains(t, wg.DisplayName(), "weather.gov")

}
