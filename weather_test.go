package main

import "testing"

func TestFetchWeatherInfo(t *testing.T) {
	weatherResponse := new(WeatherResponse)

	err := FetchWeatherInfo("臺北市", weatherResponse)
	if err != nil {
		t.Error("FetchWeatherInfo FAIL")
	}

	weather := ResolveWeatherResponse(weatherResponse)

	GetWeatherMsg(weather)

}
