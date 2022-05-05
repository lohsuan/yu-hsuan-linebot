package utils

import (
	mapset "github.com/deckarep/golang-set"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

const (
	Northern       = "北部"
	Central        = "中部"
	Southern       = "南部"
	Eastern        = "東部"
	Outlying       = "離島"
	ChooseLocation = "選擇地區"
	ChoosePlace    = "選擇縣市"
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

const availableWeatherLocation = "臺北市, 新北市, 桃園市, " +
	"臺中市, 臺南市, 高雄市, 基隆市, 新竹縣, 新竹市, 苗栗縣, 彰化縣, 南投縣, " +
	"雲林縣, 嘉義縣, 嘉義市, 屏東縣, 宜蘭縣, 花蓮縣, 臺東縣, 澎湖縣, 金門縣, 連江縣"

const locationNotAvailableMsg = "查無此地區資料，請輸入以下地區或點選快速查詢><"

func GetQuickReplyMsg() linebot.SendingMessage {
	quickReplyItems := linebot.NewQuickReplyItems(
		linebot.NewQuickReplyButton("", linebot.NewMessageAction(Northern, Northern)),
		linebot.NewQuickReplyButton("", linebot.NewMessageAction(Central, Central)),
		linebot.NewQuickReplyButton("", linebot.NewMessageAction(Southern, Southern)),
		linebot.NewQuickReplyButton("", linebot.NewMessageAction(Eastern, Eastern)),
		linebot.NewQuickReplyButton("", linebot.NewMessageAction(Outlying, Outlying)),
	)
	return linebot.NewTextMessage(ChooseLocation).WithQuickReplies(quickReplyItems)
}

func GetNorthernMsg() linebot.SendingMessage {
	quickReplyItems := linebot.NewQuickReplyItems(
		linebot.NewQuickReplyButton("", linebot.NewMessageAction(KeelungCity, "@"+KeelungCity)),
		linebot.NewQuickReplyButton("", linebot.NewMessageAction(TaipeiCity, "@"+TaipeiCity)),
		linebot.NewQuickReplyButton("", linebot.NewMessageAction(NewTaipeiCity, "@"+NewTaipeiCity)),
		linebot.NewQuickReplyButton("", linebot.NewMessageAction(YilanCounty, "@"+YilanCounty)),
		linebot.NewQuickReplyButton("", linebot.NewMessageAction(TaoyuanCity, "@"+TaoyuanCity)),
		linebot.NewQuickReplyButton("", linebot.NewMessageAction(HsinchuCity, "@"+HsinchuCity)),
	)
	return linebot.NewTextMessage(ChoosePlace).WithQuickReplies(quickReplyItems)
}

func GetCentralMsg() linebot.SendingMessage {
	quickReplyItems := linebot.NewQuickReplyItems(
		linebot.NewQuickReplyButton("", linebot.NewMessageAction(MiaoliCounty, "@"+MiaoliCounty)),
		linebot.NewQuickReplyButton("", linebot.NewMessageAction(TaichungCity, "@"+TaichungCity)),
		linebot.NewQuickReplyButton("", linebot.NewMessageAction(ChanghuaCounty, "@"+ChanghuaCounty)),
		linebot.NewQuickReplyButton("", linebot.NewMessageAction(YunlinCounty, "@"+YunlinCounty)),
		linebot.NewQuickReplyButton("", linebot.NewMessageAction(NantouCounty, "@"+NantouCounty)),
	)
	return linebot.NewTextMessage(ChoosePlace).WithQuickReplies(quickReplyItems)
}

func GetSouthernlMsg() linebot.SendingMessage {
	quickReplyItems := linebot.NewQuickReplyItems(
		linebot.NewQuickReplyButton("", linebot.NewMessageAction(ChiayiCity, "@"+ChiayiCity)),
		linebot.NewQuickReplyButton("", linebot.NewMessageAction(ChiayiCounty, "@"+ChiayiCounty)),
		linebot.NewQuickReplyButton("", linebot.NewMessageAction(TainanCity, "@"+TainanCity)),
		linebot.NewQuickReplyButton("", linebot.NewMessageAction(KaohsiungCity, "@"+KaohsiungCity)),
		linebot.NewQuickReplyButton("", linebot.NewMessageAction(PingtungCounty, "@"+PingtungCounty)),
	)
	return linebot.NewTextMessage(ChoosePlace).WithQuickReplies(quickReplyItems)
}

func GetEasternlMsg() linebot.SendingMessage {
	quickReplyItems := linebot.NewQuickReplyItems(
		linebot.NewQuickReplyButton("", linebot.NewMessageAction(HualienCounty, "@"+HualienCounty)),
		linebot.NewQuickReplyButton("", linebot.NewMessageAction(TaitungCounty, "@"+TaitungCounty)),
	)
	return linebot.NewTextMessage(ChoosePlace).WithQuickReplies(quickReplyItems)
}

func GetOutlyingMsg() linebot.SendingMessage {
	quickReplyItems := linebot.NewQuickReplyItems(
		linebot.NewQuickReplyButton("", linebot.NewMessageAction(PenghuCounty, "@"+PenghuCounty)),
		linebot.NewQuickReplyButton("", linebot.NewMessageAction(KinmenCounty, "@"+KinmenCounty)),
		linebot.NewQuickReplyButton("", linebot.NewMessageAction(LianjiangCounty, "@"+LianjiangCounty)),
	)
	return linebot.NewTextMessage(ChoosePlace).WithQuickReplies(quickReplyItems)
}

func GetOtherLocationWeather(locationName string) (linebot.SendingMessage, linebot.SendingMessage) {
	locationSet := mapset.NewSet(KeelungCity, TaipeiCity, NewTaipeiCity, YilanCounty, TaoyuanCity, HsinchuCity, MiaoliCounty, TaichungCity, ChanghuaCounty, YunlinCounty, NantouCounty, ChiayiCity, ChiayiCounty, TainanCity, KaohsiungCity, PingtungCounty, HualienCounty, TaitungCounty, PenghuCounty, KinmenCounty, LianjiangCounty)

	if !locationSet.Contains(locationName) {
		locationNotAvailable := linebot.NewTextMessage(locationNotAvailableMsg)
		quickReplyItems := linebot.NewQuickReplyItems(
			linebot.NewQuickReplyButton("", linebot.NewMessageAction(QuickReply, QuickReply)),
		)
		availableLocation := linebot.NewTextMessage(availableWeatherLocation).WithQuickReplies(quickReplyItems)
		return locationNotAvailable, availableLocation
	} else {
		return GetWeatherInfo(locationName)
	}
}
