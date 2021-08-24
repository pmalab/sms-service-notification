package models

import (
	"github.com/space-tech-dev/sms-service-notification/config"
	"github.com/golang-module/carbon"
)

// TableName overrides the table name used by User to `profiles`
func (mdl *InstructorProfile) TableName() string {
	return mdl.GetTableName(true)
}

type InstructorProfile struct {
	ID                   int           `gorm:"column:id" json:"id"`
	UUID                 string        `gorm:"column:integration_id__c" json:"uuid" primaryKey;autoIncrement:"false" sql:"index"`
	CreatedDate          carbon.Carbon `gorm:"column:createddate" json:"createdDate"`
	InstructorSFID       string        `gorm:"column:instructor__c" json:"instructorSFID"`
	InstructorUUID       string        `gorm:"column:instructor__r__integration_id__c" json:"instructorUUID"`
	IntegrationId        string        `gorm:"column:integration_id__c" json:"integrationId"`
	Intro                string        `gorm:"column:intro__c" json:"intro"`
	Language             string        `gorm:"column:language__c" json:"language"`
	LastModifiedDate     carbon.Carbon `gorm:"column:lastmodifieddate" json:"lastModifiedDate"`
	MailingCity          string        `gorm:"column:mailing_city__c" json:"mailingCity"`
	MailingCountry       string        `gorm:"column:mailing_country__c" json:"mailingCountry"`
	MailingStateProvince string        `gorm:"column:mailing_state_province__c" json:"mailingStateProvince"`
	MailingStreet        string        `gorm:"column:mailing_street__c" json:"mailingStreet"`
	MailingZipPostalCode string        `gorm:"column:mailing_zip_postal_code__c" json:"mailingZipPostalCode"`
	MaillingAddress      string        `gorm:"column:mailling_address__c" json:"maillingAddress"`
	MusicStyle           string        `gorm:"column:music_style__c" json:"musicStyle"`
	Name                 string        `gorm:"column:name" json:"name"`
	Sfid                 string        `gorm:"column:sfid" json:"sfid"`
	Slogan               string        `gorm:"column:slogan__c" json:"slogan"`
}

const TABLE_NAME_INSTRUCTOR_PROFILE = "instructor_profile__c"

// 模块初始化函数 import 包时被调用
func init() {
	//fmt.Println("program profile model module init function")
}

func NewInstructorProfile() *InstructorProfile {
	return &InstructorProfile{}
}

func (mdl *InstructorProfile) GetTableName(needFull bool) string {
	tableName := TABLE_NAME_INSTRUCTOR_PROFILE
	if needFull {
		tableName = config.DatabaseConn.Schemas["option"] + "." + tableName
	}
	return tableName
}
