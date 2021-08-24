package service

import (
	. "github.com/space-tech-dev/sms-service-notification/app/models"
)

type WelcomeAPIService struct {
	User *User
}


/**
 ** 初始化构造函数
*/

// 模块初始化函数 import 包时被调用
func init() {
	//fmt.Println("reservation service module init function")
}

func NewWelcomeAPIService() (r * WelcomeAPIService)  {
	r = &WelcomeAPIService{
		User: NewUser(),
	}
	return r
}


/**
 ** 实例函数
 */

func (srv *WelcomeAPIService) GetWelcomeAPI() string {

	return "Welcome API!"

}
