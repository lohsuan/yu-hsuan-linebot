package main

import "github.com/line/line-bot-sdk-go/v7/linebot"

const (
	northern       = "北部"
	central        = "中部"
	southern       = "南部"
	eastern        = "東部"
	outlying       = "離島"
	chooseLocation = "選擇地區"
	choosePlace    = "選擇縣市"
)

const (
	KeelungCity   = "基隆市"
	TaipeiCity    = "臺北市"
	NewTaipeiCity = "新北市"
	YilanCounty   = "宜蘭縣"
	TaoyuanCity   = "桃園市"
	HsinchuCity   = "新竹市"
)

const (
	MiaoliCounty   = "苗栗縣"
	TaichungCity   = "臺中市"
	ChanghuaCounty = "彰化縣"
	YunlinCounty   = "雲林縣"
	NantouCounty   = "南投縣"
)

const (
	ChiayiCity     = "嘉義市"
	ChiayiCounty   = "嘉義縣"
	TainanCity     = "臺南市"
	KaohsiungCity  = "高雄市"
	PingtungCounty = "屏東縣"
)

const (
	HualienCounty = "花蓮縣"
	TaitungCounty = "臺東縣"
)

const (
	PenghuCounty    = "澎湖縣"
	KinmenCounty    = "金門縣"
	LianjiangCounty = "連江縣"
)

func GetQuickReplyMesg() linebot.SendingMessage {
	quickReplyItems := linebot.NewQuickReplyItems(
		linebot.NewQuickReplyButton("", linebot.NewMessageAction(northern, northern)),
		linebot.NewQuickReplyButton("", linebot.NewMessageAction(central, central)),
		linebot.NewQuickReplyButton("", linebot.NewMessageAction(southern, southern)),
		linebot.NewQuickReplyButton("", linebot.NewMessageAction(eastern, eastern)),
		linebot.NewQuickReplyButton("", linebot.NewMessageAction(outlying, outlying)),
	)
	return linebot.NewTextMessage(chooseLocation).WithQuickReplies(quickReplyItems)
}

func GetNorthernMesg() linebot.SendingMessage {
	quickReplyItems := linebot.NewQuickReplyItems(
		linebot.NewQuickReplyButton("", linebot.NewMessageAction(KeelungCity, "@"+KeelungCity)),
		linebot.NewQuickReplyButton("", linebot.NewMessageAction(TaipeiCity, "@"+TaipeiCity)),
		linebot.NewQuickReplyButton("", linebot.NewMessageAction(NewTaipeiCity, "@"+NewTaipeiCity)),
		linebot.NewQuickReplyButton("", linebot.NewMessageAction(YilanCounty, "@"+YilanCounty)),
		linebot.NewQuickReplyButton("", linebot.NewMessageAction(TaoyuanCity, "@"+TaoyuanCity)),
		linebot.NewQuickReplyButton("", linebot.NewMessageAction(HsinchuCity, "@"+HsinchuCity)),
	)
	return linebot.NewTextMessage(choosePlace).WithQuickReplies(quickReplyItems)
}

func GetSouthernlMesg() linebot.SendingMessage {
	quickReplyItems := linebot.NewQuickReplyItems(
		linebot.NewQuickReplyButton("", linebot.NewMessageAction(ChiayiCity, "@"+ChiayiCity)),
		linebot.NewQuickReplyButton("", linebot.NewMessageAction(ChiayiCounty, "@"+ChiayiCounty)),
		linebot.NewQuickReplyButton("", linebot.NewMessageAction(TainanCity, "@"+TainanCity)),
		linebot.NewQuickReplyButton("", linebot.NewMessageAction(KaohsiungCity, "@"+KaohsiungCity)),
		linebot.NewQuickReplyButton("", linebot.NewMessageAction(PingtungCounty, "@"+PingtungCounty)),
	)
	return linebot.NewTextMessage(choosePlace).WithQuickReplies(quickReplyItems)
}

func GetEasternlMesg() linebot.SendingMessage {
	quickReplyItems := linebot.NewQuickReplyItems(
		linebot.NewQuickReplyButton("", linebot.NewMessageAction(HualienCounty, "@"+HualienCounty)),
		linebot.NewQuickReplyButton("", linebot.NewMessageAction(TaitungCounty, "@"+TaitungCounty)),
	)
	return linebot.NewTextMessage(choosePlace).WithQuickReplies(quickReplyItems)
}

func GetOutlyingMesg() linebot.SendingMessage {
	quickReplyItems := linebot.NewQuickReplyItems(
		linebot.NewQuickReplyButton("", linebot.NewMessageAction(PenghuCounty, "@"+PenghuCounty)),
		linebot.NewQuickReplyButton("", linebot.NewMessageAction(KinmenCounty, "@"+KinmenCounty)),
		linebot.NewQuickReplyButton("", linebot.NewMessageAction(LianjiangCounty, "@"+LianjiangCounty)),
	)
	return linebot.NewTextMessage(choosePlace).WithQuickReplies(quickReplyItems)
}
