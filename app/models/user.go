package models

import (
	"database/sql"
	"github.com/space-tech-dev/sms-service-notification/config"
	"gorm.io/gorm"
)

// TableName overrides the table name used by User to `profiles`
func (mdl *User) TableName() string {
	//return config.DatabaseConn.Schemas["default"] + "." + "users"
	return mdl.GetTableName(true)
}

type User struct {
	Id                    int    `gorm:"json:id" column:"id"`
	UUID                  string `gorm:"column:uuid" json:"uuid" primaryKey;autoIncrement:"false" sql:"index"`
	LeadUUID              string `gorm:"column:lead_uuid" json:"leadUUID"`
	AccountUUID           string `gorm:"column:account_uuid" json:"accountUuid"`
	Mobile                string `gorm:"column:mobile" json:"mobile"`
	OfficialAccountOpenID string `gorm:"column:official_account_openid" json:"officialAccountOpenID"`
	Password              string `gorm:"column:password" json:"password"`
	Locale                string `gorm:"column:locale" json:"locale"`
}

const TABLE_NAME_USER = "users"

// 模块初始化函数 import 包时被调用
func init() {
	//fmt.Println("user model module init function")
}

func NewUser() *User {
	return &User{}
}

func (mdl *User) GetTableName(needFull bool) string {
	tableName := TABLE_NAME_USER
	if needFull {
		tableName = config.DatabaseConn.Schemas["default"] + "." + tableName
	}
	return tableName
}

/**
 * Scope Where Conditions
 */
func (user *User) WhereUserName(uuidOrPhone string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("uuid=@value OR mobile=@value", sql.Named("value", uuidOrPhone))
	}
}

func (user *User) WhereIsValid(db *gorm.DB) *gorm.DB {
	return db.Where("account_uuid != null")
}
