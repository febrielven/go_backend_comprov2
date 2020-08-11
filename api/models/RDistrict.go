package models

import "time"

type District struct {
	ID                 int32                `sql:"AUTO_INCREMENT" gorm:"primary_key" json:"id"`
	CityID             int32                `gorm: "column:city_id;not null" json:"city_id"`
	Desc               string               `sql:"size:255"`
	CreatedAt          time.Time            `gorm:"column:created; not null" json:"created"`
	SubDistricts       []*SubDistrict       `gorm:foreignkey:DistrictID`
	SubmissionsAddress []*SubmissionAddress `gorm:"foreignkey:DistrictID"`
}

type Districts []District

func (District) TableName() string {
	return "ref_district"
}
