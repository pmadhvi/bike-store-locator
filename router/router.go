package router

import (
	"net/http"

	"github.com/pmadhvi/tech-test/bike-locator-api/handlers"
)

func Router() {
	http.HandleFunc("/bike-locator-api/healthz", handlers.HealthAPI)
	http.HandleFunc("/bike-locator-api/region/{region}/location/{location}/radius/{radius}", handlers.GetBikeStoresAPI)
}
