package main

import "github.com/line/line-bot-sdk-go/v7/linebot"

const (
	help      = "help"
	userGuild = "使用手冊"
)

const helpMesg = "回覆:\n" +
	"author: 認識我！\n" +
	"covid19: 關注疫情動態\n" +
	"weather: 查詢台北市天氣\n" +
	"@[地名]: 查詢其他地區天氣"

const availableWeatherLocation = "可查詢地區: 臺北市, 新北市, 桃園市, " +
	"臺中市, 臺南市, 高雄市, 基隆市, 新竹縣, 新竹市, 苗栗縣, 彰化縣, 南投縣, " +
	"雲林縣, 嘉義縣, 嘉義市, 屏東縣, 宜蘭縣, 花蓮縣, 臺東縣, 澎湖縣, 金門縣, 連江縣"

const quickReply = "快速查詢"

func GetHelpMesg() *linebot.TextMessage {
	return linebot.NewTextMessage(helpMesg)
}

func GetHelpWeatherMesg() linebot.SendingMessage {
	return linebot.NewTextMessage(availableWeatherLocation).WithQuickReplies(linebot.NewQuickReplyItems(
		linebot.NewQuickReplyButton("", linebot.NewMessageAction(quickReply, quickReply)),
	))
}
