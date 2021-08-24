package controllers

import (
	. "github.com/space-tech-dev/sms-service-notification/app/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type WelcomeAPIController struct {
	ServiceWelcomeAPI *WelcomeAPIService
}

// 模块初始化函数 import 包时被调用
func init() {
	//fmt.Println("Init reservation Controller")
}

func NewWelcomeAPIController() (ctl *WelcomeAPIController) {

	return &WelcomeAPIController{
		ServiceWelcomeAPI: NewWelcomeAPIService(),
	}
}

func APIGetHome(context *gin.Context) {
	ctl := NewWelcomeAPIController()

	r := ctl.ServiceWelcomeAPI.GetWelcomeAPI()

	context.JSON(http.StatusOK, r)
}
