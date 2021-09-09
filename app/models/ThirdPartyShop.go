package models

import (
	"github.com/space-tech-dev/sms-service-notification/config"
	"gorm.io/gorm"
)

// TableName overrides the table name used by User to `profiles`
func (mdl *ThirdPartyShop) TableName() string {
	return mdl.GetTableName(true)
}

type ThirdPartyShop struct {
	Studio *Studio `gorm:"foreignKey:StudioUUID;references:UUID" json:"studio"`

	ID                 int     `gorm:"column:id" json:"id"`
	//UUID               string  `gorm:"column:integration_id__c" json:"uuid" primaryKey;autoIncrement:"false" sql:"index"`
	OpenShopUUID       string  `gorm:"column:open_shop_uuid__c" json:"openShopUUID"`
	ShopAddress        string  `gorm:"column:shop_address__c" json:"shopAddress"`
	SfId      		   string  `gorm:"column:sfid" json:"sfId"`
	ThirdPartyType     string  `gorm:"column:Third_Party_Type__c" json:"thirdPartyType"`
	Name               string  `gorm:"column:name" json:"name"`
	StudioSfId            string  `gorm:"column:studio__c" json:"studioSfId"`
	StudioUUID            string  `gorm:"column:studio__r__integration_id__c" json:"studioUUID"`
}

const TABLE_NAME_THIRD_PARTY_SHOP = "third_party_shop__c"

// 模块初始化函数 import 包时被调用
func init() {
	//fmt.Println("product model module init function")
}

func NewThirdPartyShop() *ThirdPartyShop {
	return &ThirdPartyShop{}
}

func (mdl *ThirdPartyShop) GetTableName(needFull bool) string {
	tableName := TABLE_NAME_THIRD_PARTY_SHOP
	if needFull {
		tableName = config.DatabaseConn.Schemas["option"] + "." + tableName
	}
	return tableName
}

func (mdl *ThirdPartyShop) GetNameByLocale(locale string) string {
	name := mdl.Name
	return name
}

func (mdl *ThirdPartyShop) WhereIsValid(uuid string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("studio__r__integration_id__c !=?", uuid)
	}
}