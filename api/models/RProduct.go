package models

import "time"

type Product struct {
	ID             int32            `sql:"AUTO_INCREMENT" gorm:"primary_key" json:"id"`
	CreatedAt      time.Time        `gorm:"column:created; not null" json:"created"`
	SubmissionApps []*SubmissionApp `gorm:"foreignkey:ProductID"`
}

type Products []Product

func (Product) TableName() string {
	return "ref_product"
}
