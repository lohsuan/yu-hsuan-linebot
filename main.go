package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"yu-hsuan-linebot/utils"

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
					replyMsg, replyMsg2 := utils.GetOtherLocationWeather(message.Text[1:])
					_, err = bot.ReplyMessage(event.ReplyToken, replyMsg, replyMsg2).Do()
					break
				}

				switch strings.ToLower(message.Text) {
				case utils.Help, utils.UserGuild:
					replyMsg := utils.GetHelpMsg()
					replyMsg2 := utils.GetHelpWeatherMsg()
					_, err = bot.ReplyMessage(event.ReplyToken, replyMsg, replyMsg2).Do()

				case utils.QuickReply:
					replyMsg := utils.GetQuickReplyMsg()
					_, err = bot.ReplyMessage(event.ReplyToken, replyMsg).Do()

				case utils.Northern:
					replyMsg := utils.GetNorthernMsg()
					_, err = bot.ReplyMessage(event.ReplyToken, replyMsg).Do()

				case utils.Central:
					replyMsg := utils.GetCentralMsg()
					_, err = bot.ReplyMessage(event.ReplyToken, replyMsg).Do()

				case utils.Southern:
					replyMsg := utils.GetSouthernlMsg()
					_, err = bot.ReplyMessage(event.ReplyToken, replyMsg).Do()

				case utils.Eastern:
					replyMsg := utils.GetEasternlMsg()
					_, err = bot.ReplyMessage(event.ReplyToken, replyMsg).Do()

				case utils.Outlying:
					replyMsg := utils.GetOutlyingMsg()
					_, err = bot.ReplyMessage(event.ReplyToken, replyMsg).Do()

				case utils.Covid19, utils.Covid19Info:
					replyMsg := utils.GetCovidInfo()
					_, err = bot.ReplyMessage(event.ReplyToken, replyMsg).Do()

				case utils.WeatherText, utils.TaipeiWeather:
					replyMsg, replySticker := utils.GetWeatherInfo(utils.TaipeiCity)
					_, err = bot.ReplyMessage(event.ReplyToken, replyMsg, replySticker).Do()

				case utils.Author, utils.AboutAuthor:
					replyFlex := utils.GetAuthorFlex()
					replyMsg := utils.GetGreetingMsg()
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
