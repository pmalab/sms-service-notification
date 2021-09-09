package controllers

import (
	"github.com/ArtisanCloud/go-libs/fmt"
	"github.com/gin-gonic/gin"
	. "github.com/space-tech-dev/sms-service-notification/app/services"
	"github.com/space-tech-dev/sms-service-notification/config"

	//"github.com/ArtisanCloud/go-libs/fmt"
	//"net/http"
	//	"strconv"
)

type ThirdPartyAPIController struct {
	*APIController
	ServiceThirdParty *ThirdPartyService
}

// 模块初始化函数 import 包时被调用
func init() {
	//fmt.Println("Init notification controller")
}

func NewThirdPartyAPIController(context *gin.Context) (ctl *ThirdPartyAPIController) {
	return &ThirdPartyAPIController{
		APIController:       NewAPIController(context),
		ServiceThirdParty: NewThirdPartyService(),
	}
}

//
// 获取所有第三方门店信息
//
func APIGetThirdPartyShops(context *gin.Context) {
	ctl := NewThirdPartyAPIController(context)

	defer RecoverResponse(context)

	// 取得所有门店信息
	thirdPartyShopList, err :=  ctl.ServiceThirdParty.GetThirdPartyShops("1")
	fmt.Dump(thirdPartyShopList)

	if err != nil {
		ctl.RS.SetCode(config.API_ERR_CODE_FAIL_TO_GET_THIRD_PARTY_SHOP_LIST, config.API_RETURN_CODE_ERROR, "", "")
		panic(ctl.RS)
		return
	}

	ctl.RS.Success(context, thirdPartyShopList)
}
