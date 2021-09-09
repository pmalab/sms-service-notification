package en_TW

import (
	fmt2 "fmt"
	"github.com/space-tech-dev/sms-service-notification/config"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

var lang = language.Chinese

func LoadLang() {
	message.SetString(lang, fmt2.Sprintf("%s", config.MESSAGE_HUSHAN_RESERVATION_SUCCESS_TITLE), "預約成功")
	message.SetString(lang, fmt2.Sprintf("%s", config.MESSAGE_HUSHAN_RESERVATION_SUCCESS), "您已成功預約%s %v由%s老師授課的%s")

	message.SetString(lang, fmt2.Sprintf("%s", config.MESSAGE_HUSHAN_RESERVATION_NORMAL_TO_WAIT_TITLE), "預約轉候補")
	message.SetString(lang, fmt2.Sprintf("%s", config.MESSAGE_HUSHAN_RESERVATION_NORMAL_TO_WAIT), "正取名額已滿，已幫您候補%s %v由%s老師教授的%s課程")

	message.SetString(lang, fmt2.Sprintf("%s", config.MESSAGE_HUSHAN_RESERVATION_WAIT_TITLE), "預約候補")
	message.SetString(lang, fmt2.Sprintf("%s", config.MESSAGE_HUSHAN_RESERVATION_WAIT), "您已候補%s %v由%s老師教授的%s課程")

	message.SetString(lang, fmt2.Sprintf("%s", config.MESSAGE_HUSHAN_RESERVATION_WAIT_TO_NORMAL_TITLE), "候補成功")
	message.SetString(lang, fmt2.Sprintf("%s", config.MESSAGE_HUSHAN_RESERVATION_WAIT_TO_NORMAL), "您已成功候補上%s %v由%s老師教授的%s")

	message.SetString(lang, fmt2.Sprintf("%s", config.MESSAGE_HUSHAN_RESERVATION_WAIT_TO_RESERVED_PLACE_TITLE), "候補保留名額")
	message.SetString(lang, fmt2.Sprintf("%s", config.MESSAGE_HUSHAN_RESERVATION_WAIT_TO_RESERVED_PLACE), "您候補的%s %v由%s老師教授的%s，已幫您做上課的名額保留")

	message.SetString(lang, fmt2.Sprintf("%s", config.MESSAGE_HUSHAN_RESERVATION_CANCEL_TITLE), "預約取消")
	message.SetString(lang, fmt2.Sprintf("%s", config.MESSAGE_HUSHAN_RESERVATION_CANCEL), "您已取消%s %v由%s老師教授的%s課程預約")

	message.SetString(lang, fmt2.Sprintf("%s", config.MESSAGE_HUSHAN_RESERVATION_UPDATE_CYCLE_SEAT_TITLE), "更換車位成功")
	message.SetString(lang, fmt2.Sprintf("%s", config.MESSAGE_HUSHAN_RESERVATION_UPDATE_CYCLE_SEAT), "您的預約(%s %v由%s老師教授的%s)已成功更換車位")

	message.SetString(lang, fmt2.Sprintf("%s", config.MESSAGE_CYCLE_SEAT_NUMBER), "%v號車位")

	message.SetString(lang, fmt2.Sprintf("%s", config.API_ERR_CODE_FAIL_TO_GET_THIRD_PARTY_SHOP_LIST), "无法获取第三方门店信息")
}
