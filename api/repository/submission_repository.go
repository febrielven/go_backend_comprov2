package repository

import (
	"fmt"

	"github.com/febrielven/go_backend_comprov2/api/models"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type SubmissionRepository interface {
	GetAppById(*models.SubmissionApp, uuid.UUID) (*models.SubmissionApp, error)
	SaveApp(*models.SubmissionApp) (*models.SubmissionApp, error)
	UpdateApp(*models.SubmissionApp) (*models.SubmissionApp, error)
	GetCustomerBySubmissionId(*models.SubmissionCustomer, uuid.UUID) (*models.SubmissionCustomer, error)
	SaveCustomer(*models.SubmissionCustomer) (*models.SubmissionCustomer, error)
	UpdateCustomer(*models.SubmissionCustomer, uuid.UUID) (*models.SubmissionCustomer, error)
	GetAddressBySubmissionId(*models.SubmissionAddress, uuid.UUID) (*models.SubmissionAddress, error)
	// SaveAddress(*models.SubmissionAddress) (*models.SubmissionsAddress, error)
	// UpdateAddress(*models.SubmissionAddress, uuid.UUID) (*models.SubmissionAddress, error)
}

type submissionRepositoryImpl struct {
	db *gorm.DB
}

func NewSubmissionRepository(db *gorm.DB) *submissionRepositoryImpl {
	return &submissionRepositoryImpl{db}
}

func (server *submissionRepositoryImpl) GetAppById(app *models.SubmissionApp, id uuid.UUID) (*models.SubmissionApp, error) {
	var err error

	err = server.db.Debug().Model(models.SubmissionApp{}).Where("id= ?", id).Take(&app).Error

	if gorm.IsRecordNotFoundError(err) {
		fmt.Println("noRecord")
		return nil, nil
	}

	if err != nil {
		return &models.SubmissionApp{}, err
	}

	return app, nil
}

func (server *submissionRepositoryImpl) SaveApp(app *models.SubmissionApp) (*models.SubmissionApp, error) {
	var err error
	err = server.db.Debug().Create(&app).Error
	if err != nil {
		return &models.SubmissionApp{}, err
	}

	return app, nil
}
func (server *submissionRepositoryImpl) UpdateApp(app *models.SubmissionApp) (*models.SubmissionApp, error) {
	var err error
	err = server.db.Debug().Model(&app).Update(app).Error
	if err != nil {
		return &models.SubmissionApp{}, err
	}
	return app, nil

}

func (server *submissionRepositoryImpl) GetCustomerBySubmissionId(customer *models.SubmissionCustomer, submissionId uuid.UUID) (*models.SubmissionCustomer, error) {
	var err error

	err = server.db.Debug().Model(models.SubmissionCustomer{}).Where("submission_id= ?", submissionId).Take(&customer).Error
	if err != nil {
		return &models.SubmissionCustomer{}, err
	}

	return customer, nil
}

func (server *submissionRepositoryImpl) SaveCustomer(customer *models.SubmissionCustomer) (*models.SubmissionCustomer, error) {
	var err error
	err = server.db.Debug().Create(&customer).Error
	if err != nil {
		return &models.SubmissionCustomer{}, err
	}
	return customer, nil
}

func (server *submissionRepositoryImpl) UpdateCustomer(customer *models.SubmissionCustomer, submissionId uuid.UUID) (*models.SubmissionCustomer, error) {
	var err error
	err = server.db.Model(&customer).Where("submission_id = ?", submissionId).Update(customer).Error
	if err != nil {
		return &models.SubmissionCustomer{}, err
	}
	return customer, nil
}

func (server *submissionRepositoryImpl) GetAddressBySubmissionId(address *models.SubmissionAddress, submissionId uuid.UUID) (*models.SubmissionAddress, error) {
	var err error

	err = server.db.Debug().Model(models.SubmissionAddress{}).Where("submission_app_id = ", submissionId).Take(&address).Error
	if err != nil {
		return &models.SubmissionAddress{}, err
	}

	return address, nil
}

// func (server *submissionRepositoryImpl) SaveAddress(address *models.SubmissionAddress) (*models.SubmissionAddress, error) {

// 	return nil, nil
// }
