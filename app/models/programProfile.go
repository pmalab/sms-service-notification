package models

import (
	"github.com/space-tech-dev/sms-service-notification/config"
	"github.com/golang-module/carbon"
)

// TableName overrides the table name used by User to `profiles`
func (mdl *ProgramProfile) TableName() string {
	return mdl.GetTableName(true)
}

type ProgramProfile struct {
	ID               int           `gorm:"column:id" json:"id"`
	UUID             string        `gorm:"column:integration_id__c" json:"uuid" primaryKey;autoIncrement:"false" sql:"index"`
	CreatedDate      carbon.Carbon `gorm:"column:createddate" json:"createdDate"`
	Description      string        `gorm:"column:description__c" json:"description"`
	IntegrationId    string        `gorm:"column:integration_id__c" json:"integrationId"`
	Language         string        `gorm:"column:language__c" json:"language"`
	LastModifiedDate carbon.Carbon `gorm:"column:lastmodifieddate" json:"lastModifiedDate"`
	Name             string        `gorm:"column:name" json:"name"`
	ProgramSFID      string        `gorm:"column:program__c" json:"programSFID"`
	ProgramUUID      string        `gorm:"column:program__r__integration_id__c" json:"programUUID"`
	Sfid             string        `gorm:"column:sfid" json:"sfid"`
}

const TABLE_NAME_PROGRAM_PROFILE = "program_profile__c"

// 模块初始化函数 import 包时被调用
func init() {
	//fmt.Println("program profile model module init function")
}

func NewProgramProfile() *ProgramProfile {
	return &ProgramProfile{}
}

func (mdl *ProgramProfile) GetTableName(needFull bool) string {
	tableName := TABLE_NAME_PROGRAM_PROFILE
	if needFull {
		tableName = config.DatabaseConn.Schemas["option"] + "." + tableName
	}
	return tableName
}
