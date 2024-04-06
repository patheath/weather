package client

import "github.com/patheath/weather/internal/model"


type WeatherApi interface {
	GetWeather() ([]byte, error)
	ReadResponse([]byte) (*model.Weather, error)
	DisplayName() string
}

func FetchWeather(wa WeatherApi) (*model.Weather, error) {
	resp, err := wa.GetWeather()
	if err != nil {
		return nil, err
	}
	w, err := wa.ReadResponse(resp)
	if err != nil {
		return nil, err
	}
	return w, nil
}
