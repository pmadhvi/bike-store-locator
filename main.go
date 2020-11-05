package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/pmadhvi/tech-test/bike-locator-api/router"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("Hi, Welcome to Bike Locator Api!!")
	//Returns *mux.Router at "/"
	http.Handle("/", router.Router())

	os.Setenv("PORT", "9000")
	port := os.Getenv("PORT")
	//Checking for any error while listening for request on port 9000
	if err := http.ListenAndServe(fmt.Sprint(":", port), nil); err != nil {
		log.Fatal("Bike Locator Api crashed with error => ", err.Error())
	}
}
