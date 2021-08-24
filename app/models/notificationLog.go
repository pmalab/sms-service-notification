package models

import (
	"github.com/space-tech-dev/sms-service-notification/config"
	"github.com/golang-module/carbon"
)

// TableName overrides the table name used by User to `profiles`
func (mdl *NotificationLog) TableName() string {
	return mdl.GetTableName(true)
}

type NotificationLog struct {
	Account *Account `gorm:"foreignKey:AccountUUID;references:UUID" json:"account"`
	User    *User    `gorm:"foreignKey:AccountUUID;references:AccountUUID" json:"user"`

	ID               int           `gorm:"column:id" json:"id"`
	UUID             string        `gorm:"column:integration_id__c" json:"uuid" primaryKey;autoIncrement:"false" sql:"index"`
	AccountSFID      string        `gorm:"column:account__c" json:"accountSFID"`
	AccountUUID      string        `gorm:"column:account__r__integration_id__c" json:"accountUUID"`
	CreatedDate      carbon.Carbon `gorm:"column:createddate" json:"createdDate"`
	DeviceSFID       string        `gorm:"column:device__c" json:"deviceSFID"`
	DeviceUUID       string        `gorm:"column:device__r__integration_id__c" json:"deviceUUID"`
	IntegrationId    string        `gorm:"column:integration_id__c" json:"integrationId"`
	LastModifiedDate carbon.Carbon `gorm:"column:lastmodifieddate" json:"lastModifiedDate"`
	Message          string        `gorm:"column:message__c" json:"message"`
	Name             string        `gorm:"column:name" json:"name"`
	Platform         string        `gorm:"column:platform__c" json:"platform"`
	SendDate         carbon.Carbon `gorm:"column:send_date__c" json:"sendDate"`
	SendFailReason   string        `gorm:"column:send_fail_reason__c" json:"sendFailReason"`
	SendStatus       string        `gorm:"column:send_status__c" json:"sendStatus"`
	Sfid             string        `gorm:"column:sfid" json:"sfid"`
	Title            string        `gorm:"column:title__c" json:"title"`
}

const TABLE_NAME_NOTIFICATION_LOG = "notification_log__c"

const SEND_STATUS_DRAFT = "Draft"
const SEND_STATUS_SUCCESS = "Success"
const SEND_STATUS_FAILED = "Failed"

// 模块初始化函数 import 包时被调用
func init() {
	//fmt.Println("reservation model module init function")
}

func NewNotificationLog() *NotificationLog {
	return &NotificationLog{}
}

func (mdl *NotificationLog) GetTableName(needFull bool) string {
	tableName := TABLE_NAME_NOTIFICATION_LOG
	if needFull {
		tableName = config.DatabaseConn.Schemas["option"] + "." + tableName
	}
	return tableName
}
