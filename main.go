// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

var bot *linebot.Client

func main() {
	var err error
	bot, err = linebot.New(os.Getenv("ChannelSecret"), os.Getenv("ChannelAccessToken"))
	log.Println("Bot:", bot, " err:", err)
	http.HandleFunc("/callback", callbackHandler)
	port := os.Getenv("PORT")
	addr := fmt.Sprintf(":%s", port)
	http.ListenAndServe(addr, nil)
}

func GetPersonalInfoFlex() *linebot.FlexMessage {
	container, err := linebot.UnmarshalFlexMessageJSON([]byte(`{
		"type": "bubble",
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

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	events, err := bot.ParseRequest(r)

	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		return
	}

	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:

				switch message.Text {
				case "help":
					replyMsg := linebot.NewTextMessage("回覆:\nauthor -> 認識我！\njoke -> 生活太無聊，來點冷笑話xD\nepidemic -> 關注今日確診人數\nweather -> 今天天氣如何勒")

					if _, err = bot.ReplyMessage(event.ReplyToken, replyMsg).Do(); err != nil {
						log.Print("err in linebot.TextMessage: ", err)
					}

				case "epidemic", "今日確診人數":
					replyMsg := linebot.NewTextMessage("很高興認識你/妳！我是現在就讀北科大 電資學士班 大三的羅羽軒 Erin\n 下面是我的 github 連結，請多多指教！")
					replyLink := linebot.NewTextMessage("https://github.com/lohsuan")

					if _, err = bot.ReplyMessage(event.ReplyToken, replyMsg, replyLink).Do(); err != nil {
						log.Print("err in linebot.TextMessage: ", err)
					}

				case "weather", "今天天氣":
					replyMsg := linebot.NewTextMessage("很高興認識你/妳！我是現在就讀北科大 電資學士班 大三的羅羽軒 Erin\n 下面是我的 github 連結，請多多指教！")
					replyLink := linebot.NewTextMessage("https://github.com/lohsuan")

					if _, err = bot.ReplyMessage(event.ReplyToken, replyMsg, replyLink).Do(); err != nil {
						log.Print("err in linebot.TextMessage: ", err)
					}

				case "joke", "來點冷笑話":
					replyMsg := linebot.NewTextMessage("很高興認識你/妳！我是現在就讀北科大 電資學士班 大三的羅羽軒 Erin\n 下面是我的 github 連結，請多多指教！")
					replyLink := linebot.NewTextMessage("https://github.com/lohsuan")

					if _, err = bot.ReplyMessage(event.ReplyToken, replyMsg, replyLink).Do(); err != nil {
						log.Print("err in linebot.TextMessage: ", err)
					}

				case "author", "認識作者":
					replyMsg := linebot.NewTextMessage("初次見面，請多多指教！")
					replyFlex := GetPersonalInfoFlex()

					if _, err = bot.ReplyMessage(event.ReplyToken, replyMsg, replyFlex).Do(); err != nil {
						log.Print("err in linebot.TextMessage: ", err)
					}

				default:
					replyMsg := linebot.NewTextMessage("你的小助理上線啦，回覆 help 可檢視更多功能，祝你有美好的一天唷:)")
					stickerMsg := linebot.NewStickerMessage("2", "514")

					if _, err = bot.ReplyMessage(event.ReplyToken, replyMsg, stickerMsg).Do(); err != nil {
						log.Print("err in linebot.TextMessage: ", err)
					}
				}

			case *linebot.StickerMessage:
				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewStickerMessage("2", "514")).Do(); err != nil {
					log.Print("err in linebot.StickerMessage: ", err)
				}

			}
		}
	}
}
