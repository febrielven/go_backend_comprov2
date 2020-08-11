package requests

import "github.com/google/uuid"

type SubmCustomer struct {
	ID           uuid.UUID `json:"id"`
	SubmissionID uuid.UUID `json:"submission_id"`
	ProductID    int32     `json:"product_id"`
	Steps        int32     `json:"steps"`
	Name         string    `json:"name"`
	PhoneNumber  string    `json:"phone_number"`
	Email        string    `json:"email"`
}
