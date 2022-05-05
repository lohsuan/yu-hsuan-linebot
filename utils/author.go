package utils

import (
	"log"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

const Author = "author"
const AboutAuthor = "認識作者"

const myPersonalInfo = "我的基本資料"
const authorName = "羽軒 Erin"
const sallyImageUrl = "https://stickershop.line-scdn.net/stickershop/v1/sticker/52002736/iPhone/sticker_key@2x.png"

func GetAuthorFlex() linebot.SendingMessage {
	sender := linebot.NewSender(authorName, sallyImageUrl)
	container, err := linebot.UnmarshalFlexMessageJSON([]byte(personalInfoFlexMsg))
	if err != nil {
		log.Print(err)
	}
	return linebot.NewFlexMessage(myPersonalInfo, container).WithSender(sender)
}

func GetGreetingMsg() linebot.SendingMessage {
	sender := linebot.NewSender(authorName, sallyImageUrl)
	return linebot.NewTextMessage("Nice to meet you!$").WithSender(sender).AddEmoji(linebot.NewEmoji(17, "5ac2213e040ab15980c9b447", "035"))
}

const personalInfoFlexMsg string = `{
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
	  ]
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
	  ]
	}
  }`
