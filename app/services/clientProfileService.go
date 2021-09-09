package service

import (
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/gin-gonic/gin"
	"github.com/golang-module/carbon"
	"github.com/space-tech-dev/sms-service-notification/app/models"
)

type ClientProfileService struct {
	Service *Service
}

const PLATFORM_SALESFORCE = "Salesforce"
const PLATFORM_WECHAT_MINI_PROGRAM = "WeChat Mini Program"
const PLATFORM_IOS = "iOS"
const PLATFORM_ANDROID = "Android"
const PLATFORM_RETAIL = "Retail"
const PLATFORM_JD = "JD"
const PLATFORM_TMALL = "TMall"
const PLATFORM_DIANPING = "DianPing"
const PLATFORM_ALL = "All"
const PLATFORM_MGM = "MGM"
const PLATFORM_APP = "App"
const PLATFORM_WEBSITE = "Website"
const PLATFORM_TAIPEI_OLD_CYCLE_SYSTEM = "Taipei Old Cycle System"

const OS_TYPE_IOS = 1
const OS_TYPE_ANDROID = 2

const TIMEZONE = carbon.UCT
const REQUEST_TIMEZONE = carbon.Taipei

var ARRAY_PLATFORM = []string{
	PLATFORM_SALESFORCE,
	PLATFORM_WECHAT_MINI_PROGRAM,
	PLATFORM_IOS,
	PLATFORM_ANDROID,
	PLATFORM_RETAIL,
	PLATFORM_JD,
	PLATFORM_TMALL,
	PLATFORM_DIANPING,
	PLATFORM_ALL,
	PLATFORM_MGM,
	PLATFORM_APP,
	PLATFORM_WEBSITE,
	PLATFORM_TAIPEI_OLD_CYCLE_SYSTEM,
}

var ARRAY_OS_TYPE = []int{
	OS_TYPE_IOS,
	OS_TYPE_ANDROID,
}



/**
 ** 初始化构造函数
 */

// 模块初始化函数 import 包时被调用
func init() {
	//fmt.Println("clientProfile service module init function")
}

func NewClientProfileService(context *gin.Context) (r *ClientProfileService) {
	r = &ClientProfileService{
		Service: NewService(context),
	}
	return r
}

/**
 ** 静态函数
 */
func GetCurrentPlatform(c *gin.Context) string {
	return c.GetHeader("platform")

}

func GetCurrentChannel(c *gin.Context) string {

	return c.GetHeader("channel")
}

func GetCurrentUUID(c *gin.Context) string {

	return c.GetHeader("uuid")
}

func GetCurrentSource(c *gin.Context) string {

	return c.GetHeader("source")
}

func IsPlatformSF(c *gin.Context, platform string) bool {
	if platform == "" {
		platform = GetCurrentPlatform(c)
	}

	return platform == PLATFORM_SALESFORCE

}

func GetSessionLocale(c *gin.Context) string {

	// client cannot override the message locales.
	locale := GetRequestLocale(c)
	//        dd($locale);
	if locale == "" {
		locale = models.LOCALE_TW
	}
	return locale
}

func GetRequestLocale(c *gin.Context) string {

	requestLocal := c.GetHeader("locale")
	if requestLocal != "" && object.InArray(requestLocal, models.ARRAY_LOCALE) {
		return requestLocal
	}
	return ""
}

/**
 ** 实例函数
 */