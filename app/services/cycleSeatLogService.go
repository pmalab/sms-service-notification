package service

import (
	. "github.com/space-tech-dev/sms-service-notification/app/models"
	"github.com/space-tech-dev/sms-service-notification/database"
)

type CycleSeatLogService struct {
	CycleSeatLog *CycleSeatLog
}

/**
 ** 初始化构造函数
 */

// 模块初始化函数 import 包时被调用
func init() {
	//fmt.Println("reservation service module init function")
}

func NewCycleSeatLogService() (cycleSeatLogService *CycleSeatLogService) {
	cycleSeatLogService = &CycleSeatLogService{
		CycleSeatLog: NewCycleSeatLog(),
	}
	return cycleSeatLogService
}

/**
 ** 实例函数
 */

func (srv *CycleSeatLogService) GetCycleSeatLog(uuid string) (cycleSeatLog *CycleSeatLog) {

	cycleSeatLog = &CycleSeatLog{}

	db := database.DBConnection.Scopes(
		WhereUUID(uuid),
	)

	result := db.Find(cycleSeatLog)

	if result.RowsAffected > 0 {
		//fmt.Printf("user: %v", user.Account)
		return cycleSeatLog

	} else {
		return nil
	}

}
