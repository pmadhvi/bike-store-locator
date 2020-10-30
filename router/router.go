package router

import (
	"net/http"

	"github.com/pmadhvi/tech-test/bike-locator-api/handlers"
)

func Router() {
	http.HandleFunc("/bike-locator-api/healthz", handlers.HealthApi)
	http.HandleFunc("/bike-locator-api/location/{location}/radius/{radius}", handlers.GetBikeStoresApi)
}
