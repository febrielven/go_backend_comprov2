package controllers

import (
	"net/http"

	"github.com/febrielven/go_backend_comprov2/api/repository"
	"github.com/febrielven/go_backend_comprov2/api/responses"
	"github.com/jinzhu/gorm"
)

// type ProvinceController interface {
// 	GetProvinces(http.ResponseWriter, *http.Request)
// }

// func NewProvinceRepository(provinceRepository repository.ProvinceRepository) *Server {
// 	return &Server{
// 		provinceRepository: provinceRepository}
// }

func (repo *Module) NewProvinceRepository(db *gorm.DB) repository.ProvinceRepository {
	repo.provinceRepository = repository.NewProvinceRepository(db)
	return repo.provinceRepository
}

func (repo *Module) GetProvinces(w http.ResponseWriter, r *http.Request) {
	province, err := repo.provinceRepository.FindAll()

	if err != nil {
		println("err", err)

		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, province)
}
