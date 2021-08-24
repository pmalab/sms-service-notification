package lang

import (
	"github.com/space-tech-dev/sms-service-notification/resources/lang/en_US"
	"github.com/space-tech-dev/sms-service-notification/resources/lang/zh_TW"
)

func LoadLanguages()  {
	en_US.LoadLang()
	en_TW.LoadLang()
}