package service

import (
	. "github.com/space-tech-dev/sms-service-notification/app/models"
	"github.com/space-tech-dev/sms-service-notification/database"
)

type ReservationService struct {
	Reservation *Reservation
}

/**
 ** 初始化构造函数
 */

// 模块初始化函数 import 包时被调用
func init() {
	//fmt.Println("reservation service module init function")
}

func NewReservationService() (reservationService *ReservationService) {
	reservationService = &ReservationService{
		Reservation: NewReservation(),
	}
	return reservationService
}

/**
 ** 实例函数
 */

func (srv *ReservationService) GetReservation(uuid string) (reservation *Reservation) {

	reservation = &Reservation{}

	db := database.DBConnection.Scopes(
		WhereUUID(uuid),
	)
	db.Preload("Account.Devices")
	db.Preload("User")
	db.Preload("Schedule.Studio.Profiles")
	//db.Preload("Schedule.SpaceClass.Programs.Profiles")
	db.Preload("Schedule.Programs.Profiles")
	db.Preload("Schedule.Instructors.Profiles")
	db.Preload("ReservedCycleSeatLog")

	result := db.Find(reservation)

	if result.RowsAffected > 0 {
		//fmt.Printf("user: %v", user.Account)
		return reservation

	} else {
		return nil
	}

}
