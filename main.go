package main

import (
	"github.com/patheath/weather/internal/client"
	"github.com/patheath/weather/internal/client/weatherapi"
	"github.com/patheath/weather/internal/client/weathergov"
	"github.com/patheath/weather/internal/model"
)

const NumberHours = 10 // Number of hours to display

type result struct {
	w    *model.Weather
	name string
	err  error
}

func main() {

	clients := []client.Api{
		weatherapi.WeatherApi{Url: weatherapi.Url},
		weathergov.WeatherGov{Url: weathergov.Url},
	}

	ch := make(chan result, len(clients))

	for _, c := range clients {
		go func() {
			name := c.DisplayName()
			w, err := client.FetchWeather(c)
			ch <- result{
				w:    w,
				name: name,
				err:  err,
			}
		}()
	}

	for range clients {
		r := <-ch
		println(r.name)
		if r.err != nil {
			println(r.err.Error())
		} else {
			r.w.Display()
		}
	}

}
