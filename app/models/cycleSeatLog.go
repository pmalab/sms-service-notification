package models

import (
	"github.com/space-tech-dev/sms-service-notification/config"
	"github.com/golang-module/carbon"
)

// TableName overrides the table name used by User to `profiles`
func (mdl *CycleSeatLog) TableName() string {
	return mdl.GetTableName(true)
}

type CycleSeatLog struct {
	ID                   int           `gorm:"column:id" json:"id"`
	UUID                 string        `gorm:"column:integration_id__c" json:"uuid" primaryKey;autoIncrement:"false" sql:"index"`
	Account              string        `gorm:"column:account__c" json:"account__c"`
	AccountUUID          string        `gorm:"column:account__r__integration_id__c" json:"account__r__integration_id__c"`
	AvgPower             string        `gorm:"column:avg_power__c" json:"avgPower"`
	AvgRpm               string        `gorm:"column:avg_rpm__c" json:"avgRpm"`
	BookingStatus        string        `gorm:"column:booking_status__c" json:"bookingStatus"`
	Calories             string        `gorm:"column:calories__c" json:"calories"`
	CreatedDate          carbon.Carbon `gorm:"column:createddate" json:"createddate"`
	DeviceId             string        `gorm:"column:device_id__c" json:"deviceId"`
	Distance             string        `gorm:"column:distance__c" json:"distance"`
	HighPower            string        `gorm:"column:high_power__c" json:"highPower"`
	HighRpm              string        `gorm:"column:high_rpm__c" json:"highRpm"`
	IntegrationId        string        `gorm:"column:integration_id__c" json:"integrationId"`
	MigrationId          string        `gorm:"column:migration_id__c" json:"migrationId"`
	Name                 string        `gorm:"column:name" json:"name"`
	NeedToSyncFromHeroku string        `gorm:"column:need_to_sync_from_heroku__c" json:"needToSyncFromHeroku"`
	ReservationSFID      string        `gorm:"column:reservation__c" json:"reservationSFID"`
	ReservationUUID      string        `gorm:"column:reservation__r__integration_id__c" json:"reservationUUID"`
	Row                  string        `gorm:"column:row__c" json:"row"`
	RowLocation          string        `gorm:"column:row_location__c" json:"rowLocation"`
	ScheduleSFID         string        `gorm:"column:schedule__c" json:"scheduleSFID"`
	ScheduleUUID         string        `gorm:"column:schedule__r__integration_id__c" json:"scheduleUUID"`
	SeatNumber           string        `gorm:"column:seat_number__c" json:"seatNumber"`
	SystemModStamp       carbon.Carbon `gorm:"column:systemmodstamp" json:"systemmodstamp"`
	TotalEnergy          string        `gorm:"column:total_energy__c" json:"totalEnergy"`
}

const BOOK_STATUS_OCCUPIED = "Occupied"
const BOOK_STATUS_AVAILABLE = "Available"
const BOOK_STATUS_MAINTENANCE = "Under Maintenance"

var ARRAY_BOOK_STATUS = []string{

	BOOK_STATUS_OCCUPIED,
	BOOK_STATUS_AVAILABLE,
	BOOK_STATUS_MAINTENANCE,
}

const TABLE_NAME_CYCLE_SEAT_LOG = "cycle_seat_log__c"

// 模块初始化函数 import 包时被调用
func init() {
	//fmt.Println("reservation model module init function")
}

func NewCycleSeatLog() *CycleSeatLog {
	return &CycleSeatLog{}
}

func (mdl *CycleSeatLog) GetTableName(needFull bool) string {
	tableName := TABLE_NAME_CYCLE_SEAT_LOG
	if needFull {
		tableName = config.DatabaseConn.Schemas["option"] + "." + tableName
	}
	return tableName
}

func (mdl *CycleSeatLog) GetUUID() string {
	return mdl.UUID
}
