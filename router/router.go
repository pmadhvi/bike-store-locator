package router

import (
	"github.com/gorilla/mux"
	"github.com/pmadhvi/tech-test/bike-locator-api/handlers"
)

//Router define handlers based on path
func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/bikestoresapi/health", handlers.HealthHandler)
	router.HandleFunc("/bikestoresapi/region/{region}/location/{location}/radius/{radius}", handlers.GetBikeStoresHandler)
	router.HandleFunc("/", handlers.NotFoundHandler)
	return router
}
