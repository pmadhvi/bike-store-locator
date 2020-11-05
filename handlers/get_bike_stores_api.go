package handlers

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/pmadhvi/tech-test/bike-locator-api/external"
	"github.com/pmadhvi/tech-test/bike-locator-api/models"
	log "github.com/sirupsen/logrus"
)

//APIKey for google api
const (
	apiKey        = "XYZ123"
	googleHost    = "https://maps.googleapis.com"
	localHost     = "http://127.0.0.1"
	geocodePath   = "/maps/api/geocode/json"
	findPlacePath = "/maps/api/place/findplacefromtext/json"
)

//GetBikeStoresAPI returns the list of bike stores(name and address) for location sergeltorg and with radius of 2km.
func GetBikeStoresAPI(req *http.Request) (bikeStores models.BikeStores, err error) {
	//Feteching the quary parameters from url.
	params := mux.Vars(req)
	//TODO: remove this print
	fmt.Println("params =>", params)
	location := params["location"]
	region := params["region"]
	radius, err := strconv.Atoi(params["radius"])
	if err != nil {
		log.Error("Failed to convert string to integer with error: ", err.Error())
		return
	}

	//This is needed just for testing, else test will hit the actual google api
	var googleAPIHost string
	if os.Getenv("PORT") == "9000" {
		log.Info("Hurrey", os.Getenv("PORT"))
		googleAPIHost = googleHost
	} else {
		googleAPIHost = localHost
	}

	//Get the geocode for location and region
	geocodeResponse, err := getGeocodes(googleAPIHost, location, region)
	if err != nil {
		log.Error(err.Error())
		return
	}
	//TODO: Remove print later or convert it to log.
	fmt.Println("geocodeResponse in bike locator::", geocodeResponse)

	//Get the list of bikes stores
	storesResponse, err := findPlaces(googleAPIHost, location, radius, geocodeResponse.Latitude, geocodeResponse.Longitude)
	if err != nil {
		log.Error(err.Error())
		return
	}
	//TODO: Remove print later or convert it to log.
	fmt.Println("storesResponse::", storesResponse)
	return
}

func getGeocodes(googleAPIHost, location, region string) (geocode *models.Geocode, err error) {
	//Geocode request parameters
	geocodeRequest := external.GeocodeRequest{
		Address: location,
		APIKey:  apiKey,
		Region:  region,
	}

	//Defining geocode consumer with Host and Path
	geocodeConsumer := external.Consumer{Host: googleAPIHost, Path: geocodePath}

	//Calling the external GeoCodingAPI which inturns calls google api to get geocode
	geocode, err = geocodeConsumer.GeoCodingAPI(geocodeRequest)
	if err != nil {
		return
	}
	return
}

func findPlaces(googleAPIHost, location string, radius int, lat, lng float64) (storesResp models.BikeStores, err error) {
	//findPlaceRequest request parameters
	findPlaceRequest := external.FindPlaceRequest{
		Input:              location,
		InputType:          "textquery",
		APIKey:             apiKey,
		Fields:             []string{"name, formatted_address"},
		LocationBiasType:   "circle",
		LocationBiasRadius: radius,
		LocationBiasLat:    lat,
		LocationBiasLng:    lng,
	}

	// Defining geocode consumer with Host and Path
	findPlaceConsumer := external.Consumer{Host: googleAPIHost, Path: findPlacePath}

	storesResp, err = findPlaceConsumer.FindPlacesAPI(findPlaceRequest)
	if err != nil {
		return
	}
	return
}
