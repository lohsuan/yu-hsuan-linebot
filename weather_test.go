package main

import "testing"

func TestFetchWeatherInfo(t *testing.T) {
	if info, _ := FetchWeatherInfo("臺北市"); info != "" {
		t.Log("FetchWeatherInfo PASS")
	} else {
		t.Error("FetchWeatherInfo FAIL")
	}
}
