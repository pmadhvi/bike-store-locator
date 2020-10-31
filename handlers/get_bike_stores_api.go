package handlers

import (
	"encoding/json"
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

	geocodeRequest := external.GeocodeRequest{
		Address: location,
		Region:  region,
	}

	var (
		geocodeResponse external.GeocodeResponse
		storesResponse  external.FindPlaceResponse
	)

	if geocodeResponse, err = external.GeoCodingApi(geocodeRequest); err != nil {
		log.Error("Error: \n", err.Error())
		return
	}
	//TODO: Remove print later or convert it to log.
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

	if storesResponse, err = external.FindPlaceApi(findPlaceRequest); err != nil {
		log.Error("Failed to get the bike stores and error is:\n", err.Error())
		return
	}

	//TODO: Remove print later or convert it to log.
	fmt.Println("storesResponse::", storesResponse)

	// //TODO: Format the response back to models.BikeStore

	// for _, store := range storesResponse {
	// 	//TODO: run through the list of stores and Format the response back to models.BikeStore
	// }

	//TODO: return the response back
	payload, _ := json.Marshal(storesResponse)

	res.Header().Set("Content-Type", "application/json")
	res.Write([]byte(payload))

	return
}
