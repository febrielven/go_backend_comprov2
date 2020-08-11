package routes

// import (
// 	"net/http"

// 	"github.com/febrielven/go_backend_comprov2/api/controllers"
// )

// type MainRoutes interface {
// 	Routes() []*Route
// }

// type MainRoutesImpl struct {
// 	provinceController controllers.ProvinceController
// }

// func NewMainRoutes(provinceController controllers.ProvinceController) *MainRoutesImpl {
// 	return &MainRoutesImpl{provinceController}
// }

// func (r *MainRoutesImpl) Routes() []*Route {
// 	return []*Route{
// 		&Route{
// 			Patch:   "/provinces",
// 			Method:  http.MethodGet,
// 			Handler: r.provinceController.GetProvinces,
// 		},
// 	}
// }
