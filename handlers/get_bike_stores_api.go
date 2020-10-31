package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/pmadhvi/tech-test/bike-locator-api/external"
	log "github.com/sirupsen/logrus"
)

//GetBikeStoresAPI returns the list of bike stores(name and address) for location sergeltorg and with radius of 2km.
func GetBikeStoresAPI(res http.ResponseWriter, req *http.Request) {
	log.Info("Inside Bike store locator Api!!")
	location := req.URL.Query().Get("location")
	radius, err := strconv.Atoi(req.URL.Query().Get("radius"))
	if err != nil {
		log.Error("Failed to convert string to integer with error: \n", err.Error())
		return
	}
	region := req.URL.Query().Get("region")
	var (
		geocodeResponse external.GeocodeResponse
		storesResponse  external.FindPlaceResponse
	)

	geocodeRequest := external.GeocodeRequest{
		Address: location,
		Region:  region,
	}

	geocodeResponse, err = external.GeoCodingApi(geocodeRequest)
	if err != nil {
		log.Error("Failed to get the geocode for location: %s and error is: \n", location, err.Error())
		return
	}
	fmt.Println("geocodeResponse::", geocodeResponse)

	//TODO: Format the response back to models.Geocode

	findPlaceRequest := external.FindPlaceRequest{
		Input:              location,
		InputType:          "textquery",
		Fields:             []string{"name, formatted_address"},
		LocationBiasType:   "circle",
		LocationBiasRadius: radius,
		//TODO: Need to get from geocoderesp
		LocationBiasLat: -122.222641,
		LocationBiasLng: 47.6918452,
	}
	storesResponse, err = external.FindPlaceApi(findPlaceRequest)
	if err != nil {
		log.Error("Failed to get the bike stores and error is:\n", err.Error())
		return
	}
	fmt.Println("storesResponse::", storesResponse)
	// //TODO: Format the response back to models.BikeStore

	// for _, store := range storesResponse {
	// 	//TODO: run through the list of stores and Format the response back to models.BikeStore
	// }

	//TODO: return the response back
	// res.WriteHeader(200)
	// res.Header().Set("Content-Type", "application/json")

	return
}
