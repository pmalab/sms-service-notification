package models

type Salesforce struct {
	ID   int    `gorm:"column:id" json:"id"`
	UUID string `gorm:"column:integration_id__c" json:"uuid" primaryKey;autoIncrement:"false" sql:"index"`
}

const STATUS_ACTIVE = "Active"
const STATUS_INACTIVE = "Inactive"

const SF_LANGUAGE_CN = "zh_CN"
const SF_LANGUAGE_TW = "zh_TW"
const SF_LANGUAGE_EN = "en_US"

const COUNTRY_CODE_CHINA = "+86"
const COUNTRY_CODE_TW = "+886"
const SALESFORCE_TIMEZONE = "UTC"
const SALESFORCE_TIMEZONE_HOUR = 8

const _HC_LASTOP_SYNCED = "SYNCED"
const _HC_LASTOP_PENDING = "PENDING"
const _HC_LASTOP_UPDATED = "UPDATED"
const _HC_LASTOP_FAILED = "FAILED"


//const CACHE_SALESFORCE_TIME_OUT = 24 * 60 * time.Second

// 模块初始化函数 import 包时被调用
func init() {
	//fmt.Println("model module init function")
}

func NewSalesforce() *Salesforce {
	return &Salesforce{}
}

