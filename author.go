package main

import "github.com/line/line-bot-sdk-go/v7/linebot"

func GetAuthorInfo() *linebot.FlexMessage {
	container, err := linebot.UnmarshalFlexMessageJSON([]byte(`{
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
			  "color": "#2E4057"
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
				  "spacing": "sm",
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
		  ]
		},
		"footer": {
		  "type": "box",
		  "layout": "vertical",
		  "spacing": "sm",
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
			},
			{
			  "type": "box",
			  "layout": "vertical",
			  "contents": [],
			  "margin": "sm"
			}
		  ],
		  "flex": 0
		}
	  }`))
	if err != nil {
		panic(err)
	}

	return linebot.NewFlexMessage("我的基本資料", container)
}
