package controllers

import (
	"github.com/febrielven/go_backend_comprov2/api/middlewares"
	"github.com/gorilla/mux"
)

func (r *Module) InstallRoutes(router *mux.Router) {
	router.HandleFunc("/get_provinces", middlewares.Logger(r.GetProvinces)).Methods("GET")
	// router.HandleFunc("/get_city", middlewares.Logger(r.GetCity)).Methods("GET")
	router.HandleFunc("/get_city", middlewares.Logger(r.GetCity)).Methods("POST")
	router.HandleFunc("/save_submission_customer", middlewares.Logger(r.CreateAndUpdateCustomer)).Methods("POST")
	router.HandleFunc("/save_submission_address", middlewares.Logger(r.CreateAndUpdateAddress))
}
