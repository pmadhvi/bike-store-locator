// Package classification  BikeStoreLocator API.
//
// Swagger URL: http://localhost:9000/swaggerui/
//
// Terms Of Service:
//
//     Schemes: http, https
//     Host: localhost:8080
//     Version: 1.0.0
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Security:
//     - api_key:
//
//     SecurityDefinitions:
//     api_key:
//          type: apiKey
//          name: KEY
//          in: header
//
// swagger:meta
package main

import (
	"fmt"
	"net/http"

	"github.com/pmadhvi/tech-test/bike-locator-api/router"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("Hi, Welcome to Bike Locator Api!!")
	//Returns *mux.Router at "/"
	http.Handle("/", router.Router())

	//Checking for any error while listening for request on port 9000
	if err := http.ListenAndServe(fmt.Sprint(":", 9000), nil); err != nil {
		log.Fatal("Bike Locator Api crashed with error => ", err.Error())
	}
}
