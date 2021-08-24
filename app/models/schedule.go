package models

import (
	//carbon2 "github.com/ArtisanCloud/go-libs/carbon"
	"github.com/space-tech-dev/sms-service-notification/config"
	"github.com/golang-module/carbon"
	//"gorm.io/gorm"
)

// TableName overrides the table name used by User to `profiles`
func (mdl *Schedule) TableName() string {
	return mdl.GetTableName(true)
}

type Schedule struct {
	//RecordType   *RecordType     `gorm:"foreignKey:RecordTypeID;references:SFID" json:"recordType"`
	//Academies    []*Academy      `gorm:"many2many:salesforce.j_academy_to_product_to_schedule__c;foreignKey:UUID;joinForeignKey:schedule__r__integration_id__c;References:UUID;JoinReferences:academy__r__integration_id__c" json:"academy"`
	//Campaigns    []*Campaign     `gorm:"many2many:salesforce.j_campaign_to_product_to_schedule__c;foreignKey:UUID;joinForeignKey:schedule__r__integration_id__c;References:UUID;JoinReferences:campaign__r__integration_id__c" json:"campaign"`
	//SeatsLogs    []*CycleSeatLog `gorm:"foreignKey:ScheduleUUID;references:UUID" json:"seatLogs"`
	Studio      *Studio       `gorm:"foreignKey:StudioUUID;references:UUID" json:"studio"`
	SpaceClass  *SpaceClass   `gorm:"foreignKey:ClassUUID;references:UUID" json:"class"`
	Programs    []*Program    `gorm:"many2many:salesforce.j_program_to_class__c;foreignKey:ClassUUID;joinForeignKey:class__r__integration_id__c;References:UUID;JoinReferences:program__r__integration_id__c" json:"programs"`
	Instructors []*Instructor `gorm:"many2many:salesforce.j_instructor_to_schedule__c;foreignKey:UUID;joinForeignKey:schedule__r__integration_id__c;References:UUID;JoinReferences:instructor__r__integration_id__c" json:"instructors"`

	ID                      int           `gorm:"column:id" json:"id"`
	UUID                    string        `gorm:"column:integration_id__c" json:"uuid" primaryKey;autoIncrement:"false" sql:"index"`
	ApprovalStatus          string        `gorm:"column:approval_status__c" json:"approvalStatus"`
	CanDisplayOnline        string        `gorm:"column:can_display_online__c" json:"canDisplayOnline"`
	Capacity                float64       `gorm:"column:capacity__c" json:"capacity"`
	ClassSFID               string        `gorm:"column:class__c" json:"classSFID"`
	ClassUUID               string        `gorm:"column:class__r__integration_id__c" json:"classUUID"`
	ClassEndTime            carbon.Carbon `gorm:"column:class_end_time__c" json:"classEndTime"`
	ClassLabel              string        `gorm:"column:class_label__c" json:"classLabel"`
	ClassStartTime          carbon.Carbon `gorm:"column:class_start_time__c" json:"classStartTime"`
	Classroom               string        `gorm:"column:classroom__c" json:"classroom"`
	ClassroomUUID           string        `gorm:"column:classroom__r__integration_id__c" json:"classroomUUID"`
	ConsumedClasses         float64       `gorm:"column:consumed_classes__c" json:"consumedClasses"`
	CopyFromSchedule        string        `gorm:"column:copy_from_schedule__c" json:"copyFromSchedule"`
	CopyFromScheduleUUID    string        `gorm:"column:copy_from_schedule__r__integration_id__c" json:"copyFromScheduleUUID"`
	Description             string        `gorm:"column:description__c" json:"description"`
	DisplayStatus           string        `gorm:"column:display_status__c" json:"displayStatus"`
	IntegrationId           string        `gorm:"column:integration_id__c" json:"integrationId"`
	MigrationId             string        `gorm:"column:migration_id__c" json:"migrationId"`
	Name                    string        `gorm:"column:name" json:"name"`
	OpenStatus              string        `gorm:"column:open_status__c" json:"openStatus"`
	RecordTypeID            string        `gorm:"column:recordtypeid" json:"recordTypeID"`
	ReservationMaximumQuota float64       `gorm:"column:reservation_maximum_quota__c" json:"reservationMaximumQuota"`
	Status                  string        `gorm:"column:status__c" json:"status"`
	StudioSFID              string        `gorm:"column:studio__c" json:"studioSFID"`
	StudioUUID              string        `gorm:"column:studio__r__integration_id__c" json:"studioUUID"`
	TrialMaximumQuota       float64       `gorm:"column:trial_maximum_quota__c" json:"trialMaximumQuota"`
	WaitListMaximumQuota    float64       `gorm:"column:waitlist_maximum_quota__c" json:"waitListMaximumQuota"`
}

const TABLE_NAME_SCHEDULE = "schedule__c"
const OBJECT_NAME_SCHEDULE = "Schedule__c"

const RECORD_TYPE_SCHEDULE_CYCLE = "Cycle"
const RECORD_TYPE_SCHEDULE_NORMAL = "Normal"

const DISPLAY_STATUS_DRAFT = "Draft"
const DISPLAY_STATUS_INSTRUCTOR = "Visible to Instructors"
const DISPLAY_STATUS_ALL = "Visible to All"

const OPEN_STATUS_RESERVING = "Reserving"
const OPEN_STATUS_STARTED = "Started"
const OPEN_STATUS_CLOSED = "Closed"

const STATUS_PUBLISH = "Published"

// 模块初始化函数 import 包时被调用
func init() {
	//fmt.Println("schedule model module init function")
}

func NewSchedule() *Schedule {
	return &Schedule{}
}

func (mdl *Schedule) GetTableName(needFull bool) string {
	tableName := TABLE_NAME_SCHEDULE
	if needFull {
		tableName = config.DatabaseConn.Schemas["option"] + "." + tableName
	}
	return tableName
}

func (mdl *Schedule) GetUUID() string {
	return mdl.UUID
}

//
//func (mdl *Schedule) IsFuture() bool {
//	//return carbon.Parse(mdl.ClassStartTime).Compare(">", GetCarbonNow())
//	return mdl.ClassStartTime.Compare(">", carbon2.GetCarbonNow())
//}
//
//func (mdl *Schedule) IsPass() bool {
//	//return carbon.Parse(mdl.ClassStartTime).Compare("<", GetCarbonNow())
//	return mdl.ClassStartTime.Compare("<", carbon2.GetCarbonNow())
//}
//
//func (mdl *Schedule) IsDisplayedToAll() bool {
//	return mdl.DisplayStatus == DISPLAY_STATUS_ALL
//}
//
//func (mdl *Schedule) IsPublished() bool {
//	return mdl.Status == STATUS_PUBLISH
//}
//
//func (mdl *Schedule) IsReserving() bool {
//	return mdl.OpenStatus == OPEN_STATUS_RESERVING
//}
//
//func (mdl *Schedule) IsApproved() bool {
//	return mdl.ApprovalStatus == APPROVAL_STATUS_APPROVED
//}
//
//func (mdl *Schedule) IsRecordTypeCycle() bool {
//	return mdl.RecordType.DeveloperName == RECORD_TYPE_SCHEDULE_CYCLE
//}
//
///**
// * Association Conditions
// */
//// Association("Products")
//func (mdl *Schedule) LoadProducts(db *gorm.DB, conditions *gorm.DB) ([]*Product, error) {
//	mdl.Products = []*Product{}
//	db = db.Table(TABLE_NAME_PRODUCT).
//		Select(
//			TABLE_NAME_PRODUCT + ".*" + "," +
//				TABLE_NAME_J_ACADEMY_TO_PRODUCT_TO_SCHEDULE + ".product__r__integration_id__c" + "," +
//				TABLE_NAME_J_CAMPAIGN_TO_PRODUCT_TO_SCHEDULE + ".product__r__integration_id__c",
//		).
//		Joins("left join " + TABLE_NAME_J_ACADEMY_TO_PRODUCT_TO_SCHEDULE + " on " +
//			TABLE_NAME_J_ACADEMY_TO_PRODUCT_TO_SCHEDULE + ".product__r__integration_id__c" +
//			"=" +
//			TABLE_NAME_PRODUCT + ".integration_id__c",
//		).
//		Joins("left join " + TABLE_NAME_J_CAMPAIGN_TO_PRODUCT_TO_SCHEDULE + " on " +
//			TABLE_NAME_J_CAMPAIGN_TO_PRODUCT_TO_SCHEDULE + ".product__r__integration_id__c" +
//			"=" +
//			TABLE_NAME_PRODUCT + ".integration_id__c",
//		)
//
//	if mdl.Academies != nil {
//		db = db.Where(TABLE_NAME_J_ACADEMY_TO_PRODUCT_TO_SCHEDULE+".schedule__r__integration_id__c=?", mdl.UUID)
//	}
//	if mdl.Campaigns != nil {
//		db = db.Where(TABLE_NAME_J_CAMPAIGN_TO_PRODUCT_TO_SCHEDULE+".schedule__r__integration_id__c=?", mdl.UUID)
//	}
//	err := db.Find(&mdl.Products).Error
//
//	return mdl.Products, err
//}
//
//// Association("Academies")
//func (mdl *Schedule) AssociationAcademies(db *gorm.DB, conditions *gorm.DB) ([]*Academy, error) {
//	academies := []*Academy{}
//	err := db.Model(mdl).
//		Where(conditions).
//		Association("Academies").
//		Find(academies)
//
//	if err != nil {
//		panic(err)
//	}
//	mdl.Academies = academies
//	return academies, err
//}
//
//// Association("Campaigns")
//func (mdl *Schedule) AssociationCampaigns(db *gorm.DB, conditions *gorm.DB) ([]*Campaign, error) {
//	campaigns := []*Campaign{}
//	err := db.Model(mdl).
//		Where(conditions).
//		Association("Campaigns").
//		Find(campaigns)
//
//	if err != nil {
//		panic(err)
//	}
//	mdl.Campaigns = campaigns
//	return campaigns, err
//}

// 取得場館的名字
func (mdl *Schedule) GetStudioNameByLocale(locale string) string {
	var studioName string
	if mdl.Studio != nil {
		studioName = mdl.Studio.GetNameByLocale(locale)
	}
	return studioName
}

// 取得老師的名字
func (mdl *Schedule) GetInstructorNamesByLocale(locale string) string {
	var instructorNames string
	if mdl.Instructors != nil {
		for _, instructor := range mdl.Instructors {
			if len(instructorNames) <= 0 {
				instructorNames = instructor.GetNameByLocale(locale)
			} else {

				instructorNames += " & " + instructor.GetNameByLocale(locale)
			}
		}
	}
	return instructorNames
}

// 取得Program的名字
func (mdl *Schedule) GetProgramNamesByLocale(locale string) string {
	var programNames string
	if mdl.Programs != nil {
		for _, program := range mdl.Programs {
			if len(programNames) <= 0 {
				programNames = program.GetNameByLocale(locale)
			} else {

				programNames += " & " + program.GetNameByLocale(locale)
			}
		}
	}
	return programNames
}
