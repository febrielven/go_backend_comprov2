package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/febrielven/go_backend_comprov2/api/models"
	"github.com/febrielven/go_backend_comprov2/api/repository"
	"github.com/febrielven/go_backend_comprov2/api/responses"
	"github.com/jinzhu/gorm"
)

func (repo *Module) NewCityRepository(db *gorm.DB) repository.CityRepository {
	repo.cityRepository = repository.NewCityRepository(db)
	return repo.cityRepository
}

func (repo *Module) GetCity(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}

	city := models.City{}

	err = json.Unmarshal(body, &city)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}
	province_id := city.ProvinceID
	// var resultcity *models.City = &city
	// fmt.Println("resultcity (value)   :", *resultcity) //
	result, err := repo.cityRepository.FindByProvinceID(&city, uint32(province_id))

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
	}

	responses.JSON(w, http.StatusOK, result)
}
