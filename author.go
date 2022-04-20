package main

import "github.com/line/line-bot-sdk-go/v7/linebot"

func GetAuthorInfo() *linebot.FlexMessage {
	container, err := linebot.UnmarshalFlexMessageJSON([]byte(personalInfoFlexMesg))
	if err != nil {
		panic(err)
	}

	return linebot.NewFlexMessage("我的基本資料", container)
}

const personalInfoFlexMesg string = `{
	"type": "bubble",
	"size": "kilo",
	"hero": {
	  "type": "image",
	  "url": "https://i.imgur.com/ONo99SU.jpg",
	  "size": "full",
	  "aspectRatio": "20:13",
	  "aspectMode": "cover",
	  "action": {
		"type": "message",
		"label": "action",
		"text": "hello!"
	  }
	},
	"body": {
	  "type": "box",
	  "layout": "vertical",
	  "contents": [
		{
		  "type": "text",
		  "text": "羅羽軒 Erin",
		  "weight": "bold",
		  "size": "xl",
		  "color": "#004B97",
		  "margin": "xs"
		},
		{
		  "type": "box",
		  "layout": "vertical",
		  "margin": "lg",
		  "spacing": "sm",
		  "contents": [
			{
			  "type": "box",
			  "layout": "baseline",
			  "spacing": "sm",
			  "contents": [
				{
				  "type": "text",
				  "text": "School",
				  "color": "#aaaaaa",
				  "size": "sm",
				  "flex": 1
				},
				{
				  "type": "text",
				  "text": "台北科技大學",
				  "wrap": true,
				  "color": "#666666",
				  "size": "sm",
				  "flex": 3
				}
			  ]
			},
			{
			  "type": "box",
			  "layout": "baseline",
			  "spacing": "none",
			  "contents": [
				{
				  "type": "text",
				  "text": "Dept.",
				  "color": "#aaaaaa",
				  "size": "sm",
				  "flex": 1
				},
				{
				  "type": "text",
				  "text": "電資學士班 大三",
				  "wrap": true,
				  "color": "#666666",
				  "size": "sm",
				  "flex": 3
				}
			  ]
			}
		  ]
		}
	  ],
	  "background": {
		"type": "linearGradient",
		"angle": "0deg",
		"startColor": "#E8FFF5",
		"endColor": "#DFFFDF"
	  }
	},
	"footer": {
	  "type": "box",
	  "layout": "horizontal",
	  "contents": [
		{
		  "type": "button",
		  "style": "link",
		  "height": "sm",
		  "action": {
			"type": "uri",
			"label": "EMAIL",
			"uri": "mailto:angelelo88362@gmail.com"
		  }
		},
		{
		  "type": "button",
		  "style": "link",
		  "height": "sm",
		  "action": {
			"type": "uri",
			"label": "GITHUB",
			"uri": "https://github.com/lohsuan"
		  }
		}
	  ],
	  "background": {
		"type": "linearGradient",
		"angle": "0deg",
		"startColor": "#D2E9FF",
		"endColor": "#E8FFF5"
	  }
	}
  }`
