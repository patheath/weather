package client

import "github.com/patheath/weather/internal/model"


type Api interface {
	GetWeather() ([]byte, error)
	ReadResponse([]byte) (*model.Weather, error)
	DisplayName() string
}

func FetchWeather(wa Api) (*model.Weather, error) {
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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}