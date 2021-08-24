package models

import (
	"github.com/space-tech-dev/sms-service-notification/config"
	"github.com/golang-module/carbon"
)

// TableName overrides the table name used by User to `profiles`
func (mdl *Instructor) TableName() string {
	return mdl.GetTableName(true)
}

type Instructor struct {
	Profiles []*InstructorProfile `gorm:"foreignKey:InstructorUUID;references:UUID" json:"profiles"`

	ID                  int           `gorm:"column:id" json:"id"`
	UUID                string        `gorm:"column:integration_id__c" json:"uuid" primaryKey;autoIncrement:"false" sql:"index"`
	Birthday            carbon.Carbon `gorm:"column:birthday__c" json:"birthday"`
	ClassLabel          string        `gorm:"column:class_label__c" json:"classLabel"`
	ClothSize           string        `gorm:"column:cloth_size__c" json:"clothSize"`
	FacebookLink        string        `gorm:"column:facebook_link__c" json:"facebookLink"`
	Gender              string        `gorm:"column:gender__c" json:"gender"`
	Height              float64       `gorm:"column:height__c" json:"height"`
	HomePageImageLink   string        `gorm:"column:home_page_image_link__c" json:"homePageImageLink"`
	InstgramLink        string        `gorm:"column:instgram_link__c" json:"instgramLink"`
	IntegrationId       string        `gorm:"column:integration_id__c" json:"integrationId"`
	Location            string        `gorm:"column:location__c" json:"location"`
	Name                string        `gorm:"column:name" json:"name"`
	PreviewScheduleLink string        `gorm:"column:preview_schedule_link__c" json:"previewScheduleLink"`
	PublicAvatarLink    string        `gorm:"column:public_avatar_link__c" json:"publicAvatarLink"`
	ShoesSize           float64       `gorm:"column:shoes_size__c" json:"shoesSize"`
	Status              string        `gorm:"column:status__c" json:"status"`
	Weight              float64       `gorm:"column:weight__c" json:"weight"`
	//CountryCode         string `gorm:"column:country_code__c" json:"countryCode"`
	//Email               string `gorm:"column:email__c" json:"email"`
	//KkboxId             string `gorm:"column:kkbox_id__c" json:"kkboxId"`
	//LineId              string `gorm:"column:line_id__c" json:"lineId"`
	//MobilePhone         string `gorm:"column:mobilephone__c" json:"mobilePhone"`
}

const TABLE_NAME_INSTRUCTOR = "instructor__c"

// 模块初始化函数 import 包时被调用
func init() {
	//fmt.Println("product model module init function")
}

func NewInstructor() *Instructor {
	return &Instructor{}
}

func (mdl *Instructor) GetTableName(needFull bool) string {
	tableName := TABLE_NAME_INSTRUCTOR
	if needFull {
		tableName = config.DatabaseConn.Schemas["option"] + "." + tableName
	}
	return tableName
}

func (mdl *Instructor) GetUUID() string {
	return mdl.UUID
}

func (mdl *Instructor) GetNameByLocale(locale string) string {
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
