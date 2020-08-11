package models

import (
	"fmt"
	"time"

	"github.com/febrielven/go_backend_comprov2/api/requests"
	"github.com/google/uuid"
)

type SubmissionApp struct {
	ID                 uuid.UUID           `gorm:"primary_key;type:uuid" json:"id"`
	ProductID          int32               `gorm: "column:product_id;not null" json:"product_id"`
	Steps              int32               `gorm: "column:steps"; json:"steps"`
	CreatedAt          time.Time           `gorm:"column:created; not null" json:"created"`
	SubmissionCustomer *SubmissionCustomer `gorm:"foreignkey:SubmissionID"`
	SubmissionAddress  *SubmissionAddress  `gorm:"foreignkey:SubmissionID"`
}

type SubmissionApps []SubmissionApp

func (SubmissionApp) TableName() string {
	return "submission_app"
}

// func (submissionApp *SubmissionApp) BeforeSubmAppUUID(scope *gorm.Scope) error {
// 	uuid, err := uuid.NewRandom()
// 	if err != nil {
// 		return err
// 	}
// 	return scope.SetColumn("ID", uuid)
// }

func (app *SubmissionApp) Prepare(submission *requests.SubmCustomer) {

	// fmt.Printf("UUIDv4: %s\n", uuid)
	var id = uuid.Must(uuid.NewUUID())
	fmt.Println("SubmissionID: ", submission.SubmissionID)
	if submission.SubmissionID == uuid.Nil {
		app.ID = id
	} else {
		app.ID = submission.SubmissionID
	}

	app.ProductID = submission.ProductID
	app.Steps = submission.Steps

	fmt.Println("PRODUT ID", app.Steps)
}
