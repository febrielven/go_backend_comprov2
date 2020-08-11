package models

import (
	"time"
)

type Province struct {
	ID                 int32               `sql:"AUTO_INCREMENT" gorm:"primary_key"`
	Desc               string              `sql:"size:255"`
	CreatedAt          time.Time           `gorm:"column:created; not null" json:"created"`
	Citys              []City              `gorm:"foreignkey:ProvinceID"`
	SubmissionsAddress []SubmissionAddress `gorm:"foreignkey:ProvinceID"`
}

type Provinces []Province

func (Province) TableName() string {
	return "ref_province"
}
