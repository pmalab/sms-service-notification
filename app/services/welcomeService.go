package service

import (
	. "github.com/space-tech-dev/sms-service-notification/app/models"
	"github.com/space-tech-dev/sms-service-notification/config"
)

type WelcomeService struct {
	User *User
}


/**
 ** 初始化构造函数
*/

// 模块初始化函数 import 包时被调用
func init() {
	//fmt.Println("reservation service module init function")
}

func NewWelcomeService() (r * WelcomeService)  {
	r = &WelcomeService{
		User: NewUser(),
	}
	return r
}


/**
 ** 实例函数
 */

func (srv *WelcomeService) GetWelcome() string {

	return "Welcome! " + config.APP_NAME + " version:" + config.APP_VERSION

}
