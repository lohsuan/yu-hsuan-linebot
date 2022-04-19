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
	bot, err := linebot.New(os.Getenv("ChannelSecret"), os.Getenv("ChannelAccessToken"))
	log.Println("Bot:", bot, " err:", err)

	http.HandleFunc("/callback", callbackHandler)

	port := os.Getenv("PORT")
	addr := fmt.Sprintf(":%s", port)

	http.ListenAndServe(addr, nil)
}

func EmojiMsg() linebot.SendingMessage {
	return linebot.NewTextMessage(
		fmt.Sprintf("哈囉~ 我是你的小助理！$\n 祝你有美好的一天!")).AddEmoji(
		linebot.NewEmoji(12, "5ac1bfd5040ab15980c9b435", "043"))
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
				if _, err = bot.ReplyMessage(event.ReplyToken, EmojiMsg()).Do(); err != nil {
					log.Print(err)
				}

			case *linebot.StickerMessage:
				var kw string
				for _, k := range message.Keywords {
					kw = kw + "," + k
				}

				outStickerResult := fmt.Sprintf("收到貼圖訊息: %s, pkg: %s kw: %s  text: %s", message.StickerID, message.PackageID, kw, message.Text)
				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(outStickerResult)).Do(); err != nil {
					log.Print(err)
				}
			}
		}
	}
}
