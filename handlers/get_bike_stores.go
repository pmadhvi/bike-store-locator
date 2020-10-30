package handlers

import (
	"net/http"

	_ "github.com/pmadhvi/tech-test/bike-locator-api/external"
	_ "github.com/pmadhvi/tech-test/bike-locator-api/models"
	log "github.com/sirupsen/logrus"
)

func GetBikeStoresApi(res http.ResponseWriter, req *http.Request) {
	log.Info("Inside Bike store locator api!!")
	location := req.URL.Query().Get("location")
	radius := req.URL.Query().Get("radius")

	var (
		geocodeResponse external.Geocodes
		storesResponse  external.Stores
	)
	geocode_response, err = external.GeoCodingApi(location)
	if err != nil {
		log.Error("Failed to get the geocode for location: %s and err\n", location, err)
		//TODO: chnage the err to Proper format
	}

	//TODO: Format the response back to models.Geocode

	storesResponse, err := external.FindPlaceApi()
	if err != nil {
		log.Error("Failed to get the bike stores, error:\n", err)
		//TODO: chnage the err to Proper format
	}
	//TODO: Format the response back to models.BikeStore

	for _, store := range storesResponse {
		//TODO: run through the list of stores and Format the response back to models.BikeStore
	}
	return BikeStore
}
