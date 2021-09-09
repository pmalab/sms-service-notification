package models

import (
	"github.com/golang-module/carbon"
	"github.com/space-tech-dev/sms-service-notification/config"
)

// TableName overrides the table name used by Account to `profiles`
func (mdl *Device) TableName() string {
	return mdl.GetTableName(true)
}

type Device struct {
	ID               int           `gorm:"column:id" json:"id"`
	UUID             string        `gorm:"column:integration_id__c" json:"uuid" primaryKey;autoIncrement:"false" sql:"index"`
	AccountSFID      string        `gorm:"column:account__c" json:"accountSFID"`
	AccountUUID      string        `gorm:"column:account__r__integration_id__c" json:"accountUUID"`
	CreatedDate      carbon.Carbon `gorm:"column:createddate" json:"createdDate"`
	DeviceName       string        `gorm:"column:device_name__c" json:"deviceName"`
	DeviceToken      string        `gorm:"column:device_token__c" json:"deviceToken"`
	IntegrationId    string        `gorm:"column:integration_id__c" json:"integrationId"`
	LastModifiedDate carbon.Carbon `gorm:"column:lastmodifieddate" json:"lastModifiedDate"`
	Name             string        `gorm:"column:name" json:"name"`
	Platform         string        `gorm:"column:platform__c" json:"platform"`
	Sfid             string        `gorm:"column:sfid" json:"sfid"`
	Status           string        `gorm:"column:status__c" json:"status"`
}

const TABLE_NAME_DEVICE = "device__c"

//const PLATFORM_IOS = "iOS"
//const PLATFORM_ANDROID = "Android"

// 模块初始化函数 import 包时被调用
func init() {
	//fmt.Println("user model module init function")
}

func NewDevice() *Device {
	return &Device{}
}

func (mdl *Device) GetTableName(needFull bool) string {
	tableName := TABLE_NAME_DEVICE
	if needFull {
		tableName = config.DatabaseConn.Schemas["option"] + "." + tableName
	}
	return tableName
}

func (mdl *Device) GetUUID() string {
	return mdl.UUID
}

/**
 *  Relationships
 */

/**
 *  Conditions
 */
