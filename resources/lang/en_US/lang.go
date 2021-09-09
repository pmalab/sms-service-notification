package en_US

import (
	fmt2 "fmt"
	"github.com/space-tech-dev/sms-service-notification/config"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

var lang = language.English

func LoadLang() {
	message.SetString(lang, fmt2.Sprintf("%d", config.API_RETURN_CODE_ERROR), "return error")
	message.SetString(lang, fmt2.Sprintf("%s", config.MESSAGE_HUSHAN_RESERVATION_SUCCESS_TITLE), "Booking Success")
	message.SetString(lang, fmt2.Sprintf("%s", config.MESSAGE_HUSHAN_RESERVATION_SUCCESS), "%s %v %s %s")

	message.SetString(lang, fmt2.Sprintf("%s", config.MESSAGE_HUSHAN_RESERVATION_NORMAL_TO_WAIT_TITLE), "Booking Convert To Wait List")
	message.SetString(lang, fmt2.Sprintf("%s", config.MESSAGE_HUSHAN_RESERVATION_NORMAL_TO_WAIT), "%s %v %s %s")

	message.SetString(lang, fmt2.Sprintf("%s", config.MESSAGE_HUSHAN_RESERVATION_WAIT_TITLE), "Booking Wait List")
	message.SetString(lang, fmt2.Sprintf("%s", config.MESSAGE_HUSHAN_RESERVATION_WAIT), "%s %v %s %s")

	message.SetString(lang, fmt2.Sprintf("%s", config.MESSAGE_HUSHAN_RESERVATION_WAIT_TO_NORMAL_TITLE), "Booking Wait List Success")
	message.SetString(lang, fmt2.Sprintf("%s", config.MESSAGE_HUSHAN_RESERVATION_WAIT_TO_NORMAL), "%s %v %s %s")

	message.SetString(lang, fmt2.Sprintf("%s", config.MESSAGE_HUSHAN_RESERVATION_WAIT_TO_RESERVED_PLACE_TITLE), "Booking Wait List Reserved Place")
	message.SetString(lang, fmt2.Sprintf("%s", config.MESSAGE_HUSHAN_RESERVATION_WAIT_TO_RESERVED_PLACE), "%s %v %s %s")

	message.SetString(lang, fmt2.Sprintf("%s", config.MESSAGE_HUSHAN_RESERVATION_CANCEL_TITLE), "Booking Cancel")
	message.SetString(lang, fmt2.Sprintf("%s", config.MESSAGE_HUSHAN_RESERVATION_CANCEL), "%s %v %s %s")

	message.SetString(lang, fmt2.Sprintf("%s", config.MESSAGE_HUSHAN_RESERVATION_UPDATE_CYCLE_SEAT_TITLE), "Change Seat Success")
	message.SetString(lang, fmt2.Sprintf("%s", config.MESSAGE_HUSHAN_RESERVATION_UPDATE_CYCLE_SEAT), "%s %v %s %s")

	message.SetString(lang, fmt2.Sprintf("%s", config.MESSAGE_CYCLE_SEAT_NUMBER), "seat number %v")

	message.SetString(lang, fmt2.Sprintf("%d", config.API_ERR_CODE_FAIL_TO_GET_THIRD_PARTY_SHOP_LIST), "无法获取第三方门店信息")
}
