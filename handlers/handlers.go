package handlers

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
)

//HealthHandler returns the health status of the application server.
func HealthHandler(res http.ResponseWriter, req *http.Request) {
	log.Info("HealthHandler!!")
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	res.Write([]byte(`{"message": "alive"}`))
}

//NotFoundHandler is for handling invalid url path
func NotFoundHandler(res http.ResponseWriter, req *http.Request) {
	log.Info("NotFoundHandler!!")
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusNotFound)
	res.Write([]byte(`{"message": "not found"}`))
}

//GetBikeStoresHandler handler to find the list of bike stores
func GetBikeStoresHandler(res http.ResponseWriter, req *http.Request) {
	log.Info("GetBikeStoresHandler!!")
	res.Header().Set("Content-Type", "application/json")
	//Get the list of bike stores
	bikeStores, err := GetBikeStoresAPI(req)
	if err != nil {
		respondErrorJSON(res, err)
	}
	//Converting from models.BikeStores to json format
	response, err := json.Marshal(bikeStores)
	if err != nil {
		respondErrorJSON(res, err)
	}
	respondJSON(res, response)
}

//respondJSON returns payload in json format
func respondJSON(w http.ResponseWriter, response []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}

//respondErrorJSON returns error in json format
func respondErrorJSON(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(err.Error()))
}
