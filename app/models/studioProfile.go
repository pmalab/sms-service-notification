package models

import (
	"github.com/space-tech-dev/sms-service-notification/config"
	"github.com/golang-module/carbon"
)

// TableName overrides the table name used by User to `profiles`
func (mdl *StudioProfile) TableName() string {
	return mdl.GetTableName(true)
}

type StudioProfile struct {
	ID               int           `gorm:"column:id" json:"id"`
	UUID             string        `gorm:"column:integration_id__c" json:"uuid" primaryKey;autoIncrement:"false" sql:"index"`
	Address          string        `gorm:"column:address__c" json:"address"`
	BusGuidance      string        `gorm:"column:bus_guidance__c" json:"busGuidance"`
	BusinessHour     string        `gorm:"column:business_hour__c" json:"businessHour"`
	CreatedDate      carbon.Carbon `gorm:"column:createddate" json:"createdDate"`
	Description      string        `gorm:"column:description__c" json:"description"`
	DrivingGuidance  string        `gorm:"column:driving_guidance__c" json:"drivingGuidance"`
	IntegrationId    string        `gorm:"column:integration_id__c" json:"integrationId"`
	Language         string        `gorm:"column:language__c" json:"language"`
	LastModifiedDate carbon.Carbon `gorm:"column:lastmodifieddate" json:"lastModifiedDate"`
	MetroGuidance    string        `gorm:"column:metro_guidance__c" json:"metroGuidance"`
	Name             string        `gorm:"column:name" json:"name"`
	Sfid             string        `gorm:"column:sfid" json:"sfid"`
	StudioSFID       string        `gorm:"column:studio__c" json:"studioSFID"`
	StudioUUID       string        `gorm:"column:studio__r__integration_id__c" json:"studioUUID"`
}

const TABLE_NAME_STUDIO_PROFILE = "studio_profile__c"

// 模块初始化函数 import 包时被调用
func init() {
	//fmt.Println("program profile model module init function")
}

func NewStudioProfile() *StudioProfile {
	return &StudioProfile{}
}

func (mdl *StudioProfile) GetTableName(needFull bool) string {
	tableName := TABLE_NAME_STUDIO_PROFILE
	if needFull {
		tableName = config.DatabaseConn.Schemas["option"] + "." + tableName
	}
	return tableName
}
