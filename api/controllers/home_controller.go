package controllers

import (
	"fmt"
	"net/http"

	"github.com/febrielven/go_backend_comprov2/api/responses"
)

// type HomeController interface {
// 	Home(http.ResponseWriter, *http.Request)
// }

func (repo *Module) Home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Welcome To This Awesome API")
	responses.JSON(w, http.StatusOK, "Wellcome to Api")
}
