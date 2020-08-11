package repository

import (
	"fmt"

	"github.com/febrielven/go_backend_comprov2/api/models"
	"github.com/jinzhu/gorm"
)

type ProvinceRepository interface {
	FindAll() (*models.Provinces, error)
}

type provincesRepositoryImpl struct {
	db *gorm.DB
}

func NewProvinceRepository(db *gorm.DB) *provincesRepositoryImpl {
	return &provincesRepositoryImpl{db}
}

func (server *provincesRepositoryImpl) FindAll() (*models.Provinces, error) {
	fmt.Println("FindAll...")
	var err error
	var provinces models.Provinces
	err = server.db.Debug().Model(&models.Province{}).Find(&provinces).Error
	if err != nil {
		return &provinces, err
	}
	return &provinces, nil
}
