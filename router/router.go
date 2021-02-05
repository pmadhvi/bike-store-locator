package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pmadhvi/tech-test/bike-locator-api/handlers"
)

//Router define handlers based on path
func Router() *mux.Router {
	router := mux.NewRouter()

	// swagger:operation GET /bikestoresapi/health HealthApi
	// ---
	// summary: Returns the helath of the api.
	// description: Returns the helath of the api
	// responses:
	//   "200":
	//     "$ref": "#/responses/healthRes"
	//   "400":
	//     "$ref": "#/responses/badReq"
	//   "404":
	//     "$ref": "#/responses/notFoundReq"
	router.HandleFunc("/bikestoresapi/health", handlers.HealthHandler)

	// swagger:operation GET /bikestoresapi/radius/{radius} GetBikeStoreApi
	// ---
	// summary: Returns the list of bike stores with the radius.
	// description: Returns the list of bike stores with the radius.
	// parameters:
	// - name: radius
	//   in: path
	//   description: radius of the area to look into
	//   type: integer
	//   format: int64
	//   required: true
	// responses:
	//   "200":
	//     "$ref": "#/responses/bikeStoresRes"
	//   "400":
	//     "$ref": "#/responses/badReq"
	//   "404":
	//     "$ref": "#/responses/notFoundReq"
	router.HandleFunc("/bikestoresapi/radius/{radius}", handlers.GetBikeStoresHandler)

	sh := http.StripPrefix("/swaggerui/", http.FileServer(http.Dir("./swaggerui/")))
	router.PathPrefix("/swaggerui/").Handler(sh)

	return router
}
