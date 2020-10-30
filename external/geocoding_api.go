package external

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	log "github.com/sirupsen/logrus"
)

type GeocodeResponse struct {
	Results ResultList `json:"results"`
	Status  string     `json:"status"`
}
type ResultList []Result

type Result struct {
	Geometry Geometry `json:"geometry"`
}

type Geometry struct {
	Location Location `json:"location"`
}

type Location struct {
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lng"`
}

const (
	key = "XYZ"
)

func GeoCodingApi(location) (geocodeResp GeocodeResponse, err error) {
	log.Info("Inside GeoCoding api!!")
	baseUrl, err := url.Parse("https://maps.googleapis.com/maps/api/geocode/json")
	if err != nil {
		log.Errorf("Malformed URL: ", err.Error())
		return
	}
	params := url.Values{}
	params.Add("address", location)
	params.Add("key", key)
	baseUrl.RawQuery = params.Encode()

	fmt.Println("url :", baseUrl.String())
	resp, err := http.Get(baseUrl.String())
	if err != nil {
		log.Errorf("Failed to get the geocode for location: %s with error: %s\n", location, err.Error())
		return
	}
	//Defer the call to close the response body
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&geocodeResp)
	if err != nil {
		log.Errorf("Failed to decode the response with error: %s\n", err.Error())
		return
	}
	return
}
