package models

import "github.com/space-tech-dev/sms-service-notification/config"

// TableName overrides the table name used by User to `profiles`
func (mdl *Studio) TableName() string {
	return mdl.GetTableName(true)
}

const LOCATION_TAIPEI = "Taipei"

type Studio struct {
	Profiles []*StudioProfile `gorm:"foreignKey:StudioUUID;references:UUID" json:"profiles"`

	ID                 int     `gorm:"column:id" json:"id"`
	UUID               string  `gorm:"column:integration_id__c" json:"uuid" primaryKey;autoIncrement:"false" sql:"index"`
	BusinessDays       string  `gorm:"column:business_days__c" json:"businessDays"`
	Channel            string  `gorm:"column:channel__c" json:"channel"`
	IntegrationId      string  `gorm:"column:integration_id__c" json:"integrationId"`
	IsVirtual          string  `gorm:"column:is_virtual__c" json:"isVirtual"`
	Latitude           float64 `gorm:"column:latitude__c" json:"latitude"`
	Location           string  `gorm:"column:location__c" json:"location"`
	Longitude          float64 `gorm:"column:longitude__c" json:"longitude"`
	MainStudio         string  `gorm:"column:main_studio__c" json:"mainStudio"`
	MainStudioUUID     string  `gorm:"column:main_studio__r__integration_id__c" json:"mainStudioUUID"`
	MigrationId        string  `gorm:"column:migration_id__c" json:"migrationId"`
	Name               string  `gorm:"column:name" json:"name"`
	Ownerid            string  `gorm:"column:ownerid" json:"ownerid"`
	Phone              string  `gorm:"column:phone__c" json:"phone"`
	PictureUrl         string  `gorm:"column:picture_url__c" json:"pictureUrl"`
	RecordtypeID       string  `gorm:"column:recordtypeid" json:"recordtypeID"`
	Status             string  `gorm:"column:status__c" json:"status"`
	StudioAbbreviation string  `gorm:"column:studio_abbreviation__c" json:"studioAbbreviation"`
	StudioManager      string  `gorm:"column:studio_manager__c" json:"studioManager"`
	StudioManagerUUID  string  `gorm:"column:studio_manager__r__integration_id__c" json:"studioManagerUUID"`
}

const TABLE_NAME_STUDIO = "studio__c"

// 模块初始化函数 import 包时被调用
func init() {
	//fmt.Println("product model module init function")
}

func NewStudio() *Studio {
	return &Studio{}
}

func (mdl *Studio) GetTableName(needFull bool) string {
	tableName := TABLE_NAME_STUDIO
	if needFull {
		tableName = config.DatabaseConn.Schemas["option"] + "." + tableName
	}
	return tableName
}

func (mdl *Studio) GetUUID() string {
	return mdl.UUID
}

func (mdl *Studio) GetNameByLocale(locale string) string {
	name := mdl.Name
	if mdl.Profiles != nil {
		for _, profile := range mdl.Profiles {
			if profile.Language == locale {
				name = profile.Name
				break
			}
		}
	}
	return name
}
