package controllers

import (
	. "github.com/space-tech-dev/sms-service-notification/app/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type WelcomeController struct {
	ServiceWelcome *WelcomeService
}

// 模块初始化函数 import 包时被调用
func init() {
	//fmt.Println("Init reservation Controller")
}

func NewWelcomeController() (ctl *WelcomeController) {

	return &WelcomeController{
		ServiceWelcome: NewWelcomeService(),
	}
}

func WebGetHome(context *gin.Context) {
	ctl := NewWelcomeController()

	r := ctl.ServiceWelcome.GetWelcome()

	context.JSON(http.StatusOK, r)
}
