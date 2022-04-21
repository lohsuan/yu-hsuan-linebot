package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

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
						replyMsg := linebot.NewTextMessage("查無此地區資料，請輸入或點選快速查詢><")
						replyMsg2 := linebot.NewTextMessage("臺北市, 新北市, 桃園市, 臺中市, 臺南市, 高雄市, 基隆市, 新竹縣, 新竹市, 苗栗縣, 彰化縣, 南投縣, 雲林縣, 嘉義縣, 嘉義市, 屏東縣, 宜蘭縣, 花蓮縣, 臺東縣, 澎湖縣, 金門縣, 連江縣")
						replyMsg2.WithQuickReplies(linebot.NewQuickReplyItems(
							linebot.NewQuickReplyButton("", linebot.NewMessageAction("快速查詢", "快速查詢")),
						))
						if _, err = bot.ReplyMessage(event.ReplyToken, replyMsg, replyMsg2).Do(); err != nil {
							log.Print("err in linebot.TextMessage: ", err)
						}
					} else {
						replyMsg, replySticker := GetWeatherInfo(locationName)

						if _, err = bot.ReplyMessage(event.ReplyToken, replyMsg, replySticker).Do(); err != nil {
							log.Print("err in linebot.TextMessage: ", err)
						}
					}
					break
				}

				switch strings.ToLower(message.Text) {
				case help, userGuild:
					replyMsg := GetHelpMesg()
					replyMsg2 := GetHelpWeatherMesg()
					_, err = bot.ReplyMessage(event.ReplyToken, replyMsg, replyMsg2).Do()

				case quickReply:
					replyMsg := GetQuickReplyMesg()
					bot.ReplyMessage(event.ReplyToken, replyMsg).Do()

				case northern:
					replyMsg := GetNorthernMesg()
					bot.ReplyMessage(event.ReplyToken, replyMsg).Do()

				case central:
					replyMsg := GetNorthernMesg()
					bot.ReplyMessage(event.ReplyToken, replyMsg).Do()

				case southern:
					replyMsg := GetSouthernlMesg()
					bot.ReplyMessage(event.ReplyToken, replyMsg).Do()

				case eastern:
					replyMsg := GetEasternlMesg()
					bot.ReplyMessage(event.ReplyToken, replyMsg).Do()

				case outlying:
					replyMsg := GetOutlyingMesg()
					bot.ReplyMessage(event.ReplyToken, replyMsg).Do()

				case "covid19", "關注疫情動態":
					sendr := linebot.NewSender("疾管署", "https://i.imgur.com/ZvY23Ag.png")
					replyMsg, err := GetCovidInfo()
					if err != nil {
						bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("資料發生錯誤，請稍後再試")).Do()
						break
					}
					bot.ReplyMessage(event.ReplyToken, replyMsg.WithSender(sendr)).Do()

				case "weather", "臺北市天氣":
					replyMsg, replySticker := GetWeatherInfo("臺北市")
					bot.ReplyMessage(event.ReplyToken, replyMsg, replySticker).Do()

				case "author", "認識作者":
					replyMsg := GetGreetingMesg()
					replyFlex := GetAuthorFlex()
					bot.ReplyMessage(event.ReplyToken, replyFlex, replyMsg).Do()

				default:
					replyMsg, stickerMsg := GetDefaultMesg()
					bot.ReplyMessage(event.ReplyToken, replyMsg, stickerMsg).Do()
				}

			case *linebot.StickerMessage:
				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewStickerMessage("2", "514")).Do(); err != nil {
					log.Print("err in linebot.StickerMessage: ", err)
				}

			default:
				replyMsg, stickerMsg := GetDefaultMesg()
				bot.ReplyMessage(event.ReplyToken, replyMsg, stickerMsg).Do()

			}
			if err != nil {
				log.Print(err)
			}
		}
	}
}

func GetDefaultMesg() (*linebot.TextMessage, *linebot.StickerMessage) {
	const defaultMesg = "你的小助理上線啦！回覆 help 可檢視更多功能，祝你有美好的一天:)"

	replyMsg := linebot.NewTextMessage(defaultMesg)
	stickerMsg := linebot.NewStickerMessage("2", "514")
	return replyMsg, stickerMsg
}
