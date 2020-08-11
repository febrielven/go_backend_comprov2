
package models

import "time"

type SubDistrict struct {
	ID                 int32                `sql:"AUTO_INCREMENT" gorm:"primary_key" json:"id"`
	DistrictID         int32                `gorm: "column:district_id;not null" json:"district_id"`
	Desc               string               `sql:"size:255"`
	CreatedAt          time.Time            `gorm:"column:created; not null" json:"created"`
	SubmissionsAddress []*SubmissionAddress `gorm:"foreignkey:SubDistrictID"`
}

type SubDistricts []SubDistrict

func (SubDistrict) TableName() string {
	return "ref_subdistrict"
}
