package service

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v2/client"
	"github.com/alibabacloud-go/tea/tea"
	. "github.com/space-tech-dev/sms-service-notification/config"
	"log"
)

var (
	AliYunClient *dysmsapi20170525.Client
)

type AliYunService struct {
	//NotificationLog *NotificationLog
}

/**
** 初始化构造函数
 */

// 模块初始化函数 import 包时被调用
func init() {
	//fmt.Println("notification service module init function")
}

func NewAliYunService() (aliYunSrv *AliYunService) {
	aliYunSrv = &AliYunService{
		//NotificationLog: NewNotificationLog(),
	}
	return aliYunSrv
}

// CreateClient /**
func CreateClient (aliYun *AliYunConfig) (_result *dysmsapi20170525.Client, _err error) {
	config := &openapi.Config{
		// 您的AccessKey ID
		AccessKeyId: &aliYun.AccessKeyId,
		// 您的AccessKey S
		AccessKeySecret: &aliYun.AccessKeySecret,
	}
	// 访问的域名
	config.Endpoint = tea.String("dysmsapi.aliyuncs.com")
	_result = &dysmsapi20170525.Client{}
	_result, _err = dysmsapi20170525.NewClient(config)
	if _err != nil {
		// throw an exception here
		log.Fatal("Aliyun Auth init error: ", _err)
		return
	}
	AliYunClient = _result
	return _result, _err
}

func (srv *AliYunService) sendMsg()  {

}

