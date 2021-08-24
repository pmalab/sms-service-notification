package models

import (
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
