package models

import "github.com/space-tech-dev/sms-service-notification/config"

// TableName overrides the table name used by User to `profiles`
func (mdl *SpaceClass) TableName() string {
	return mdl.GetTableName(true)
}

type SpaceClass struct {
	Programs []*Program `gorm:"many2many:salesforce.j_program_to_class__c;foreignKey:UUID;joinForeignKey:class__r__integration_id__c;References:UUID;JoinReferences:program__r__integration_id__c" json:"programs"`

	ID              int     `gorm:"column:id" json:"id"`
	UUID            string  `gorm:"column:integration_id__c" json:"uuid" primaryKey;autoIncrement:"false" sql:"index"`
	ClassImage      string  `gorm:"column:class_image__c" json:"classImage" primaryKey;autoIncrement:"false" sql:"index"`
	ClassLabel      string  `gorm:"column:class_label__c" json:"classLabel"`
	ClassTheme      string  `gorm:"column:class_theme__c" json:"classTheme"`
	ConsumedClasses float64 `gorm:"column:consumed_classes__c" json:"consumedClasses"`
	CurrencyISOCode string  `gorm:"column:currencyisocode" json:"currencyisocode"`
	DurationMinutes string  `gorm:"column:duration_minutes__c" json:"durationMinutes"`
	IntegrationId   string  `gorm:"column:integration_id__c" json:"integrationId"`
	MigrationId     string  `gorm:"column:migration_id__c" json:"migrationId"`
	Name            string  `gorm:"column:name" json:"name"`
	OwnerID         string  `gorm:"column:ownerid" json:"ownerid"`
	RecordTypeID    string  `gorm:"column:recordtypeid" json:"recordtypeid"`
	Status          string  `gorm:"column:status__c" json:"status"`
}

const TABLE_NAME_SPACE_CLASS = "class__c"
const OBJECT_NAME_SPACE_CLASS = "Class__c"

// 模块初始化函数 import 包时被调用
func init() {
	//fmt.Println("spaceClass model module init function")
}

func NewSpaceClass() *SpaceClass {
	return &SpaceClass{}
}

func (mdl *SpaceClass) GetTableName(needFull bool) string {
	tableName := TABLE_NAME_SPACE_CLASS
	if needFull {
		tableName = config.DatabaseConn.Schemas["option"] + "." + tableName
	}
	return tableName

}

func (mdl *SpaceClass) GetUUID() string {
	return mdl.UUID
}

//func (mdl *SpaceClass) IsGroup() bool {
//	return mdl.RecordTypeID == GetCachedRecordTypeID(OBJECT_NAME_SPACE_CLASS, RECORD_TYPE_GROUP)
//}
//
//func (mdl *SpaceClass) IsCampaign() bool {
//	return mdl.RecordTypeID == GetCachedRecordTypeID(OBJECT_NAME_SPACE_CLASS, RECORD_TYPE_CAMPAIGN)
//}
//
//func (mdl *SpaceClass) IsAcademy() bool {
//	return mdl.RecordTypeID == GetCachedRecordTypeID(OBJECT_NAME_SPACE_CLASS, RECORD_TYPE_ACADEMY)
//}
//
//func (mdl *SpaceClass) IsSpecial() bool {
//	return mdl.IsCampaign() || mdl.IsAcademy()
//}
