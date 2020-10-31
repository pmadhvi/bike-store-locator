package main

import (
	"net/http"

	"github.com/pmadhvi/tech-test/bike-locator-api/router"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("Hi, Welcome to Bike Locator Api!!")
	router.Router()
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		log.Fatal("Bike Locator Api crashed!!")
	}
}
