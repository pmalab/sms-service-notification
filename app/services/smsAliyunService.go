package service

import (
	"fmt"
	UBT "github.com/ArtisanCloud/ubt-go"
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v2/client"
	"github.com/alibabacloud-go/tea/tea"
	. "github.com/space-tech-dev/sms-service-notification/config"
	logger "github.com/space-tech-dev/sms-service-notification/loggerManager"
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

func (srv *AliYunService) SendMsg(mobile string, sign string, templateId string, templateParam string) (_err error) {

	sendSmsRequest := &dysmsapi20170525.SendSmsRequest{
		PhoneNumbers: tea.String(mobile),
		SignName: tea.String(sign),
		TemplateCode: tea.String(templateId),
		TemplateParam: tea.String(templateParam),
	}
	// 复制代码运行请自行打印 API 的返回值
	response, _err := AliYunClient.SendSms(sendSmsRequest)
	if _err != nil {
		fmt.Println(_err)
	}
	logger.UBTHandler.Info("request sms", &UBT.ExtraMessage{
	Module: "request.sms",
	BusinessInfo: map[string]string{
		"data": response.String(),
		"request_id": *response.Body.RequestId,
		"code": *response.Body.Code,
		"message": *response.Body.Message,
	}})
	fmt.Println(response.Body)
	return _err
}

