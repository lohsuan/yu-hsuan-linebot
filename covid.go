package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

const covid19 = "covid19"
const covid19Info = "關注疫情動態"
const errorFetchingDataMesg = "資料發生錯誤，請稍後再試"
const globalCovidUrl = "https://covid19dashboard.cdc.gov.tw/dash2"
const taiwanCovidUrl = "https://covid19dashboard.cdc.gov.tw/dash3"
const covidInfoFlexAlt = "疫情概況一覽"

type GlabalCovid struct {
	Zero struct {
		Cases     string `json:"cases"`
		Deaths    string `json:"deaths"`
		Cfr       string `json:"cfr"`
		Countries int    `json:"countries"`
	} `json:"0"`
}

type TaiwanCovid struct {
	Zero struct {
		Case       string `json:"確診"`
		Death      int    `json:"死亡"`
		Inform     string `json:"送驗"`
		Except     string `json:"排除"`
		LastCase   int    `json:"昨日確診"`
		LastExcept string `json:"昨日排除"`
		LastInform string `json:"昨日送驗"`
	} `json:"0"`
}

func FetchGlabalCovidInfo(target interface{}) error {
	url := globalCovidUrl

	response, err := http.Get(url)
	if err != nil || response.StatusCode != 200 {
		return err
	}
	defer response.Body.Close()

	return json.NewDecoder(response.Body).Decode(target)
}

func FetchTaiwanCovidInfo(target interface{}) error {
	url := taiwanCovidUrl

	response, err := http.Get(url)
	if err != nil || response.StatusCode != 200 {
		return err
	}
	defer response.Body.Close()

	return json.NewDecoder(response.Body).Decode(target)
}

func GetCovidInfo() linebot.SendingMessage {
	flexMessage := fmt.Sprint(constFlexMessage)

	globalCovid := new(GlabalCovid)
	err := FetchGlabalCovidInfo(globalCovid)
	if err != nil {
		return linebot.NewTextMessage(errorFetchingDataMesg)
	}
	flexMessage = PrepareFlexMesgForGlobalInfo(flexMessage, globalCovid)

	taiwanCovid := new(TaiwanCovid)
	err = FetchTaiwanCovidInfo(taiwanCovid)
	if err != nil {
		return linebot.NewTextMessage(errorFetchingDataMesg)
	}
	flexMessage = PrepareFlexMesgForTaiwanInfo(flexMessage, taiwanCovid)

	container, _ := linebot.UnmarshalFlexMessageJSON([]byte(flexMessage))
	return linebot.NewFlexMessage(covidInfoFlexAlt, container)
}

func PrepareFlexMesgForGlobalInfo(flexMessage string, globalCovid *GlabalCovid) string {
	flexMessage = strings.Replace(flexMessage, "CASES", globalCovid.Zero.Cases, 1)
	flexMessage = strings.Replace(flexMessage, "DEATHS", globalCovid.Zero.Deaths, 1)
	flexMessage = strings.Replace(flexMessage, "CFR", globalCovid.Zero.Cfr, 1)
	flexMessage = strings.Replace(flexMessage, "COUNTRIES", strconv.Itoa(globalCovid.Zero.Countries), 1)
	return flexMessage
}

func PrepareFlexMesgForTaiwanInfo(flexMessage string, taiwanCovid *TaiwanCovid) string {
	flexMessage = strings.Replace(flexMessage, "LASTCASE", strconv.Itoa(taiwanCovid.Zero.LastCase), 1)
	flexMessage = strings.Replace(flexMessage, "LASTINFORM", taiwanCovid.Zero.LastInform, 1)
	flexMessage = strings.Replace(flexMessage, "LASTEXCEPT", taiwanCovid.Zero.LastExcept, 1)
	flexMessage = strings.Replace(flexMessage, "CASE", taiwanCovid.Zero.Case, 1)
	flexMessage = strings.Replace(flexMessage, "INFORM", taiwanCovid.Zero.Inform, 1)
	flexMessage = strings.Replace(flexMessage, "DEATH", strconv.Itoa(taiwanCovid.Zero.Death), 1)
	return flexMessage
}

const constFlexMessage string = `{
	"type": "carousel",
	"contents": [
	  {
		"type": "bubble",
		"size": "micro",
		"body": {
		  "type": "box",
		  "layout": "vertical",
		  "contents": [
			{
			  "type": "text",
			  "text": "昨日新增",
			  "weight": "bold",
			  "size": "xl",
			  "margin": "md",
			  "color": "#00BB00"
			},
			{
			  "type": "separator",
			  "margin": "xs"
			},
			{
			  "type": "box",
			  "layout": "vertical",
			  "margin": "xxl",
			  "spacing": "sm",
			  "contents": [
				{
				  "type": "box",
				  "layout": "horizontal",
				  "contents": [
					{
					  "type": "text",
					  "text": "確診",
					  "size": "sm",
					  "color": "#555555",
					  "flex": 0
					},
					{
					  "type": "text",
					  "text": "LASTCASE",
					  "size": "sm",
					  "color": "#111111",
					  "align": "end"
					}
				  ]
				},
				{
				  "type": "box",
				  "layout": "horizontal",
				  "contents": [
					{
					  "type": "text",
					  "text": "通報",
					  "size": "sm",
					  "color": "#555555",
					  "flex": 0
					},
					{
					  "type": "text",
					  "size": "sm",
					  "color": "#111111",
					  "align": "end",
					  "text": "LASTINFORM"
					}
				  ]
				},
				{
				  "type": "box",
				  "layout": "horizontal",
				  "contents": [
					{
					  "type": "text",
					  "text": "排除",
					  "size": "sm",
					  "color": "#555555"
					},
					{
					  "type": "text",
					  "text": "LASTEXCEPT",
					  "size": "sm",
					  "color": "#111111",
					  "align": "end"
					}
				  ]
				}
			  ]
			},
			{
			  "type": "separator",
			  "margin": "lg"
			},
			{
			  "type": "box",
			  "layout": "horizontal",
			  "margin": "md",
			  "contents": [
				{
				  "type": "text",
				  "text": "Taiwan",
				  "size": "xs",
				  "color": "#aaaaaa",
				  "flex": 0
				}
			  ]
			}
		  ]
		},
		"action": {
		  "type": "uri",
		  "label": "action",
		  "uri": "https://www.cdc.gov.tw/"
		},
		"styles": {
		  "footer": {
			"separator": true
		  }
		}
	  },
	  {
		"type": "bubble",
		"size": "micro",
		"body": {
		  "type": "box",
		  "layout": "vertical",
		  "contents": [
			{
			  "type": "text",
			  "text": "國內統計",
			  "weight": "bold",
			  "size": "xl",
			  "margin": "md",
			  "color": "#00AEAE"
			},
			{
			  "type": "separator",
			  "margin": "xs"
			},
			{
			  "type": "box",
			  "layout": "vertical",
			  "margin": "xxl",
			  "spacing": "sm",
			  "contents": [
				{
				  "type": "box",
				  "layout": "horizontal",
				  "contents": [
					{
					  "type": "text",
					  "text": "確診",
					  "size": "sm",
					  "color": "#555555",
					  "flex": 0
					},
					{
					  "type": "text",
					  "text": "CASE",
					  "size": "sm",
					  "color": "#111111",
					  "align": "end"
					}
				  ]
				},
				{
				  "type": "box",
				  "layout": "horizontal",
				  "contents": [
					{
					  "type": "text",
					  "text": "通報",
					  "size": "sm",
					  "color": "#555555",
					  "flex": 0
					},
					{
					  "type": "text",
					  "size": "sm",
					  "color": "#111111",
					  "align": "end",
					  "text": "INFORM"
					}
				  ]
				},
				{
				  "type": "box",
				  "layout": "horizontal",
				  "contents": [
					{
					  "type": "text",
					  "text": "死亡",
					  "size": "sm",
					  "color": "#555555",
					  "flex": 0
					},
					{
					  "type": "text",
					  "text": "DEATH",
					  "size": "sm",
					  "color": "#111111",
					  "align": "end"
					}
				  ]
				}
			  ]
			},
			{
			  "type": "separator",
			  "margin": "lg"
			},
			{
			  "type": "box",
			  "layout": "horizontal",
			  "margin": "md",
			  "contents": [
				{
				  "type": "text",
				  "text": "Taiwan",
				  "size": "xs",
				  "color": "#aaaaaa",
				  "flex": 0
				}
			  ]
			}
		  ]
		},
		"action": {
		  "type": "uri",
		  "label": "action",
		  "uri": "https://www.cdc.gov.tw/"
		},
		"styles": {
		  "footer": {
			"separator": true
		  }
		}
	  },
	  {
		"type": "bubble",
		"size": "micro",
		"body": {
		  "type": "box",
		  "layout": "vertical",
		  "contents": [
			{
			  "type": "text",
			  "text": "全球病例",
			  "weight": "bold",
			  "size": "xl",
			  "margin": "md",
			  "color": "#0066CC"
			},
			{
			  "type": "separator",
			  "margin": "xs"
			},
			{
			  "type": "box",
			  "layout": "vertical",
			  "margin": "xxl",
			  "spacing": "sm",
			  "contents": [
				{
				  "type": "box",
				  "layout": "horizontal",
				  "contents": [
					{
					  "type": "text",
					  "text": "確診",
					  "size": "sm",
					  "color": "#555555",
					  "flex": 0
					},
					{
					  "type": "text",
					  "size": "sm",
					  "color": "#111111",
					  "align": "end",
					  "text": "CASES"
					}
				  ]
				},
				{
				  "type": "box",
				  "layout": "horizontal",
				  "contents": [
					{
					  "type": "text",
					  "text": "死亡",
					  "size": "sm",
					  "color": "#555555",
					  "flex": 0
					},
					{
					  "type": "text",
					  "text": "DEATHS",
					  "size": "sm",
					  "color": "#111111",
					  "align": "end"
					}
				  ]
				},
				{
				  "type": "box",
				  "layout": "horizontal",
				  "contents": [
					{
					  "type": "text",
					  "text": "致死率",
					  "size": "sm",
					  "color": "#555555",
					  "flex": 0
					},
					{
					  "type": "text",
					  "text": "CFR",
					  "size": "sm",
					  "color": "#111111",
					  "align": "end"
					}
				  ]
				}
			  ]
			},
			{
			  "type": "separator",
			  "margin": "lg"
			},
			{
			  "type": "box",
			  "layout": "horizontal",
			  "margin": "md",
			  "contents": [
				{
				  "type": "text",
				  "text": "Countries",
				  "size": "xs",
				  "color": "#aaaaaa",
				  "flex": 0
				},
				{
				  "type": "text",
				  "text": "COUNTRIES",
				  "color": "#aaaaaa",
				  "size": "xs",
				  "align": "end"
				}
			  ]
			}
		  ]
		},
		"action": {
		  "type": "uri",
		  "label": "action",
		  "uri": "https://www.cdc.gov.tw/"
		},
		"styles": {
		  "footer": {
			"separator": true
		  }
		}
	  }
	]
  }`
