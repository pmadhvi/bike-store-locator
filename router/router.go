package router

import (
	"github.com/gorilla/mux"
	"github.com/pmadhvi/tech-test/bike-locator-api/handlers"
)

//Router define handlers based on path
func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/bike-locator-api/healthz", handlers.HealthAPI)
	router.HandleFunc("/bike-locator-api/region/{region}/location/{location}/radius/{radius}", handlers.GetBikeStoresAPI)
	return router
}
