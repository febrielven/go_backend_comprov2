package models

import (
	"fmt"
	"time"

	"github.com/febrielven/go_backend_comprov2/api/requests"
	"github.com/google/uuid"
)

type SubmissionCustomer struct {
	ID           uuid.UUID `gorm:"primary_key;type:uuid" json:"id"`
	SubmissionID uuid.UUID `gorm: "column:submission_app_id;not null;type:uuid" json:"submission_id"`
	Name         string    `sql:"size:255"`
	PhoneNumber  string    `sql:"size:255"`
	Email        string    `sql:"size:255"`
	CreatedAt    time.Time `gorm:"column:created; not null" json:"created"`
}
type SubmissionCustomers []SubmissionCustomer

func (SubmissionCustomer) TableName() string {
	return "submission_customer"
}

// func (submissionCustomer *SubmissionCustomer) BeforeSubmCustomerUUID(scope *gorm.Scope) error {
// 	uuid, err := uuid.NewRandom()
// 	if err != nil {
// 		return err
// 	}
// 	return scope.SetColumn("ID", uuid)
// }

func (customer *SubmissionCustomer) Prepare(submission *requests.SubmCustomer) {

	var newID = uuid.Must(uuid.NewRandom())
	fmt.Printf("UUIDv4: %s\n", newID)
	if submission.ID == uuid.Nil {
		customer.ID = newID
	} else {
		customer.ID = submission.ID
	}

	customer.SubmissionID = submission.SubmissionID
	customer.Name = submission.Name
	customer.PhoneNumber = submission.PhoneNumber
	customer.Email = submission.Email
}
