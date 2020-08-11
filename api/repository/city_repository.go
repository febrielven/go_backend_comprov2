package repository

import (
	"github.com/febrielven/go_backend_comprov2/api/models"
	"github.com/jinzhu/gorm"
)

type CityRepository interface {
	FindAll() (*models.Citys, error)
	FindByProvinceID(*models.City, uint32) (*models.Citys, error)
}

type cityRepositoryImpl struct {
	db *gorm.DB
}

func NewCityRepository(db *gorm.DB) *cityRepositoryImpl {
	return &cityRepositoryImpl{db}
}

func (server *cityRepositoryImpl) FindAll() (*models.Citys, error) {
	var err error
	var citys models.Citys

	err = server.db.Debug().Model(&models.City{}).Find(&citys).Error
	if err != nil {
		return nil, err
	}

	return &citys, nil
}

func (server *cityRepositoryImpl) FindByProvinceID(city *models.City, provinceID uint32) (*models.Citys, error) {
	var err error
	var citys models.Citys

	err = server.db.Debug().Model(&models.Province{}).Where("province_id= ?", provinceID).Find(&citys).Error
	if err != nil {
		return nil, err
	}

	return &citys, nil

}
