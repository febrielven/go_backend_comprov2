package controllers

import (
	"github.com/febrielven/go_backend_comprov2/api/repository"
	"github.com/jinzhu/gorm"
)

type Module struct {
	provinceRepository   repository.ProvinceRepository
	cityRepository       repository.CityRepository
	submissionRepository repository.SubmissionRepository
}

func (r *Module) InitializeRepository(db *gorm.DB) {
	r.NewProvinceRepository(db)
	r.NewCityRepository(db)
	r.NewSubmissionRepository(db)
}
