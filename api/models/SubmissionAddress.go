package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type SubmissionAddress struct {
	ID            uuid.UUID `sql:"primary_key;type:uuid" json:"id"`
	SubmissionID  uuid.UUID `gorm:"column:submission_app_id;not null;type:uuid" json:"submission_id"`
	ProvinceID    int32     `gorm:"column:province_id;not null" json:"province_id"`
	CityID        int32     `gorm:"column:city_id;not null" json:"city_id"`
	DistrictID    int32     `gorm:"column:district_id;not null" json:"district_id"`
	SubDistrictID int32     `gorm:"column:sub_district_id;not null" json:"sub_district_id"`
	FullAddress   string    `gorm:"size:255" json:"full_address"`
	CreatedAt     time.Time `gorm:"column:created; not null" json:"created"`
}

type SubmissionsAddress []SubmissionAddress

func (SubmissionAddress) TableName() string {
	return "submission_adresss"
}

func (submissionAddress *SubmissionAddress) BeforeSubmAddressUUID(scope *gorm.Scope) error {
	uuid, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	return scope.SetColumn("ID", uuid)
}
