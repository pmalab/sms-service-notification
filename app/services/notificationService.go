package service

import (
	_ "embed"
	"flag"
	fmt2 "fmt"
	. "github.com/space-tech-dev/sms-service-notification/app/models"
	config2 "github.com/space-tech-dev/sms-service-notification/config"
	"github.com/space-tech-dev/sms-service-notification/database"
	logger "github.com/space-tech-dev/sms-service-notification/loggerManager"
	"github.com/ArtisanCloud/go-libs/carbon"
	"github.com/ArtisanCloud/go-libs/str"
	UBT "github.com/ArtisanCloud/ubt-go"
	"github.com/google/uuid"
	"github.com/sideshow/apns2"
	"github.com/sideshow/apns2/certificate"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"log"
	"strconv"
)

type NotificationService struct {
	//NotificationLog *NotificationLog
}

/**
** 初始化构造函数
 */

// 模块初始化函数 import 包时被调用
func init() {
	//fmt.Println("notification service module init function")
}

func NewNotificationService() (notificationSrv *NotificationService) {
	notificationSrv = &NotificationService{
		//NotificationLog: NewNotificationLog(),
	}
	return notificationSrv
}

//
// 創建 notification log
//
func (srv *NotificationService) CreateNotificationLog(
	accountUUID string,
	deviceUUID string,
	platform string,
	title string,
	message string,
	sendStatus string,
) (notificationLog *NotificationLog) {

	notificationLog = &NotificationLog{
		UUID:        uuid.NewString(),
		AccountUUID: accountUUID,
		DeviceUUID:  deviceUUID,
		Platform:    platform,
		Title:       title,
		Message:     message,
		SendStatus:  sendStatus,
		SendDate:    carbon.GetCarbonNow(),
	}

	return notificationLog
}

//
// 保存 notification log
//
func (srv *NotificationService) SaveNotificationLog(notificationLog *NotificationLog) error {

	return database.DBConnection.Transaction(func(tx *gorm.DB) error {

		result := tx.Omit(clause.Associations).Create(notificationLog)
		if result.Error != nil {
			return result.Error
		}

		return nil
	})
}

// 載入檔案
//go:embed apns/apns.p12
var certs []byte

//
// 推送iOS訊息
//
func (srv *NotificationService) PushIOSMessage(notificationLog *NotificationLog, deviceToken string, message string, title string) {

	cert, err := certificate.FromP12Bytes(certs, "Space")
	if err != nil {
		log.Fatal("Cert Error:", err)
	}
	flag.Parse()

	notification := &apns2.Notification{}
	//notification.DeviceToken = *deviceToken
	notification.DeviceToken = deviceToken
	notification.Topic = "cn.space.Space-Cycle"
	notification.Payload = []byte(`{"aps":{"alert":{ "body": "` + message + `", "title": "` + title + `"}} ,"key1":"val2","key":"val" }`) // See Payload section below

	// If you want to test push notifications for builds running directly from XCode (Development), use
	// client := apns2.NewClient(cert).Development()
	// For apps published to the app store or installed as an ad-hoc distribution use Production()

	client := apns2.NewClient(cert).Production()

	if config2.AppConfigure != nil && str.Lower(config2.AppConfigure.Env) == "development" {
		client = apns2.NewClient(cert).Development()
	}

	res, err := client.Push(notification)

	// 記錄發送的Log
	logger.UBTHandler.Info("push message result", &UBT.ExtraMessage{
		Module: "push.message.result",
		BusinessInfo: map[string]string{
			"accountUUID": notificationLog.AccountUUID,
			"StatusCode":  strconv.Itoa(res.StatusCode),
			"ApnsID":      res.ApnsID,
			"Reason":      res.Reason,
		}})

	// 發送失敗，更新Notification Log發送狀態
	if res.StatusCode != 200 || err != nil {
		notificationLog.SendStatus = SEND_STATUS_FAILED
		notificationLog.SendFailReason = res.Reason
	}

	srv.SaveNotificationLog(notificationLog)

	if err != nil {
		log.Fatal("Error:", err)
	}

	//log.Printf("%v\n%v %v %v\n",deviceToken, res.StatusCode, res.ApnsID, res.Reason)
}

////
//// 推送iOS訊息 by account
////
//func (srv *NotificationService) PushIOSMessageByAccount(account *Account, title string, message string) {
//
//	if account.Devices == nil {
//		account.LoadDevices(database.DBConnection, nil)
//	}
//
//	if account.Devices != nil && len(account.Devices) > 0 {
//		for _, device := range account.Devices {
//			if device.Platform == PLATFORM_IOS && device.Status == STATUS_ACTIVE {
//				//fmt2.Dump(device.DeviceToken)
//				//var deviceToken = flag.String("dt"+strconv.Itoa(i), device.DeviceToken, "DeviceToken")
//				//var deviceToken = flag.String("dt", "a5f9a7610e0973b102ff4daf87e1e001d66e30fb39854e71db55332bd3adfbab", "DeviceToken")
//				var deviceToken = device.DeviceToken
//				//fmt2.Dump(deviceToken)
//
//				// 創建 Notification Log
//				notificationLog := srv.CreateNotificationLog(
//					device.AccountUUID,
//					device.UUID,
//					device.Platform,
//					title,
//					message,
//					SEND_STATUS_SUCCESS,
//				)
//
//				// 推送
//				srv.PushIOSMessage(notificationLog, deviceToken, message, title)
//			}
//		}
//	} else {
//		logger.UBTHandler.Info("account has no device", &UBT.ExtraMessage{
//			Module: "go.notification",
//			BusinessInfo: map[string]string{
//				"accountUUID":     account.UUID,
//			}})
//	}
//}

//
// 推送訊息-預約正取成功
//
func (srv *NotificationService) PushReservationSuccess(reservation *Reservation) {
	// 取得用戶的語言
	locale := reservation.GetUserLocale()

	studioName := reservation.Schedule.GetStudioNameByLocale(locale)
	classStartTime := reservation.Schedule.ClassStartTime.ToDateTimeString()
	instructorNames := reservation.Schedule.GetInstructorNamesByLocale(locale)
	programNames := reservation.Schedule.GetProgramNamesByLocale(locale)

	// 產生Printer 根據語言取得訊息
	var p *message.Printer
	if locale == SF_LANGUAGE_EN {
		p = message.NewPrinter(language.English)
	} else {
		p = message.NewPrinter(language.Chinese)
	}

	// 取得標題
	title := p.Sprintf(fmt2.Sprintf("%s", config2.MESSAGE_HUSHAN_RESERVATION_SUCCESS_TITLE))

	// 取得訊息
	message := p.Sprintf(fmt2.Sprintf("%s", config2.MESSAGE_HUSHAN_RESERVATION_SUCCESS), studioName, classStartTime, instructorNames, programNames)
	if reservation.ReservedCycleSeatLog != nil {
		message += ", " + p.Sprintf(fmt2.Sprintf("%s", config2.MESSAGE_CYCLE_SEAT_NUMBER), reservation.ReservedCycleSeatLog.SeatNumber)
	}
	//fmt.Dump(locale, classStartTime, message)

	if reservation.Account.Devices != nil && len(reservation.Account.Devices) > 0 {

		for _, device := range reservation.Account.Devices {
			// iOS 推送
			if device.Platform == PLATFORM_IOS && device.Status == STATUS_ACTIVE {
				//fmt2.Dump(device.DeviceToken)
				//var deviceToken = flag.String("dt"+strconv.Itoa(i), device.DeviceToken, "DeviceToken")
				//var deviceToken = flag.String("dt", "a5f9a7610e0973b102ff4daf87e1e001d66e30fb39854e71db55332bd3adfbab", "DeviceToken")
				var deviceToken = device.DeviceToken
				//fmt2.Dump(deviceToken)

				// 創建 Notification Log
				notificationLog := srv.CreateNotificationLog(
					device.AccountUUID,
					device.UUID,
					device.Platform,
					title,
					message,
					SEND_STATUS_SUCCESS,
				)

				// 推送
				srv.PushIOSMessage(notificationLog, deviceToken, message, title)
			}
		}
	} else {
		logger.UBTHandler.Info("account has no device", &UBT.ExtraMessage{
			Module: "go.notification.reservation.success",
			BusinessInfo: map[string]string{
				"reservationUUID": reservation.UUID,
				"accountUUID":     reservation.AccountUUID,
			}})
	}
}

//
// 推送訊息-預約正取轉候補
//
func (srv *NotificationService) PushReservationNormalToWait(reservation *Reservation) {
	// 取得用戶的語言
	locale := reservation.GetUserLocale()

	studioName := reservation.Schedule.GetStudioNameByLocale(locale)
	classStartTime := reservation.Schedule.ClassStartTime.ToDateTimeString()
	instructorNames := reservation.Schedule.GetInstructorNamesByLocale(locale)
	programNames := reservation.Schedule.GetProgramNamesByLocale(locale)

	// 產生Printer 根據語言取得訊息
	var p *message.Printer
	if locale == SF_LANGUAGE_EN {
		p = message.NewPrinter(language.English)
	} else {
		p = message.NewPrinter(language.Chinese)
	}

	// 取得標題
	title := p.Sprintf(fmt2.Sprintf("%s", config2.MESSAGE_HUSHAN_RESERVATION_NORMAL_TO_WAIT_TITLE))

	// 取得訊息
	message := p.Sprintf(fmt2.Sprintf("%s", config2.MESSAGE_HUSHAN_RESERVATION_NORMAL_TO_WAIT), studioName, classStartTime, instructorNames, programNames)
	//fmt2.Dump(locale, classStartTime, message)

	if reservation.Account.Devices != nil && len(reservation.Account.Devices) > 0 {

		for _, device := range reservation.Account.Devices {

			// iOS 推送
			if device.Platform == PLATFORM_IOS && device.Status == STATUS_ACTIVE {

				var deviceToken = device.DeviceToken

				// 創建 Notification Log
				notificationLog := srv.CreateNotificationLog(
					device.AccountUUID,
					device.UUID,
					device.Platform,
					title,
					message,
					SEND_STATUS_SUCCESS,
				)

				// 推送
				srv.PushIOSMessage(notificationLog, deviceToken, message, title)
			}
		}
	} else {
		logger.UBTHandler.Info("account has no device", &UBT.ExtraMessage{
			Module: "go.notification.reservation.normalToWait",
			BusinessInfo: map[string]string{
				"reservationUUID": reservation.UUID,
				"accountUUID":     reservation.AccountUUID,
			}})
	}
}

//
// 推送訊息-預約候補
//
func (srv *NotificationService) PushReservationWait(reservation *Reservation) {
	// 取得用戶的語言
	locale := reservation.GetUserLocale()

	studioName := reservation.Schedule.GetStudioNameByLocale(locale)
	classStartTime := reservation.Schedule.ClassStartTime.ToDateTimeString()
	instructorNames := reservation.Schedule.GetInstructorNamesByLocale(locale)
	programNames := reservation.Schedule.GetProgramNamesByLocale(locale)

	// 產生Printer 根據語言取得訊息
	var p *message.Printer
	if locale == SF_LANGUAGE_EN {
		p = message.NewPrinter(language.English)
	} else {
		p = message.NewPrinter(language.Chinese)
	}

	// 取得標題
	title := p.Sprintf(fmt2.Sprintf("%s", config2.MESSAGE_HUSHAN_RESERVATION_WAIT_TITLE))

	// 取得訊息
	message := p.Sprintf(fmt2.Sprintf("%s", config2.MESSAGE_HUSHAN_RESERVATION_WAIT), studioName, classStartTime, instructorNames, programNames)
	//fmt2.Dump(locale, classStartTime, message)

	if reservation.Account.Devices != nil && len(reservation.Account.Devices) > 0 {

		for _, device := range reservation.Account.Devices {

			// iOS 推送
			if device.Platform == PLATFORM_IOS && device.Status == STATUS_ACTIVE {

				var deviceToken = device.DeviceToken

				// 創建 Notification Log
				notificationLog := srv.CreateNotificationLog(
					device.AccountUUID,
					device.UUID,
					device.Platform,
					title,
					message,
					SEND_STATUS_SUCCESS,
				)

				// 推送
				srv.PushIOSMessage(notificationLog, deviceToken, message, title)
			}
		}
	} else {
		logger.UBTHandler.Info("account has no device", &UBT.ExtraMessage{
			Module: "go.notification.reservation.wait",
			BusinessInfo: map[string]string{
				"reservationUUID": reservation.UUID,
				"accountUUID":     reservation.AccountUUID,
			}})
	}
}

//
// 推送訊息-預約候補轉正席
//
func (srv *NotificationService) PushReservationWaitToNormal(reservation *Reservation) {
	// 取得用戶的語言
	locale := reservation.GetUserLocale()

	studioName := reservation.Schedule.GetStudioNameByLocale(locale)
	classStartTime := reservation.Schedule.ClassStartTime.ToDateTimeString()
	instructorNames := reservation.Schedule.GetInstructorNamesByLocale(locale)
	programNames := reservation.Schedule.GetProgramNamesByLocale(locale)

	// 產生Printer 根據語言取得訊息
	var p *message.Printer
	if locale == SF_LANGUAGE_EN {
		p = message.NewPrinter(language.English)
	} else {
		p = message.NewPrinter(language.Chinese)
	}

	// 取得標題
	title := p.Sprintf(fmt2.Sprintf("%s", config2.MESSAGE_HUSHAN_RESERVATION_WAIT_TO_NORMAL_TITLE))

	// 取得訊息
	message := p.Sprintf(fmt2.Sprintf("%s", config2.MESSAGE_HUSHAN_RESERVATION_WAIT_TO_NORMAL), studioName, classStartTime, instructorNames, programNames)
	if reservation.ReservedCycleSeatLog != nil {
		message += ", " + p.Sprintf(fmt2.Sprintf("%s", config2.MESSAGE_CYCLE_SEAT_NUMBER), reservation.ReservedCycleSeatLog.SeatNumber)
	}
	//fmt2.Dump(locale, classStartTime, message)

	if reservation.Account.Devices != nil && len(reservation.Account.Devices) > 0 {

		for _, device := range reservation.Account.Devices {

			// iOS 推送
			if device.Platform == PLATFORM_IOS && device.Status == STATUS_ACTIVE {

				var deviceToken = device.DeviceToken

				// 創建 Notification Log
				notificationLog := srv.CreateNotificationLog(
					device.AccountUUID,
					device.UUID,
					device.Platform,
					title,
					message,
					SEND_STATUS_SUCCESS,
				)

				// 推送
				srv.PushIOSMessage(notificationLog, deviceToken, message, title)
			}
		}
	} else {
		logger.UBTHandler.Info("account has no device", &UBT.ExtraMessage{
			Module: "go.notification.reservation.waitToNormal",
			BusinessInfo: map[string]string{
				"reservationUUID": reservation.UUID,
				"accountUUID":     reservation.AccountUUID,
			}})
	}
}

//
// 推送訊息-預約候補轉保留名額
//
func (srv *NotificationService) PushReservationWaitToReservedPlace(reservation *Reservation) {
	// 取得用戶的語言
	locale := reservation.GetUserLocale()

	studioName := reservation.Schedule.GetStudioNameByLocale(locale)
	classStartTime := reservation.Schedule.ClassStartTime.ToDateTimeString()
	instructorNames := reservation.Schedule.GetInstructorNamesByLocale(locale)
	programNames := reservation.Schedule.GetProgramNamesByLocale(locale)

	// 產生Printer 根據語言取得訊息
	var p *message.Printer
	if locale == SF_LANGUAGE_EN {
		p = message.NewPrinter(language.English)
	} else {
		p = message.NewPrinter(language.Chinese)
	}

	// 取得標題
	title := p.Sprintf(fmt2.Sprintf("%s", config2.MESSAGE_HUSHAN_RESERVATION_WAIT_TO_RESERVED_PLACE_TITLE))

	// 取得訊息
	message := p.Sprintf(fmt2.Sprintf("%s", config2.MESSAGE_HUSHAN_RESERVATION_WAIT_TO_RESERVED_PLACE), studioName, classStartTime, instructorNames, programNames)
	//if reservation.ReservedCycleSeatLog != nil {
	//	message += ", " + p.Sprintf(fmt2.Sprintf("%s", config2.MESSAGE_CYCLE_SEAT_NUMBER), reservation.ReservedCycleSeatLog.SeatNumber)
	//}
	//fmt2.Dump(locale, classStartTime, message)

	if reservation.Account.Devices != nil && len(reservation.Account.Devices) > 0 {

		for _, device := range reservation.Account.Devices {

			// iOS 推送
			if device.Platform == PLATFORM_IOS && device.Status == STATUS_ACTIVE {

				var deviceToken = device.DeviceToken

				// 創建 Notification Log
				notificationLog := srv.CreateNotificationLog(
					device.AccountUUID,
					device.UUID,
					device.Platform,
					title,
					message,
					SEND_STATUS_SUCCESS,
				)

				// 推送
				srv.PushIOSMessage(notificationLog, deviceToken, message, title)
			}
		}
	} else {
		logger.UBTHandler.Info("account has no device", &UBT.ExtraMessage{
			Module: "go.notification.reservation.waitToReserved",
			BusinessInfo: map[string]string{
				"reservationUUID": reservation.UUID,
				"accountUUID":     reservation.AccountUUID,
			}})
	}
}


//
// 推送訊息-預約取消
//
func (srv *NotificationService) PushReservationCancel(reservation *Reservation) {
	// 取得用戶的語言
	locale := reservation.GetUserLocale()

	studioName := reservation.Schedule.GetStudioNameByLocale(locale)
	classStartTime := reservation.Schedule.ClassStartTime.ToDateTimeString()
	instructorNames := reservation.Schedule.GetInstructorNamesByLocale(locale)
	programNames := reservation.Schedule.GetProgramNamesByLocale(locale)

	// 產生Printer 根據語言取得訊息
	var p *message.Printer
	if locale == SF_LANGUAGE_EN {
		p = message.NewPrinter(language.English)
	} else {
		p = message.NewPrinter(language.Chinese)
	}

	// 取得標題
	title := p.Sprintf(fmt2.Sprintf("%s", config2.MESSAGE_HUSHAN_RESERVATION_CANCEL_TITLE))

	// 取得訊息
	message := p.Sprintf(fmt2.Sprintf("%s", config2.MESSAGE_HUSHAN_RESERVATION_CANCEL), studioName, classStartTime, instructorNames, programNames)
	//fmt2.Dump(locale, classStartTime, message)

	if reservation.Account.Devices != nil && len(reservation.Account.Devices) > 0 {

		for _, device := range reservation.Account.Devices {

			// iOS 推送
			if device.Platform == PLATFORM_IOS && device.Status == STATUS_ACTIVE {

				var deviceToken = device.DeviceToken

				// 創建 Notification Log
				notificationLog := srv.CreateNotificationLog(
					device.AccountUUID,
					device.UUID,
					device.Platform,
					title,
					message,
					SEND_STATUS_SUCCESS,
				)

				// 推送
				srv.PushIOSMessage(notificationLog, deviceToken, message, title)
			}
		}
	} else {
		logger.UBTHandler.Info("account has no device", &UBT.ExtraMessage{
			Module: "go.notification.reservation.cancel",
			BusinessInfo: map[string]string{
				"reservationUUID": reservation.UUID,
				"accountUUID":     reservation.AccountUUID,
			}})
	}
}

//
// 推送訊息-預約取消
//
func (srv *NotificationService) PushReservationUpdateCycleSeat(reservation *Reservation, oldCycleSeatLog *CycleSeatLog) {
	// 取得用戶的語言
	locale := reservation.GetUserLocale()

	studioName := reservation.Schedule.GetStudioNameByLocale(locale)
	classStartTime := reservation.Schedule.ClassStartTime.ToDateTimeString()
	instructorNames := reservation.Schedule.GetInstructorNamesByLocale(locale)
	programNames := reservation.Schedule.GetProgramNamesByLocale(locale)

	// 產生Printer 根據語言取得訊息
	var p *message.Printer
	if locale == SF_LANGUAGE_EN {
		p = message.NewPrinter(language.English)
	} else {
		p = message.NewPrinter(language.Chinese)
	}

	// 取得標題
	title := p.Sprintf(fmt2.Sprintf("%s", config2.MESSAGE_HUSHAN_RESERVATION_UPDATE_CYCLE_SEAT_TITLE))

	// 取得訊息
	message := p.Sprintf(fmt2.Sprintf("%s", config2.MESSAGE_HUSHAN_RESERVATION_UPDATE_CYCLE_SEAT), studioName, classStartTime, instructorNames, programNames)
	//fmt2.Dump(locale, classStartTime, message)
	message += ", " + p.Sprintf(fmt2.Sprintf("%s", config2.MESSAGE_CYCLE_SEAT_NUMBER), reservation.ReservedCycleSeatLog.SeatNumber)

	if reservation.Account.Devices != nil && len(reservation.Account.Devices) > 0 {

		for _, device := range reservation.Account.Devices {

			// iOS 推送
			if device.Platform == PLATFORM_IOS && device.Status == STATUS_ACTIVE {

				var deviceToken = device.DeviceToken

				// 創建 Notification Log
				notificationLog := srv.CreateNotificationLog(
					device.AccountUUID,
					device.UUID,
					device.Platform,
					title,
					message,
					SEND_STATUS_SUCCESS,
				)

				// 推送
				srv.PushIOSMessage(notificationLog, deviceToken, message, title)
			}
		}
	} else {
		logger.UBTHandler.Info("account has no device", &UBT.ExtraMessage{
			Module: "go.notification.reservation.cancel",
			BusinessInfo: map[string]string{
				"reservationUUID":     reservation.UUID,
				"oldCycleSeatLogUUID": oldCycleSeatLog.UUID,
				"newCycleSeatLogUUID": reservation.ReservedCycleNumberUUID,
				"accountUUID":         reservation.AccountUUID,
			}})
	}
}
