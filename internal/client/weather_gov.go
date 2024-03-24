package client

import (
	"io"
	"log"
	"net/http"

	"github.com/patheath/weather/internal/model"
)

func FetchWeather() model.Weather {
	resp := getWeather()
	w := readResponse(resp)
	return w
}

func readResponse(resp *http.Response) model.Weather {

	defer resp.Body.Close()

	//Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	sb := string(body)
	//log.Printf(sb)

	h := []model.Forecast{}
	f := model.Forecast{
		Hour: 1,
		Temp: 30,
		Short: sb[:10],
	}
	h = append(h, f)

	return model.Weather{
		Hourly: h,
	}
}

func getWeather() *http.Response {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.weather.gov/gridpoints/TOP/64,85/forecast/hourly", nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("User-Agent", "(weather-cli, patheath@gmail.com)")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	return resp
}
