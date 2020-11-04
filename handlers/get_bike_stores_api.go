package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/pmadhvi/tech-test/bike-locator-api/external"
	"github.com/pmadhvi/tech-test/bike-locator-api/models"
	log "github.com/sirupsen/logrus"
)

//APIKey for google api
const APIKey = "XYZ123"

//GetBikeStoresAPI returns the list of bike stores(name and address) for location sergeltorg and with radius of 2km.
func GetBikeStoresAPI(res http.ResponseWriter, req *http.Request) {
	log.Info("Inside Bike store locator Api!!")
	fmt.Print(req.URL.Query())
	fmt.Printf("radios = %v", req.URL.Query().Get("radius"))
	//Feteching the quary parameters from url.
	location := req.URL.Query().Get("location")
	fmt.Printf("location = %v", req.URL.Query().Get("location"))
	region := req.URL.Query().Get("region")
	fmt.Printf("region = %v", req.URL.Query().Get("region"))
	radius, err := strconv.Atoi(req.URL.Query().Get("radius"))
	fmt.Println("err ", err) 
	fmt.Printf("radius = %v\n", radius) 
	if err != nil {
		log.Error("Failed to convert string to integer with error: ", err.Error())
		return
	}
	

	// Defining geocode consumer with Host and Path
	geocodeConsumer := external.Consumer{Host: "https://maps.googleapis.com", Path: "/maps/api/geocode/json"}

	//Geocode request parameters
	geocodeRequest := external.GeocodeRequest{
		Address: location,
		APIKey:  APIKey,
		Region:  region,
	}

	var (
		geocodeResponse *models.Geocode
		storesResponse  *external.FindPlaceResponse
	)

	geocodeResponse, err = geocodeConsumer.GeoCodingAPI(geocodeRequest)
	if err != nil {
		log.Error(err.Error())
		return
	}
	//TODO: Remove print later or convert it to log.
	fmt.Println("geocodeResponse in bike locator::", geocodeResponse)

	findPlaceRequest := external.FindPlaceRequest{
		Input:              location,
		InputType:          "textquery",
		Fields:             []string{"name, formatted_address"},
		LocationBiasType:   "circle",
		LocationBiasRadius: radius,
		LocationBiasLat:    geocodeResponse.Latitude,
		LocationBiasLng:    geocodeResponse.Longitude,
	}

	// Defining geocode consumer with Host and Path
	findPlaceConsumer := external.Consumer{Host: "https://maps.googleapis.com", Path: "maps/api/place/textsearch/json"}

	storesResponse, err = findPlaceConsumer.FindPlacesAPI(findPlaceRequest)
	if err != nil {
		log.Error("Failed to get the bike stores and error is: ", err.Error())
		return
	}

	// //TODO: Remove print later or convert it to log.
	fmt.Println("storesResponse::", storesResponse)

	// //TODO: Format the response back to models.BikeStore

	// for _, store := range storesResponse {
	// 	//TODO: run through the list of stores and Format the response back to models.BikeStore
	// }

	//TODO: return the response back
	//payload, _ := json.Marshal(storesResponse)

	res.Header().Set("Content-Type", "application/json")
	res.Write([]byte("Hello World"))

	return
}
