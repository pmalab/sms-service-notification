package controllers

import (
	"fmt"
	. "github.com/space-tech-dev/sms-service-notification/app/http"
	"github.com/space-tech-dev/sms-service-notification/config"
	"github.com/gin-gonic/gin"
)

type APIController struct {
	//ServiceNotification *service.NotificationService
	RS          *APIResponse
}

func NewAPIController() *APIController {
	return &APIController{
		//ServiceNotification: service.NewNotificationService(),
		RS:          NewAPIResponse(),
	}
}



func RecoverResponse(context *gin.Context) {

	if p := recover(); p != nil {
		switch rs := p.(type) {

		// 获取业务流程中的异常错误
		case *APIResponse:
			rs.ThrowJSONResponse(context)

		// 若非APIResponse，也许默认抛出一个若非APIResponse
		default:
			fmt.Printf("Unknown panic: %v", p)

			defaultRS := &APIResponse{}
			defaultRS.SetReturnCode(config.API_RETURN_CODE_ERROR,"Inner Error")
			defaultRS.ThrowJSONResponse(context)

		}
	}
}
