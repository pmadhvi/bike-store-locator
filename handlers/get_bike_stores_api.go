package handlers

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/pmadhvi/tech-test/bike-locator-api/external"
	"github.com/pmadhvi/tech-test/bike-locator-api/models"
	log "github.com/sirupsen/logrus"
)

const (
	apiKey         = "AIzaSyAUeAoC5FJvYiSwS2sVBXxRMU1ojQMicwU"
	query          = "bicycle_store near Sergeltorg"
	placeType      = "bicycle_store"
	googleHost     = "https://maps.googleapis.com"
	localHost      = "http://127.0.0.1"
	textSearchPath = "/maps/api/place/textsearch/json"
	region         = "se"
)

//GetBikeStoresAPI returns the list of bike stores(name and address) for location sergeltorg and with radius of 2km.
func GetBikeStoresAPI(req *http.Request) (bikeStores models.BikeStores, err error) {
	//Feteching the quary parameters from url.
	params := mux.Vars(req)
	radius := params["radius"]

	//This is needed just for testing, else test will hit the actual google api
	var googleAPIHost string
	if os.Getenv("PORT") == "9000" {
		log.Info("Hurrey", os.Getenv("PORT"))
		googleAPIHost = googleHost
	} else {
		googleAPIHost = localHost
	}

	//Get the list of bikes stores
	bikeStores, err = findBikeStores(googleAPIHost, radius)
	if err != nil {
		log.Error(err.Error())
		return
	}

	return
}

//findBikeStores returns list of places
func findBikeStores(googleAPIHost, radius string) (storesResp models.BikeStores, err error) {
	//textSearchRequest request parameters
	textSearchRequest := external.TextSearchRequest{
		Query:     query,
		APIKey:    apiKey,
		Radius:    radius,
		Region:    region,
		PlaceType: placeType,
	}

	// Defining geocode consumer with Host and Path
	textSearchConsumer := external.Consumer{Host: googleAPIHost, Path: textSearchPath}

	storesResp, err = textSearchConsumer.TextSearchAPI(textSearchRequest)
	if err != nil {
		return
	}
	return
}
