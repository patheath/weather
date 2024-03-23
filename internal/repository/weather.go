package repository

import (
	"io"
	"log"
	"net/http"
)

type weather struct {
	hourly []string
}

func Weather() []string {

	resp := getWeather()
	w := readResponse(resp)
	return w.hourly
}

func readResponse(resp *http.Response) weather {

	defer resp.Body.Close()

	//Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	sb := string(body)
	log.Printf(sb)

	h := []string{}
	h = append(h, sb)

	return weather{
		hourly: h,
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
