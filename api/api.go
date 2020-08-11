package api

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/febrielven/go_backend_comprov2/api/controllers"
	"github.com/febrielven/go_backend_comprov2/api/database"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var (
	port        = flag.Int("P", 5070, "Set Port")
	resetTables = flag.Bool("rt", false, "reset tables")
)

var module = controllers.Module{}

func Run() {
	flag.Parse()
	db := database.Connect()

	if db != nil {
		defer db.Close()
	}

	fmt.Println("Database Connected...")

	// db.Debug().AutoMigrate(
	// 	&models.City{})
	// provinceRepository := repository.NewProvinceRepository(db)
	// provinceController := controllers.NewProvinceRepository(provinceRepository)
	// routeMain := routes.NewMainRoutes(provinceController)
	// var server = controllers.ProvinceControllerImpl{}
	// router := mux.NewRouter().StrictSlash(true)
	// routes.InitiallizeRoutes(router, routeMain)
	module.InitializeRepository(db)
	router := mux.NewRouter().StrictSlash(true)
	module.InstallRoutes(router)

	headers := handlers.AllowedHeaders([]string{"Content-Type", "X-Request", "Location", "Entity", "Accept"})
	methods := handlers.AllowedMethods([]string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete})
	origins := handlers.AllowedOrigins([]string{"*"})

	fmt.Println("API lintening on", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), handlers.CORS(headers, methods, origins)(router)))
}
