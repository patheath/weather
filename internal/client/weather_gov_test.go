package client

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

	_, err := getWeather(ts.URL)
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

	w, _ := readResponse(jsonResp)

	check := w.Hourly[len(resp.Properties.Periods)-1]

	assert.Contains(t, check.Short, "Very hot and humid")
	assert.Equal(t, check.Temp, 103)
	assert.Equal(t, check.Hour, 9)

}

func TestReadResponseWrongContent(t *testing.T) {

	res := []byte("Bad response")
	_, err := readResponse(res)
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

	_, err = readResponse(jsonResp)
	assert.Equal(t, err.Error(), "No results from weather.gov provider")

}
