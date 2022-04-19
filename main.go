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
				case "author":
					replyMsg := linebot.NewTextMessage("Author: Yu-Hsuan, Lo (Erin)\nGithub: https://github.com/lohsuan")

					if _, err = bot.ReplyMessage(event.ReplyToken, replyMsg).Do(); err != nil {
						log.Print("err in linebot.TextMessage: ", err)
					}

				default:
					replyMsg := linebot.NewTextMessage("小助理上線啦，回覆 help 可取得更多功能\n祝你有美好的一天唷:)")
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
