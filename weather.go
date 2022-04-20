package main

import (
	"bytes"
	"log"
	"net/http"
	"os"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func FetchWeatherInfo(locationName string) (string, error) {
	url := "https://opendata.cwb.gov.tw/api/v1/rest/datastore/F-C0032-001?Authorization=" + os.Getenv("WeatherAuthorization") + "&locationName=" + locationName

	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != 200 {
		log.Print("Error fetching weather info: ", err)
		return "", err
	}

	buf := new(bytes.Buffer)
	if _, err := buf.ReadFrom(resp.Body); err != nil {
		log.Print("Error read from req.body to string: ", err)
		return "", err
	}

	return buf.String(), nil
}

func GetWeatherInfo(locationName string) *linebot.TextMessage {
	data, err := FetchWeatherInfo("臺北市")

	if err != nil {
		return linebot.NewTextMessage("抱歉，得不到此地區天氣資料><")
	} else {
		return linebot.NewTextMessage(data)
	}
}
