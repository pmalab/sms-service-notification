package models

import (
	"github.com/golang-module/carbon"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Model interface {
	GetTableName() string
	GetUUID() string
}

const STATUS_DRAFT = "Draft"
const STATUS_CANCELED = "Canceled"

const APPROVAL_STATUS_DRAFT = "Draft"
const APPROVAL_STATUS_PENDING = "Pending"
const APPROVAL_STATUS_APPROVED = "Approved"
const APPROVAL_STATUS_REJECTED = "Rejected"

const LOCALE_EN = "en_US"
const LOCALE_CN = "zh_CN"
const LOCALE_TW = "zh_TW"

var ARRAY_LOCALE = []string{
	LOCALE_EN,
	LOCALE_CN,
	LOCALE_TW,
}

var ARRAY_TIMEZONE = []string{
	LOCALE_EN,
	LOCALE_CN,
	LOCALE_TW,
}

type MyModel struct {
	ID   int    `gorm:"autoIncrement:true;unique; column:id" json:"-"`
	UUID string `gorm:"primaryKey;autoIncrement:false;unique; column:integration_id__c" json:"uuid" sql:"index"`

	CreatedAt carbon.Carbon `gorm:"column:createddate" json:"createddate"`
	UpdatedAt carbon.Carbon `gorm:"column:lastmodifieddate" json:"lastmodifieddate"`
}

func NewMyModel() *MyModel {
	now := carbon.Now()
	return &MyModel{
		UUID:      uuid.NewString(),
		CreatedAt: now,
		UpdatedAt: now,
	}
}

// 模块初始化函数 import 包时被调用
func init() {
	//fmt.Println("model model module init function")
}

/**
 * Scope Where Conditions
 */
func WhereUUID(uuid string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("integration_id__c=?", uuid)
	}
}

/**
 * Association Relationship
 */
func AssociationRelationship(db *gorm.DB, conditions *gorm.DB, mdl interface{}, relationship string) *gorm.Association {

	tx := db.Model(mdl)

	if conditions != nil {
		tx = tx.Where(conditions)
	}

	return tx.Association(relationship)
}
