package service

import (
	. "github.com/space-tech-dev/sms-service-notification/app/models"
	"github.com/space-tech-dev/sms-service-notification/database"
)

type ThirdPartyService struct {
	ThirdPartyShop *ThirdPartyShop
}

/**
 ** 初始化构造函数
 */

// 模块初始化函数 import 包时被调用
func init() {
	//fmt.Println("reservation service module init function")
}

func NewThirdPartyService() (thirdPartyService *ThirdPartyService) {
	thirdPartyService = &ThirdPartyService{
		ThirdPartyShop: NewThirdPartyShop(),
	}
	return thirdPartyService
}

/**
 ** 实例函数
 */

func (srv *ThirdPartyService) GetThirdPartyShops(uuid string) (thirdPartyShops  []*ThirdPartyShop,  err error) {

	  thirdPartyShops = []*ThirdPartyShop{}

	db := database.DBConnection.Scopes(
		srv.ThirdPartyShop.WhereIsValid(uuid),
	)
	db.Preload("Studio")

	result := db.Find(&thirdPartyShops)
	err = result.Error
	if err != nil {
		return nil, err
	}

	return thirdPartyShops, nil

}