package models

import "github.com/space-tech-dev/sms-service-notification/config"

// TableName overrides the table name used by User to `profiles`
func (mdl *Program) TableName() string {
	return mdl.GetTableName(true)
}

type Program struct {
	Profiles []*ProgramProfile `gorm:"foreignKey:ProgramUUID;references:UUID" json:"profiles"`

	ID            int    `gorm:"column:id" json:"id"`
	UUID          string `gorm:"column:integration_id__c" json:"uuid" primaryKey;autoIncrement:"false" sql:"index"`
	Calorie       string `gorm:"column:calorie__c" json:"calorie"`
	Cardio        string `gorm:"column:cardio__c" json:"cardio"`
	ClassLabel    string `gorm:"column:class_label__c" json:"classLabel"`
	CoreStability string `gorm:"column:core_stability__c" json:"coreStability"`
	Flexibility   string `gorm:"column:flexibility__c" json:"flexibility"`
	IntegrationId string `gorm:"column:integration_id__c" json:"integrationId"`
	MigrationId   string `gorm:"column:migration_id__c" json:"migrationId"`
	Modality      string `gorm:"column:modality__c" json:"modality"`
	ModalityUUID  string `gorm:"column:modality__r__integration_id__c" json:"modalityUUID"`
	MusicJourney  string `gorm:"column:music_journey__c" json:"musicJourney"`
	Name          string `gorm:"column:name" json:"name"`
	Restorative   string `gorm:"column:restorative__c" json:"restorative"`
	Status        string `gorm:"column:status__c" json:"status"`
	Strength      string `gorm:"column:strength__c" json:"strength"`
	Sweat         string `gorm:"column:sweat__c" json:"sweat"`
}

const TABLE_NAME_PROGRAM = "program__c"

// 模块初始化函数 import 包时被调用
func init() {
	//fmt.Println("program model module init function")
}

func NewProgram() *Program {
	return &Program{}
}

func (mdl *Program) GetTableName(needFull bool) string {
	tableName := TABLE_NAME_PROGRAM
	if needFull {
		tableName = config.DatabaseConn.Schemas["option"] + "." + tableName
	}
	return tableName
}

func (mdl *Program) GetNameByLocale(locale string) string {
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
