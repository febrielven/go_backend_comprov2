package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/febrielven/go_backend_comprov2/api/models"
	"github.com/febrielven/go_backend_comprov2/api/repository"
	"github.com/febrielven/go_backend_comprov2/api/requests"
	"github.com/febrielven/go_backend_comprov2/api/responses"
	"github.com/jinzhu/gorm"
)

func (repo *Module) NewSubmissionRepository(db *gorm.DB) repository.SubmissionRepository {
	repo.submissionRepository = repository.NewSubmissionRepository(db)
	return repo.submissionRepository
}

func (repo *Module) CreateAndUpdateCustomer(w http.ResponseWriter, r *http.Request) {

	var err error
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}
	submission := requests.SubmCustomer{}

	err = json.Unmarshal(body, &submission)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}

	/** Submission App **/
	// var dataSubmission *requests.SubmCustomer = &submission
	// fmt.Println("dataSubmission: ", *dataSubmission)

	app := models.SubmissionApp{}
	app.Prepare(&submission)
	fmt.Println("update; ", &app)
	appFindById, err := repo.submissionRepository.GetAppById(&app, app.ID)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
	}

	if appFindById == nil {
		appCreate, err := repo.submissionRepository.SaveApp(&app)
		if err != nil {
			responses.ERROR(w, http.StatusInternalServerError, err)
		}
		submission.SubmissionID = appCreate.ID
		customer := models.SubmissionCustomer{}
		customer.Prepare(&submission)
		customerCreate, err := repo.submissionRepository.SaveCustomer(&customer)
		if err != nil {
			responses.ERROR(w, http.StatusInternalServerError, err)
		}

		w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, customerCreate.ID))
		responses.JSON(w, http.StatusOK, customerCreate)
	} else {
		// app.Prepare(&submission)
		// fmt.Println("update; ", &app)
		// _, err := repo.submissionRepository.UpdateApp(&app)
		// if err != nil {
		// 	responses.ERROR(w, http.StatusInternalServerError, err)
		// }
		customer := models.SubmissionCustomer{}
		customerFindBySubmId, err := repo.submissionRepository.GetCustomerBySubmissionId(&customer, app.ID)
		if err != nil {
			responses.ERROR(w, http.StatusInternalServerError, err)
		}
		submission.ID = customerFindBySubmId.ID
		customer.Prepare(&submission)
		customerUpdate, err := repo.submissionRepository.UpdateCustomer(&customer, submission.SubmissionID)
		if err != nil {
			responses.ERROR(w, http.StatusInternalServerError, err)
		}

		responses.JSON(w, http.StatusOK, customerUpdate)
	}

}

func (repo *Module) CreateAndUpdateAddress(w http.ResponseWriter, r *http.Request) {
	var err error

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}

	address := models.SubmissionAddress{}
	err = json.Unmarshal(body, &address)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}

	fmt.Println("dataSubmission: ", address)

}
