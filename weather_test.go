package main

import (
	"bytes"
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFetchWeatherInfo(t *testing.T) {
	if os.Getenv("CI") == "true" {
		t.Skip()
	}
	weatherResponse := new(WeatherResponse)

	err := FetchWeatherInfo("臺北市", weatherResponse)
	if err != nil || weatherResponse.Success != "true" {
		t.Error("FetchWeatherInfo FAIL")
	}
}

func TestResolveWeatherResponse(t *testing.T) {
	weatherResponse := new(WeatherResponse)
	json.NewDecoder(bytes.NewReader([]byte(standardWeatherResponse))).Decode(weatherResponse)

	weather := ResolveWeatherResponse(weatherResponse)

	assert.Equal(t, weather.LocationName, "臺北市")
	assert.Equal(t, weather.State, "晴時多雲")
	assert.Equal(t, weather.RainProb, "20")
	assert.Equal(t, weather.StartTime, "2022-04-25 18:00:00")
	assert.Equal(t, weather.EndTime, "2022-04-26 06:00:00")
	assert.Equal(t, weather.Confort, "舒適至悶熱")
	assert.Equal(t, weather.MaxTemp, "30")
	assert.Equal(t, weather.MinTemp, "24")
}

const standardWeatherResponse string = `{
	"success": "true",
	"result": {
	  "resource_id": "F-C0032-001",
	  "fields": [
		{
		  "id": "datasetDescription",
		  "type": "String"
		},
		{
		  "id": "locationName",
		  "type": "String"
		},
		{
		  "id": "parameterName",
		  "type": "String"
		},
		{
		  "id": "parameterValue",
		  "type": "String"
		},
		{
		  "id": "parameterUnit",
		  "type": "String"
		},
		{
		  "id": "startTime",
		  "type": "Timestamp"
		},
		{
		  "id": "endTime",
		  "type": "Timestamp"
		}
	  ]
	},
	"records": {
	  "datasetDescription": "三十六小時天氣預報",
	  "location": [
		{
		  "locationName": "臺北市",
		  "weatherElement": [
			{
			  "elementName": "Wx",
			  "time": [
				{
				  "startTime": "2022-04-25 18:00:00",
				  "endTime": "2022-04-26 06:00:00",
				  "parameter": {
					"parameterName": "晴時多雲",
					"parameterValue": "2"
				  }
				},
				{
				  "startTime": "2022-04-26 06:00:00",
				  "endTime": "2022-04-26 18:00:00",
				  "parameter": {
					"parameterName": "多雲午後短暫雷陣雨",
					"parameterValue": "22"
				  }
				},
				{
				  "startTime": "2022-04-26 18:00:00",
				  "endTime": "2022-04-27 06:00:00",
				  "parameter": {
					"parameterName": "晴時多雲",
					"parameterValue": "2"
				  }
				}
			  ]
			},
			{
			  "elementName": "PoP",
			  "time": [
				{
				  "startTime": "2022-04-25 18:00:00",
				  "endTime": "2022-04-26 06:00:00",
				  "parameter": {
					"parameterName": "20",
					"parameterUnit": "百分比"
				  }
				},
				{
				  "startTime": "2022-04-26 06:00:00",
				  "endTime": "2022-04-26 18:00:00",
				  "parameter": {
					"parameterName": "80",
					"parameterUnit": "百分比"
				  }
				},
				{
				  "startTime": "2022-04-26 18:00:00",
				  "endTime": "2022-04-27 06:00:00",
				  "parameter": {
					"parameterName": "20",
					"parameterUnit": "百分比"
				  }
				}
			  ]
			},
			{
			  "elementName": "MinT",
			  "time": [
				{
				  "startTime": "2022-04-25 18:00:00",
				  "endTime": "2022-04-26 06:00:00",
				  "parameter": {
					"parameterName": "24",
					"parameterUnit": "C"
				  }
				},
				{
				  "startTime": "2022-04-26 06:00:00",
				  "endTime": "2022-04-26 18:00:00",
				  "parameter": {
					"parameterName": "24",
					"parameterUnit": "C"
				  }
				},
				{
				  "startTime": "2022-04-26 18:00:00",
				  "endTime": "2022-04-27 06:00:00",
				  "parameter": {
					"parameterName": "24",
					"parameterUnit": "C"
				  }
				}
			  ]
			},
			{
			  "elementName": "CI",
			  "time": [
				{
				  "startTime": "2022-04-25 18:00:00",
				  "endTime": "2022-04-26 06:00:00",
				  "parameter": {
					"parameterName": "舒適至悶熱"
				  }
				},
				{
				  "startTime": "2022-04-26 06:00:00",
				  "endTime": "2022-04-26 18:00:00",
				  "parameter": {
					"parameterName": "舒適至悶熱"
				  }
				},
				{
				  "startTime": "2022-04-26 18:00:00",
				  "endTime": "2022-04-27 06:00:00",
				  "parameter": {
					"parameterName": "舒適至悶熱"
				  }
				}
			  ]
			},
			{
			  "elementName": "MaxT",
			  "time": [
				{
				  "startTime": "2022-04-25 18:00:00",
				  "endTime": "2022-04-26 06:00:00",
				  "parameter": {
					"parameterName": "30",
					"parameterUnit": "C"
				  }
				},
				{
				  "startTime": "2022-04-26 06:00:00",
				  "endTime": "2022-04-26 18:00:00",
				  "parameter": {
					"parameterName": "34",
					"parameterUnit": "C"
				  }
				},
				{
				  "startTime": "2022-04-26 18:00:00",
				  "endTime": "2022-04-27 06:00:00",
				  "parameter": {
					"parameterName": "30",
					"parameterUnit": "C"
				  }
				}
			  ]
			}
		  ]
		}
	  ]
	}
  }`
