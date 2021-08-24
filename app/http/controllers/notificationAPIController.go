package controllers

import (
	"github.com/space-tech-dev/sms-service-notification/app/http/request"
	. "github.com/space-tech-dev/sms-service-notification/app/services"
	"github.com/space-tech-dev/sms-service-notification/config"
	logger "github.com/space-tech-dev/sms-service-notification/loggerManager"
	UBT "github.com/ArtisanCloud/ubt-go"
	"github.com/gin-gonic/gin"
	//"github.com/ArtisanCloud/go-libs/fmt"
	//"net/http"
	//	"strconv"
)

type NotificationAPIController struct {
	*APIController
	ServiceNotification *NotificationService
	ServiceReservation  *ReservationService
	ServiceCycleSeatLog  *CycleSeatLogService
}

// 模块初始化函数 import 包时被调用
func init() {
	//fmt.Println("Init notification controller")
}

func NewNotificationAPIController(context *gin.Context) (ctl *NotificationAPIController) {
	return &NotificationAPIController{
		APIController:       NewAPIController(),
		ServiceNotification: NewNotificationService(),
		ServiceReservation:  NewReservationService(),
		ServiceCycleSeatLog:  NewCycleSeatLogService(),
	}
}

//
// 推送訊息-預約正取成功
//
func APIPushReservationSuccess(context *gin.Context) {
	ctl := NewNotificationAPIController(context)

	params, _ := context.Get("params")
	para := params.(request.ParaPushReservationSuccess)
	reservationUUID := para.ReservationUUID

	logger.UBTHandler.Info("push message request", &UBT.ExtraMessage{
		Module: "go.notification.reservation.success",
		BusinessInfo: map[string]string{
			"reservationUUID": reservationUUID,
		}})

	defer RecoverResponse(context)

	// 取得預約
	ctl.ServiceReservation.Reservation = ctl.ServiceReservation.GetReservation(reservationUUID)
	//fmt.Dump(ctl.ServiceReservation.Reservation)

	// 預約=null，報錯
	if ctl.ServiceReservation.Reservation == nil {
		ctl.RS.SetCode(config.API_ERR_CODE_NO_RESERVATION, config.API_RETURN_CODE_ERROR, "", "")
		panic(ctl.RS)
		return
	}

	//推送訊息
	ctl.ServiceNotification.PushReservationSuccess(ctl.ServiceReservation.Reservation)
	ctl.RS.Success(context, nil)

}

//
// 推送訊息-預約正取轉候補
//
func APIPushReservationNormalToWait(context *gin.Context) {
	ctl := NewNotificationAPIController(context)

	params, _ := context.Get("params")
	para := params.(request.ParaPushReservationNormalToWait)
	reservationUUID := para.ReservationUUID

	logger.UBTHandler.Info("push message request", &UBT.ExtraMessage{
		Module: "go.notification.reservation.normalToWait",
		BusinessInfo: map[string]string{
			"reservationUUID": reservationUUID,
		}})

	defer RecoverResponse(context)

	// 取得預約
	ctl.ServiceReservation.Reservation = ctl.ServiceReservation.GetReservation(reservationUUID)
	//fmt2.Dump(ctl.ServiceReservation.Reservation)

	// 預約=null，報錯
	if ctl.ServiceReservation.Reservation == nil {
		ctl.RS.SetCode(config.API_ERR_CODE_NO_RESERVATION, config.API_RETURN_CODE_ERROR, "", "")
		panic(ctl.RS)
		return
	}

	// 推送訊息
	ctl.ServiceNotification.PushReservationNormalToWait(ctl.ServiceReservation.Reservation)
	ctl.RS.Success(context, nil)
}

//
// 推送訊息-預約候補
//
func APIPushReservationWait(context *gin.Context) {
	ctl := NewNotificationAPIController(context)

	params, _ := context.Get("params")
	para := params.(request.ParaPushReservationWait)
	reservationUUID := para.ReservationUUID

	logger.UBTHandler.Info("push message request", &UBT.ExtraMessage{
		Module: "go.notification.reservation.wait",
		BusinessInfo: map[string]string{
			"reservationUUID": reservationUUID,
		}})

	defer RecoverResponse(context)

	// 取得預約
	ctl.ServiceReservation.Reservation = ctl.ServiceReservation.GetReservation(reservationUUID)
	//fmt2.Dump(ctl.ServiceReservation.Reservation)

	// 預約=null，報錯
	if ctl.ServiceReservation.Reservation == nil {
		ctl.RS.SetCode(config.API_ERR_CODE_NO_RESERVATION, config.API_RETURN_CODE_ERROR, "", "")
		panic(ctl.RS)
		return
	}

	// 推送訊息
	ctl.ServiceNotification.PushReservationWait(ctl.ServiceReservation.Reservation)
	ctl.RS.Success(context, nil)
}

//
// 推送訊息-預約候補轉正取
//
func APIPushReservationWaitToNormal(context *gin.Context) {
	ctl := NewNotificationAPIController(context)

	params, _ := context.Get("params")
	para := params.(request.ParaPushReservationWaitToNormal)
	reservationUUID := para.ReservationUUID

	logger.UBTHandler.Info("push message request", &UBT.ExtraMessage{
		Module: "go.notification.reservation.waitToNormal",
		BusinessInfo: map[string]string{
			"reservationUUID": reservationUUID,
		}})

	defer RecoverResponse(context)

	// 取得預約
	ctl.ServiceReservation.Reservation = ctl.ServiceReservation.GetReservation(reservationUUID)
	//fmt2.Dump(ctl.ServiceReservation.Reservation)

	// 預約=null，報錯
	if ctl.ServiceReservation.Reservation == nil {
		ctl.RS.SetCode(config.API_ERR_CODE_NO_RESERVATION, config.API_RETURN_CODE_ERROR, "", "")
		panic(ctl.RS)
		return
	}

	// 推送訊息
	ctl.ServiceNotification.PushReservationWaitToNormal(ctl.ServiceReservation.Reservation)
	ctl.RS.Success(context, nil)
}

//
// 推送訊息-預約候補轉保留名額
//
func APIPushReservationWaitToReservedPlace(context *gin.Context) {
	ctl := NewNotificationAPIController(context)

	params, _ := context.Get("params")
	para := params.(request.ParaPushReservationWaitToReservedPlace)
	reservationUUID := para.ReservationUUID

	logger.UBTHandler.Info("push message request", &UBT.ExtraMessage{
		Module: "go.notification.reservation.waitToReserved",
		BusinessInfo: map[string]string{
			"reservationUUID": reservationUUID,
		}})

	defer RecoverResponse(context)

	// 取得預約
	ctl.ServiceReservation.Reservation = ctl.ServiceReservation.GetReservation(reservationUUID)
	//fmt2.Dump(ctl.ServiceReservation.Reservation)

	// 預約=null，報錯
	if ctl.ServiceReservation.Reservation == nil {
		ctl.RS.SetCode(config.API_ERR_CODE_NO_RESERVATION, config.API_RETURN_CODE_ERROR, "", "")
		panic(ctl.RS)
		return
	}

	// 推送訊息
	ctl.ServiceNotification.PushReservationWaitToReservedPlace(ctl.ServiceReservation.Reservation)
	ctl.RS.Success(context, nil)
}


// 推送訊息-預約取消
func APIPushReservationCancel(context *gin.Context) {
	ctl := NewNotificationAPIController(context)

	params, _ := context.Get("params")
	para := params.(request.ParaPushReservationCancel)
	reservationUUID := para.ReservationUUID

	logger.UBTHandler.Info("push message request", &UBT.ExtraMessage{
		Module: "go.notification.reservation.cancel",
		BusinessInfo: map[string]string{
			"reservationUUID": reservationUUID,
		}})

	defer RecoverResponse(context)

	// 取得預約
	ctl.ServiceReservation.Reservation = ctl.ServiceReservation.GetReservation(reservationUUID)
	//fmt2.Dump(ctl.ServiceReservation.Reservation)

	// 預約=null，報錯
	if ctl.ServiceReservation.Reservation == nil {
		ctl.RS.SetCode(config.API_ERR_CODE_NO_RESERVATION, config.API_RETURN_CODE_ERROR, "", "")
		panic(ctl.RS)
		return
	}

	// 推送訊息
	ctl.ServiceNotification.PushReservationCancel(ctl.ServiceReservation.Reservation)
	ctl.RS.Success(context, nil)
}

// 推送訊息-更換車位
func APIPushReservationUpdateCycleSeat(context *gin.Context) {
	ctl := NewNotificationAPIController(context)

	params, _ := context.Get("params")
	para := params.(request.ParaPushReservationUpdateCycleSeat)
	reservationUUID := para.ReservationUUID
	oldCycleSeatLogUUID := para.OldCycleSeatLogUUID

	logger.UBTHandler.Info("push message request", &UBT.ExtraMessage{
		Module: "go.notification.reservation.update.cycle.eat",
		BusinessInfo: map[string]string{
			"reservationUUID": reservationUUID,
			"oldCycleSeatLogUUID": oldCycleSeatLogUUID,
		}})

	defer RecoverResponse(context)

	// 取得預約
	ctl.ServiceReservation.Reservation = ctl.ServiceReservation.GetReservation(reservationUUID)
	//fmt2.Dump(ctl.ServiceReservation.Reservation)
	ctl.ServiceCycleSeatLog.CycleSeatLog = ctl.ServiceCycleSeatLog.GetCycleSeatLog(oldCycleSeatLogUUID)
	//fmt2.Dump(ctl.ServiceCycleSeatLog.CycleSeatLog)

	// 預約=null，報錯
	if ctl.ServiceReservation.Reservation == nil {
		ctl.RS.SetCode(config.API_ERR_CODE_NO_RESERVATION, config.API_RETURN_CODE_ERROR, "", "")
		panic(ctl.RS)
		return
	}

	if ctl.ServiceCycleSeatLog.CycleSeatLog == nil {
		ctl.RS.SetCode(config.API_ERR_CODE_NO_CYCLE_SEAT_LOG, config.API_RETURN_CODE_ERROR, "", "")
		panic(ctl.RS)
		return
	}

	// 推送訊息
	ctl.ServiceNotification.PushReservationUpdateCycleSeat(ctl.ServiceReservation.Reservation, ctl.ServiceCycleSeatLog.CycleSeatLog)
	ctl.RS.Success(context, nil)
}
