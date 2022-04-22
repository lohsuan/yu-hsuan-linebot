package main

import "github.com/line/line-bot-sdk-go/v7/linebot"

const (
	help      = "help"
	userGuild = "使用手冊"
)

const helpMsg = "回覆:\n" +
	"author: 認識我！\n" +
	"covid19: 關注疫情動態\n" +
	"weather: 查詢台北市天氣\n" +
	"@[地名]: 查詢其他地區天氣"

const locationHint = "可查詢地區: "

const quickReply = "快速查詢"

func GetHelpMsg() *linebot.TextMessage {
	return linebot.NewTextMessage(helpMsg)
}

func GetHelpWeatherMsg() linebot.SendingMessage {
	return linebot.NewTextMessage(locationHint + availableWeatherLocation).WithQuickReplies(linebot.NewQuickReplyItems(
		linebot.NewQuickReplyButton("", linebot.NewMessageAction(quickReply, quickReply)),
	))
}
