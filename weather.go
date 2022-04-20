package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

type Weather struct {
	LocationName string `json:"locationName"`
	State        string `json:"state"`
	RainProb     string `json:"rainProb"`
	StartTime    string `json:"startTime"`
	EndTime      string `json:"endTime"`
	Confort      string `json:"confort"`
	MaxTemp      string `json:"maxTemp"`
	MinTemp      string `json:"minTemp"`
}

type WeatherResponse struct {
	Success string `json:"success"`
	Result  struct {
		ResourceID string `json:"resource_id"`
		Fields     []struct {
			ID   string `json:"id"`
			Type string `json:"type"`
		} `json:"fields"`
	} `json:"result"`
	Records struct {
		DatasetDescription string `json:"datasetDescription"`
		Location           []struct {
			LocationName   string `json:"locationName"`
			WeatherElement []struct {
				ElementName string `json:"elementName"`
				Time        []struct {
					StartTime string `json:"startTime"`
					EndTime   string `json:"endTime"`
					Parameter struct {
						ParameterName  string `json:"parameterName"`
						ParameterValue string `json:"parameterValue"`
					} `json:"parameter"`
				} `json:"time"`
			} `json:"weatherElement"`
		} `json:"location"`
	} `json:"records"`
}

func FetchWeatherInfo(locationName string, target interface{}) error {
	url := "https://opendata.cwb.gov.tw/api/v1/rest/datastore/F-C0032-001?Authorization=" + os.Getenv("WeatherAuthorization") + "&locationName=" + locationName

	response, err := http.Get(url)
	if err != nil || response.StatusCode != 200 {
		return err
	}
	defer response.Body.Close()

	return json.NewDecoder(response.Body).Decode(target)
}

func ResolveWeatherResponse(weatherResponse *WeatherResponse) Weather {
	var weather Weather
	weather.LocationName = weatherResponse.Records.Location[0].LocationName
	weatherElements := weatherResponse.Records.Location[0].WeatherElement
	weather.StartTime = weatherElements[0].Time[0].StartTime
	weather.EndTime = weatherElements[0].Time[0].EndTime
	weather.State = weatherElements[0].Time[0].Parameter.ParameterName
	weather.RainProb = weatherElements[1].Time[0].Parameter.ParameterName
	weather.MinTemp = weatherElements[2].Time[0].Parameter.ParameterName
	weather.Confort = weatherElements[3].Time[0].Parameter.ParameterName
	weather.MaxTemp = weatherElements[4].Time[0].Parameter.ParameterName
	return weather
}

func GetWeatherMesg(weather Weather) *linebot.TextMessage {
	message := fmt.Sprintln("【天氣小助理】")
	message += fmt.Sprintln("今天" + weather.LocationName + "的天氣: " + weather.State)
	message += fmt.Sprintln("溫度: " + weather.MinTemp + "°C - " + weather.MaxTemp + "°C")
	message += fmt.Sprintln("降雨機率: " + weather.RainProb + "%")
	message += fmt.Sprintln("舒適度: " + weather.Confort)
	message += fmt.Sprint("時間: " + weather.StartTime[5:16] + "~" + weather.EndTime[5:16])

	if i, _ := strconv.Atoi(weather.RainProb); i > 70 {
		message += "\n提醒您，降雨機率高，出門記得帶把傘唷！"
	} else if i, _ := strconv.Atoi(weather.MaxTemp); i > 27 {
		message += "\n提醒您，今天有點熱，外出要小心中暑唷！"
	} else if i, _ := strconv.Atoi(weather.MinTemp); i < 15 {
		message += "\n提醒您，今天天氣偏涼，記得多穿一件外套唷！"
	}
	return linebot.NewTextMessage(message)
}

func GetWeatherSticker(weather Weather) *linebot.StickerMessage {
	rainProb, _ := strconv.Atoi(weather.RainProb)

	if rainProb > 70 {
		return linebot.NewStickerMessage("789", "10893")
	} else if rainProb > 30 {
		return linebot.NewStickerMessage("789", "10872")
	} else {
		return linebot.NewStickerMessage("789", "10871")
	}
}

func GetWeatherInfo(locationName string) (*linebot.TextMessage, *linebot.StickerMessage) {
	weatherResponse := new(WeatherResponse)

	err := FetchWeatherInfo(locationName, weatherResponse)
	if err != nil {
		return linebot.NewTextMessage("抱歉，得不到此地區天氣資料><"), linebot.NewStickerMessage("6136", "10551376")
	}

	weather := ResolveWeatherResponse(weatherResponse)

	return GetWeatherMesg(weather), GetWeatherSticker(weather)
}
