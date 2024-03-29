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

func TestReadResponse(t *testing.T) {

	resp := &Response{
		Properties: Properties{
			Periods: []Period{
				{Number: 1, Temperature: 36, ShortForecast: "Sunny"},
				{Number: 2, Temperature: 42, ShortForecast: "Rainy and cold"},
				{Number: 9, Temperature: 103, ShortForecast: "Very hot and humid"},
			},
		},
	}

	// TODO - return my JSON instead of String in reponse
	// add checks to test

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		jsonResp, err := json.Marshal(resp)
		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}
		_, err = w.Write(jsonResp)
		if err != nil {
			log.Fatalf("Error sending JSON response. Err: %s", err)
		}
	}))
	defer ts.Close()

	res, err := http.Get(ts.URL)
	if err != nil {
		log.Fatal(err)
	}

	w := readResponse(res)

	check := w.Hourly[len(resp.Properties.Periods)-1]

	assert.Contains(t, check.Short, "Very hot and humid")
	assert.Equal(t, check.Temp, 103)
	assert.Equal(t, check.Hour, 9)

}
