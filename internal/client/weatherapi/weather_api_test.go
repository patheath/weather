package weatherapi

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

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

func TestReadResponseBadStatusCode(t *testing.T) {

	ts := httptest.NewServer(http.NotFoundHandler())
	defer ts.Close()

	wa := WeatherApi{Url: ts.URL}

	_, err := wa.GetWeather()
	assert.Contains(t, err.Error(), http.StatusText(http.StatusNotFound))
}

func TestReadResponseOK(t *testing.T) {

	resp := &Response{
		Forecast: Forecast{
			Forecastday: []Forecastday{
				{
					Hour: []Hour{
						{TempF: 35.7, Condition: Condition{
							Text: "Sunny",
						}},
						{TempF: 42, Condition: Condition{
							Text: "Rainy and cold",
						}},
						{TempF: 103.1, Condition: Condition{
							Text: "Very hot and humid",
						}},
					},
				},
			},
		},
	}

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatal("Error converting response to JSON")
	}

	wa := WeatherApi{Url: "example.com"}
	w, _ := wa.ReadResponse(jsonResp)

	check := w.Hourly[len(resp.Forecast.Forecastday[0].Hour)-1]

	assert.Contains(t, check.Short, "Very hot and humid")
	assert.Equal(t, check.Temp, 103)
	assert.Equal(t, check.Hour, 2) //Fixme not hardcoded.

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
	assert.Contains(t, wa.DisplayName(), "weatherapi.com", )

}
