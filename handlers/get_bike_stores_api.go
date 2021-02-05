package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pmadhvi/tech-test/bike-locator-api/external"
	"github.com/pmadhvi/tech-test/bike-locator-api/models"
	log "github.com/sirupsen/logrus"
)

const (
	apiKey         = "AIzaSyCIkeRRD02JxTZZIRzXn-eLsgBRUulICok"
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

	//Get the list of bikes stores
	bikeStores, err = findBikeStores(googleHost, radius)
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
