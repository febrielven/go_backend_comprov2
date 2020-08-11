package routes

// import (
// 	"net/http"

// 	"github.com/febrielven/go_backend_comprov2/api/middlewares"
// 	"github.com/gorilla/mux"
// )

// type Route struct {
// 	Patch   string
// 	Method  string
// 	Handler http.HandlerFunc
// }

// func InitiallizeRoutes(router *mux.Router, mainRoutes MainRoutes) {

// 	allRoutes := mainRoutes.Routes()

// 	for _, route := range allRoutes {
// 		handler := middlewares.Logger(route.Handler)
// 		router.HandleFunc(route.Patch, handler).Methods(route.Method)
// 	}

// }
