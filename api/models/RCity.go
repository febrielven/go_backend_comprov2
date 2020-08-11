package models

import "time"

type City struct {
	ID                 int32                `sql:"AUTO_INCREMENT" gorm:"primary_key" json:"id"`
	ProvinceID         int32                `sql:"size:11" gorm:"column:province_id" json:"province_id"`
	Desc               string               `sql:"size:255" gorm:"column:descs" json:"desc"`
	CreatedAt          time.Time            `gorm:"column:created; not null" json:"created"`
	Districts          []*District          `gorm:"foreignkey:CityID"`
	SubmissionsAddress []*SubmissionAddress `gorm:"foreignkey:CityID"`
}

type Citys []City

func (City) TableName() string {
	return "ref_city"
}
