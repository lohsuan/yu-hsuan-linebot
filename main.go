package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

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

				if message.Text[0] == '@' {
					replyMsg, replyMsg2 := GetOtherLocationWeather(message.Text[1:])
					_, err = bot.ReplyMessage(event.ReplyToken, replyMsg, replyMsg2).Do()
					break
				}

				switch strings.ToLower(message.Text) {
				case help, userGuild:
					replyMsg := GetHelpMsg()
					replyMsg2 := GetHelpWeatherMsg()
					_, err = bot.ReplyMessage(event.ReplyToken, replyMsg, replyMsg2).Do()

				case quickReply:
					replyMsg := GetQuickReplyMsg()
					_, err = bot.ReplyMessage(event.ReplyToken, replyMsg).Do()

				case northern:
					replyMsg := GetNorthernMsg()
					_, err = bot.ReplyMessage(event.ReplyToken, replyMsg).Do()

				case central:
					replyMsg := GetCentralMsg()
					_, err = bot.ReplyMessage(event.ReplyToken, replyMsg).Do()

				case southern:
					replyMsg := GetSouthernlMsg()
					_, err = bot.ReplyMessage(event.ReplyToken, replyMsg).Do()

				case eastern:
					replyMsg := GetEasternlMsg()
					_, err = bot.ReplyMessage(event.ReplyToken, replyMsg).Do()

				case outlying:
					replyMsg := GetOutlyingMsg()
					_, err = bot.ReplyMessage(event.ReplyToken, replyMsg).Do()

				case covid19, covid19Info:
					replyMsg := GetCovidInfo()
					_, err = bot.ReplyMessage(event.ReplyToken, replyMsg).Do()

				case weather, taipeiWeather:
					replyMsg, replySticker := GetWeatherInfo(TaipeiCity)
					_, err = bot.ReplyMessage(event.ReplyToken, replyMsg, replySticker).Do()

				case author, aboutAuthor:
					replyFlex := GetAuthorFlex()
					replyMsg := GetGreetingMsg()
					_, err = bot.ReplyMessage(event.ReplyToken, replyFlex, replyMsg).Do()

				default:
					replyMsg, stickerMsg := GetDefaultMsg()
					_, err = bot.ReplyMessage(event.ReplyToken, replyMsg, stickerMsg).Do()

				}

			default:
				replyMsg, stickerMsg := GetDefaultMsg()
				_, err = bot.ReplyMessage(event.ReplyToken, replyMsg, stickerMsg).Do()

			}
			if err != nil {
				log.Print(err)
			}
		}
	}
}

func GetDefaultMsg() (*linebot.TextMessage, *linebot.StickerMessage) {
	const defaultMsg = "你的小助理上線啦！回覆 help 可檢視更多功能，祝你有美好的一天:)"
	replyMsg := linebot.NewTextMessage(defaultMsg)
	stickerMsg := linebot.NewStickerMessage("2", "514")
	return replyMsg, stickerMsg
}
