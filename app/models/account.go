package models

import (
	"database/sql"
	"github.com/space-tech-dev/sms-service-notification/config"
	"gorm.io/gorm"
)

// TableName overrides the table name used by Account to `profiles`
func (mdl *Account) TableName() string {
	return mdl.GetTableName(true)
}

type Account struct {
	//User    *User     `gorm:"foreignKey:AccountUUID;references:UUID"`
	//Reservations      []*Reservation `gorm:"foreignKey:AccountUUID;references:UUID"`
	//Memberships       []*Membership  `gorm:"foreignKey:AccountUUID;references:UUID"`
	//SharedMemberships []*Membership  `gorm:"many2many:salesforce.membership_sharing__c;foreignKey:UUID;joinForeignKey:share_with__r__integration_id__c;References:UUID;JoinReferences:membership__r__integration_id__c"`
	//Identities []*PicklistOption `gorm:"-" json:"identities"`

	Devices []*Device `gorm:"foreignKey:AccountUUID;references:UUID"`

	ID                           int    `gorm:"column:id" json:"id"`
	UUID                         string `gorm:"column:integration_id__c" json:"uuid" primaryKey;autoIncrement:"false" sql:"index"`
	AccountNumber                string `gorm:"column:accountnumber" json:"accountNumber"`
	AccountSource                string `gorm:"column:accountsource" json:"accountSource"`
	Active                       bool   `gorm:"column:active__c" json:"active"`
	AgreedtoJoinShes             string `gorm:"column:agreedto_joinshes__pc" json:"agreedtoJoinShes"`
	AlwaysGoStudioC              string `gorm:"column:always_go_studio__c" json:"alwaysGoStudioC"`
	AlwaysGoStudioUUID           string `gorm:"column:always_go_studio__r__integration_id__c" json:"alwaysGoStudioUUID"`
	AnnualRevenue                string `gorm:"column:annualrevenue" json:"annualRevenue"`
	ApprovalStatus               string `gorm:"column:approval_status__c" json:"approvalStatus"`
	ApprovedBy                   string `gorm:"column:approved_by__c" json:"approvedBy"`
	BackIsActive                 string `gorm:"column:back_is_active__c" json:"backIsActive"`
	Billingcity                  string `gorm:"column:billingcity" json:"billingcity"`
	Billingcountry               string `gorm:"column:billingcountry" json:"billingcountry"`
	Billinggeocodeaccuracy       string `gorm:"column:billinggeocodeaccuracy" json:"billinggeocodeaccuracy"`
	Billinglatitude              string `gorm:"column:billinglatitude" json:"billinglatitude"`
	Billinglongitude             string `gorm:"column:billinglongitude" json:"billinglongitude"`
	Billingpostalcode            string `gorm:"column:billingpostalcode" json:"billingpostalcode"`
	Billingstate                 string `gorm:"column:billingstate" json:"billingstate"`
	Billingstreet                string `gorm:"column:billingstreet" json:"billingstreet"`
	BusinessAaccount             string `gorm:"column:business_account__c" json:"businessAaccount"`
	BusinessAccountUUID          string `gorm:"column:business_account__r__integration_id__c" json:"businessAccountUUID"`
	Channel                      string `gorm:"column:channel__pc" json:"channel"`
	CiRecommendId                string `gorm:"column:ci_recommend_id__c" json:"ciRecommendId"`
	ClothSize                    string `gorm:"column:cloth_size__pc" json:"clothSize"`
	ConversionDate               string `gorm:"column:conversion_date__c" json:"conversionDate"`
	CountryCode                  string `gorm:"column:country_code__pc" json:"countryCode"`
	CreatedbyIntegrationId       string `gorm:"column:createdby__integration_id__c" json:"createdbyIntegrationId"`
	Createdbyid                  string `gorm:"column:createdbyid" json:"createdbyid"`
	Createddate                  string `gorm:"column:createddate" json:"createddate"`
	Currencyisocode              string `gorm:"column:currencyisocode" json:"currencyisocode"`
	Description                  string `gorm:"column:description" json:"description"`
	EmergencyContactName         string `gorm:"column:emergency_contact_name__pc" json:"emergencyContactName"`
	EmergencyContactPhoneNumber  string `gorm:"column:emergency_contact_phone_number__pc" json:"emergencyContactPhoneNumber"`
	EmergencyContactRelationship string `gorm:"column:emergency_contact_relationship__pc" json:"emergencyContactRelationship"`
	EnglishName                  string `gorm:"column:english_name__c" json:"englishName"`
	Fax                          string `gorm:"column:fax" json:"fax"`
	Firstname                    string `gorm:"column:firstname" json:"firstname"`
	Gender                       string `gorm:"column:gender__pc" json:"gender"`
	HeadPortraitFileId           string `gorm:"column:head_portrait_file_id__c" json:"headPortraitFileId"`
	HeadPortraitPublicUrl        string `gorm:"column:head_portrait_public_url__c" json:"headPortraitPublicUrl"`
	HeadPortraitUrl              string `gorm:"column:head_portrait_url__c" json:"headPortraitUrl"`
	Height                       string `gorm:"column:height__pc" json:"height"`
	IdCardNo                     string `gorm:"column:id_card_no__c" json:"idCardNo"`
	IdentifiedBy                 string `gorm:"column:identified_by__c" json:"identifiedBy"`
	Identity                     string `gorm:"column:identity__c" json:"identity"`
	Industry                     string `gorm:"column:industry" json:"industry"`
	IntegrationId                string `gorm:"column:integration_id__c" json:"integrationId"`
	IsNewcomer                   string `gorm:"column:is_newcomer__c" json:"isNewcomer"`
	IsTt                         string `gorm:"column:is_tt__pc" json:"isTt"`
	Isdeleted                    string `gorm:"column:isdeleted" json:"isdeleted"`
	Ispersonaccount              string `gorm:"column:ispersonaccount" json:"ispersonaccount"`
	LastmodifiedbyIntegrationId  string `gorm:"column:lastmodifiedby__integration_id__c" json:"lastmodifiedbyIntegrationId"`
	Lastmodifiedbyid             string `gorm:"column:lastmodifiedbyid" json:"lastmodifiedbyid"`
	Lastmodifieddate             string `gorm:"column:lastmodifieddate" json:"lastmodifieddate"`
	Lastname                     string `gorm:"column:lastname" json:"lastname"`
	LinkableAccountName          string `gorm:"column:linkable_account_name__c" json:"linkableAccountName"`
	MasterrecordIntegrationId    string `gorm:"column:masterrecord__integration_id__c" json:"masterrecordIntegrationId"`
	Masterrecordid               string `gorm:"column:masterrecordid" json:"masterrecordid"`
	MemberId                     string `gorm:"column:member_id__c" json:"memberId"`
	MemberLevel                  string `gorm:"column:member_level__pc" json:"memberLevel"`
	MgmRecommendedBy             string `gorm:"column:mgm_recommended_by__c" json:"mgmRecommendedBy"`
	MgmRecommendedByUUID         string `gorm:"column:mgm_recommended_by__r__integration_id__c" json:"mgmRecommendedByUUID"`
	Middlename                   string `gorm:"column:middlename" json:"middlename"`
	MigrationId                  string `gorm:"column:migration_id__c" json:"migrationId"`
	Name                         string `gorm:"column:name" json:"name"`
	Nationality                  string `gorm:"column:nationality__pc" json:"nationality"`
	NickName                     string `gorm:"column:nick_name__pc" json:"nickName"`
	Numberofemployees            string `gorm:"column:numberofemployees" json:"numberofemployees"`
	Occupation                   string `gorm:"column:occupation__pc" json:"occupation"`
	OwneUUID                     string `gorm:"column:owner__integration_id__c" json:"owneUUID"`
	Ownerid                      string `gorm:"column:ownerid" json:"ownerid"`
	Ownership                    string `gorm:"column:ownership" json:"ownership"`
	ParentIntegrationId          string `gorm:"column:parent__integration_id__c" json:"parentIntegrationId"`
	Parentid                     string `gorm:"column:parentid" json:"parentid"`
	PassportNo                   string `gorm:"column:passport_no__c" json:"passportNo"`
	Personassistantname          string `gorm:"column:personassistantname" json:"personassistantname"`
	Personassistantphone         string `gorm:"column:personassistantphone" json:"personassistantphone"`
	Personbirthdate              string `gorm:"column:personbirthdate" json:"personbirthdate"`
	Personcontactid              string `gorm:"column:personcontactid" json:"personcontactid"`
	Persondepartment             string `gorm:"column:persondepartment" json:"persondepartment"`
	Persondonotcall              string `gorm:"column:persondonotcall" json:"persondonotcall"`
	Personemail                  string `gorm:"column:personemail" json:"personemail"`
	Personhasoptedoutofemail     string `gorm:"column:personhasoptedoutofemail" json:"personhasoptedoutofemail"`
	Personhasoptedoutoffax       string `gorm:"column:personhasoptedoutoffax" json:"personhasoptedoutoffax"`
	Personhomephone              string `gorm:"column:personhomephone" json:"personhomephone"`
	Personmailingcity            string `gorm:"column:personmailingcity" json:"personmailingcity"`
	Personmailingcountry         string `gorm:"column:personmailingcountry" json:"personmailingcountry"`
	Personmailingpostalcode      string `gorm:"column:personmailingpostalcode" json:"personmailingpostalcode"`
	Personmailingstate           string `gorm:"column:personmailingstate" json:"personmailingstate"`
	Personmailingstreet          string `gorm:"column:personmailingstreet" json:"personmailingstreet"`
	Personmobilephone            string `gorm:"column:personmobilephone" json:"personmobilephone"`
	Personotherphone             string `gorm:"column:personotherphone" json:"personotherphone"`
	Phone                        string `gorm:"column:phone" json:"phone"`
	RealName                     string `gorm:"column:real_name__c" json:"realName"`
	RecommendedBy                string `gorm:"column:recommended_by__pc" json:"recommendedBy"`
	RecommendedByPUUID           string `gorm:"column:recommended_by__pr__integration_id__c" json:"recommendedByPUUID"`
	RecordtypeID                 string `gorm:"column:recordtypeid" json:"recordtypeID"`
	RegisterFrom                 string `gorm:"column:register_from__pc" json:"registerFrom"`
	Salutation                   string `gorm:"column:salutation" json:"salutation"`
	Sfid                         string `gorm:"column:sfid" json:"sfid"`
	Shippingcity                 string `gorm:"column:shippingcity" json:"shippingcity"`
	Shippingcountry              string `gorm:"column:shippingcountry" json:"shippingcountry"`
	Shippinggeocodeaccuracy      string `gorm:"column:shippinggeocodeaccuracy" json:"shippinggeocodeaccuracy"`
	Shippinglatitude             string `gorm:"column:shippinglatitude" json:"shippinglatitude"`
	Shippinglongitude            string `gorm:"column:shippinglongitude" json:"shippinglongitude"`
	Shippingpostalcode           string `gorm:"column:shippingpostalcode" json:"shippingpostalcode"`
	Shippingstate                string `gorm:"column:shippingstate" json:"shippingstate"`
	Shippingstreet               string `gorm:"column:shippingstreet" json:"shippingstreet"`
	ShoesSize                    string `gorm:"column:shoes_size__pc" json:"shoesSize"`
	Sicdesc                      string `gorm:"column:sicdesc" json:"sicdesc"`
	StaffNo                      string `gorm:"column:staff_no__c" json:"staffNo"`
	StudioSFID                   string `gorm:"column:studio__c" json:"studioSFID"`
	StudioUUID                   string `gorm:"column:studio__r__integration_id__c" json:"studioUUID"`
	Suffix                       string `gorm:"column:suffix" json:"suffix"`
	Systemmodstamp               string `gorm:"column:systemmodstamp" json:"systemmodstamp"`
	Type                         string `gorm:"column:type" json:"type"`
	UnapproveReason              string `gorm:"column:unapprove_reason__c" json:"unapproveReason"`
	Website                      string `gorm:"column:website" json:"website"`
	WechatId                     string `gorm:"column:wechat_id__c" json:"wechatId"`
	Weight                       string `gorm:"column:weight__pc" json:"weight"`
	WhereDoYouKnowSpace          string `gorm:"column:where_do_you_know_space__pc" json:"whereDoYouKnowSpace"`
}

const TABLE_NAME_ACCOUNT = "account"

const TYPE_BUSINESS = "Business"
const TYPE_MEMBER = "Member"
const TYPE_TEACHER = "Teacher"

const CHANNEL_SPACE = "Space"
const CHANNEL_WEWORK = "WeWork"
const CHANNEL_MERCEDES = "Mercedes"

var ARRAY_ACCOUNT_TYPE = []string{
	TYPE_BUSINESS,
	TYPE_MEMBER,
	TYPE_TEACHER,
}

const RECORD_TYPE_BUSINESS_ACCOUNT = "Business_Account"

const RECORD_TYPE_PERSON_ACCOUNT = "Person_Account"

var ARRAY_ACCOUNT_RECORD_TYPE = []string{
	RECORD_TYPE_BUSINESS_ACCOUNT,
	RECORD_TYPE_PERSON_ACCOUNT,
}

const IDENTITY_MEMBER = "Member"
const IDENTITY_INVESTOR = "Investor"
const IDENTITY_PUBLIC_FIGURE = "Public Figure"
const IDENTITY_EMPLOYEE = "Employee"
const IDENTITY_EMPLOYEE_FAMILY = "Employee Family"
const IDENTITY_EMPLOYEE_LEED = "Leed"
const IDENTITY_ALL = "All"

var ARRAY_IDENTITY = []string{
	IDENTITY_MEMBER,
	IDENTITY_INVESTOR,
	IDENTITY_PUBLIC_FIGURE,
	IDENTITY_EMPLOYEE,
	IDENTITY_EMPLOYEE_FAMILY,
	IDENTITY_EMPLOYEE_LEED,
	IDENTITY_ALL,
}

const IDENTIFIED_BY_ID = "ID"
const IDENTIFIED_BY_PASSPORT = "Passport"

// 模块初始化函数 import 包时被调用
func init() {
	//fmt.Println("user model module init function")
}

func NewAccount() *Account {
	return &Account{}
}

func (mdl *Account) GetTableName(needFull bool) string {
	tableName := TABLE_NAME_ACCOUNT
	if needFull {
		tableName = config.DatabaseConn.Schemas["option"] + "." + tableName
	}
	return tableName
}

func (mdl *Account) GetUUID() string {
	return mdl.UUID
}

/**
 *  Relationships
 */

// -- Devices
func (mdl *Account) LoadDevices(db *gorm.DB, conditions *gorm.DB) ([]*Device, error) {
	devices := []*Device{}
	err := AssociationRelationship(db, conditions, mdl, "Devices").Find(devices)
	if err != nil {
		panic(err)
	}

	mdl.Devices = devices
	return devices, err
}

/**
 *  Conditions
 */
func (mdl *Account) IsActive() bool {
	return mdl.Active
}

/**
 * Scope Where Conditions
 */
func (mdl *Account) WhereAccountName(uuidOrPhone string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("integration_id__c=@value OR personmobilephone=@value", sql.Named("value", uuidOrPhone))
	}
}

func (mdl *Account) WherePersonMobilePhone(phone string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("personmobilephone=@value", sql.Named("value", phone))
	}
}

func (mdl *Account) WhereIsActive(db *gorm.DB) *gorm.DB {
	//return db.Where("status__c = ?", "active")
	return db.Where("active__c = ?", true)
}

//func (mdl *Account) GetUserLocale() string {
//	locale := SF_LANGUAGE_TW
//	if mdl.User != nil {
//		locale = mdl.User.Locale
//	}
//	return locale
//}
