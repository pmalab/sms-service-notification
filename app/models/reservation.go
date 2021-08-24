package models

import (
	"github.com/space-tech-dev/sms-service-notification/config"
	"github.com/golang-module/carbon"
)

// TableName overrides the table name used by User to `profiles`
func (mdl *Reservation) TableName() string {
	return mdl.GetTableName(true)
}

type Reservation struct {
	//Membership      *Membership       `gorm:"foreignKey:MembershipUUID;references:UUID" json:"membership"`
	//ReservationLogs []*ReservationLog `gorm:"foreignKey:ReservationUUID;references:UUID" json:"reservationLogs"`
	Account              *Account      `gorm:"foreignKey:AccountUUID;references:UUID" json:"account"`
	User                 *User         `gorm:"foreignKey:AccountUUID;references:AccountUUID" json:"user"`
	Schedule             *Schedule     `gorm:"foreignKey:ScheduleUUID;references:UUID" json:"schedule"`
	ReservedCycleSeatLog *CycleSeatLog `gorm:"foreignKey:ReservedCycleNumberUUID;references:UUID" json:"reservedCycleSeatLog"`

	ID                      int           `gorm:"column:id" json:"id"`
	UUID                    string        `gorm:"column:integration_id__c" json:"uuid" primaryKey;autoIncrement:"false" sql:"index"`
	AccountSFID             string        `gorm:"column:account__c" json:"accountSFID"`
	AccountUUID             string        `gorm:"column:account__r__integration_id__c" json:"accountUUID"`
	CancelTime              carbon.Carbon `gorm:"column:cancel_time__c" json:"cancelTime"`
	CheckinTime             carbon.Carbon `gorm:"column:checkin_time__c" json:"checkinTime"`
	ConsumeMembershipSFID   string        `gorm:"column:consume_membership__c" json:"consumeMembershipSFID"`
	ConsumeMembershipUUID   string        `gorm:"column:consume_membership__r__integration_id__c" json:"consumeMembershipUUID"`
	ConsumedClasses         float64       `gorm:"column:consumed_classes__c" json:"consumedClasses"`
	ConsumedDate            carbon.Carbon `gorm:"column:consumed_date__c" json:"consumedDate"`
	CreatedDate             carbon.Carbon `gorm:"column:createddate" json:"createddate"`
	Description             string        `gorm:"column:description__c" json:"description"`
	MembershipSFID          string        `gorm:"column:membership__c" json:"membershipSFID"`
	MembershipUUID          string        `gorm:"column:membership__r__integration_id__c" json:"membershipUUID"`
	Name                    string        `gorm:"column:name" json:"name"`
	OperationStatus         string        `gorm:"column:operation_status__c" json:"operationStatus"`
	RecordTypeID            string        `gorm:"column:recordtypeid" json:"recordTypeID"`
	RequestCycleNumberSFID  string        `gorm:"column:request_cycle_number__c" json:"requestCycleNumberSFID"`
	RequestCycleNumberUUID  string        `gorm:"column:request_cycle_number__r__integration_id__c" json:"requestCycleNumberUUID"`
	RequestType             string        `gorm:"column:request_type__c" json:"requestType"`
	ReservationStatus       string        `gorm:"column:reservation_status__c" json:"reservationStatus"`
	ReservedCycleNumberSFID string        `gorm:"column:reserved_cycle_number__c" json:"reservedCycleNumberSFID"`
	ReservedCycleNumberUUID string        `gorm:"column:reserved_cycle_number__r__integration_id__c" json:"reservedCycleNumberUUID"`
	ScheduleSFID            string        `gorm:"column:schedule__c" json:"ScheduleSFID"`
	ScheduleUUID            string        `gorm:"column:schedule__r__integration_id__c" json:"scheduleUUID"`
	SearchClassName         string        `gorm:"column:search_class_name__c" json:"searchClassName"`
	Status                  string        `gorm:"column:status__c" json:"status"`
	SystemModStamp          carbon.Carbon `gorm:"column:systemmodstamp" json:"updatedDate"`
	WaitingListNo           float64       `gorm:"column:waiting_list_no__c" json:"waitingListNo"`
	WaitingListOrder        float64       `gorm:"column:waiting_list_order__c" json:"waitingListOrder"`
}

const TABLE_NAME_RESERVATION = "reservation__c"
const OBJECT_NAME_RESERVATION = "Reservation__c"

const RESERVATION_RECORD_TYPE_NORMAL = "Normal"
const RESERVATION_RECORD_TYPE_WAITING_LIST = "Waiting_List"

var ARRAY_RESERVATION_RECORD_TYPE = []string{
	RESERVATION_RECORD_TYPE_NORMAL,
	RESERVATION_RECORD_TYPE_WAITING_LIST,
}

const REQUEST_TYPE_CONFIRM = "Confirm"
const REQUEST_TYPE_WAIT_LIST = "Wait List"

const RESERVATION_STATUS_DRAFT = "Draft"
const RESERVATION_STATUS_CONFIRMED = "Confirmed"
const RESERVATION_STATUS_WAIT_LIST = "Wait List"
const RESERVATION_STATUS_FAILED = "Failed"

const OPERATION_STATUS_NONE = "None"
const OPERATION_STATUS_CANCELLING = "Cancelling"
const OPERATION_STATUS_CANCELLED = "Cancelled"
const OPERATION_STATUS_LATE_CANCELLED = "Late Cancelled"
const OPERATION_STATUS_NO_SHOW = "No Show"
const OPERATION_STATUS_CHECK_IN = "Check In"
const OPERATION_STATUS_CANCEL_FAILED = "Cancel Failed"
const OPERATION_STATUS_AUTO_CANCELLED = "Auto Cancelled"

// 模块初始化函数 import 包时被调用
func init() {
	//fmt.Println("reservation model module init function")
}

func NewReservation() *Reservation {
	return &Reservation{}
}

func (mdl *Reservation) GetTableName(needFull bool) string {
	tableName := TABLE_NAME_RESERVATION
	if needFull {
		tableName = config.DatabaseConn.Schemas["option"] + "." + tableName
	}
	return tableName
}

//// -- Schedule
//func (mdl *Reservation) LoadSchedule(db *gorm.DB, conditions *gorm.DB) (*Schedule, error) {
//	mdl.Schedule = &Schedule{}
//	err := AssociationRelationship(db, conditions, mdl, "Schedule").Find(&mdl.Schedule)
//	if err != nil {
//		panic(err)
//	}
//	return mdl.Schedule, err
//}

func (mdl *Reservation) GetUserLocale() string {
	locale := SF_LANGUAGE_TW
	if mdl.User != nil {
		locale = mdl.User.Locale
	}
	return locale
}
