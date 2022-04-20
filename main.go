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

	mapset "github.com/deckarep/golang-set"

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

	locationSet := mapset.NewSet("臺北市", "新北市", "桃園市", "臺中市", "臺南市", "高雄市", "基隆市", "新竹縣", "新竹市", "苗栗縣", "彰化縣", "南投縣", "雲林縣", "嘉義縣", "嘉義市", "屏東縣", "宜蘭縣", "花蓮縣", "臺東縣", "澎湖縣", "金門縣", "連江縣")

	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:

				if message.Text[0] == '@' {
					var locationName = message.Text[1:]
					if !locationSet.Contains(locationName) {
						replyMsg := linebot.NewTextMessage("抱歉，資料庫不夠強大，請輸入下方可查詢地區")
						replyMsg2 := linebot.NewTextMessage("臺北市, 新北市, 桃園市, 臺中市, 臺南市, 高雄市, 基隆市, 新竹縣, 新竹市, 苗栗縣, 彰化縣, 南投縣, 雲林縣, 嘉義縣, 嘉義市, 屏東縣, 宜蘭縣, 花蓮縣, 臺東縣, 澎湖縣, 金門縣, 連江縣")

						if _, err = bot.ReplyMessage(event.ReplyToken, replyMsg, replyMsg2).Do(); err != nil {
							log.Print("err in linebot.TextMessage: ", err)
						}
					} else {
						replyMsg, replySticker := GetWeatherInfo(locationName)

						if _, err = bot.ReplyMessage(event.ReplyToken, replyMsg, replySticker).Do(); err != nil {
							log.Print("err in linebot.TextMessage: ", err)
						}
					}
				}

				switch message.Text {
				case "help":
					replyMsg := linebot.NewTextMessage("回覆:\nauthor -> 認識我！\njoke -> 來點冷笑話\ncovid19 -> 關注今日確診人數\nweather -> 查詢台北市天氣\n@[地名]: 查詢其他地區天氣")
					replyMsg2 := linebot.NewTextMessage("可查詢地區: 臺北市, 新北市, 桃園市, 臺中市, 臺南市, 高雄市, 基隆市, 新竹縣, 新竹市, 苗栗縣, 彰化縣, 南投縣, 雲林縣, 嘉義縣, 嘉義市, 屏東縣, 宜蘭縣, 花蓮縣, 臺東縣, 澎湖縣, 金門縣, 連江縣")

					if _, err = bot.ReplyMessage(event.ReplyToken, replyMsg, replyMsg2).Do(); err != nil {
						log.Print("err in linebot.TextMessage: ", err)
					}

				case "covid19", "今日確診人數":
					replyMsg := linebot.NewTextMessage("很高興認識你/妳！我是現在就讀北科大 電資學士班 大三的羅羽軒 Erin\n 下面是我的 github 連結，請多多指教！")
					replyLink := linebot.NewTextMessage("https://github.com/lohsuan")

					if _, err = bot.ReplyMessage(event.ReplyToken, replyMsg, replyLink).Do(); err != nil {
						log.Print("err in linebot.TextMessage: ", err)
					}

				case "weather", "今天天氣":
					var locationName = "臺北市"
					replyMsg, replySticker := GetWeatherInfo(locationName)

					if _, err = bot.ReplyMessage(event.ReplyToken, replyMsg, replySticker).Do(); err != nil {
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
					replyFlex := GetAuthorInfo()

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
