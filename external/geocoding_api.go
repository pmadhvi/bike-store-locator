package external

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	log "github.com/sirupsen/logrus"
)

type GeocodeRequest struct {
	Address string `json:"address"`
	Region  string `json:"status"`
}

type GeocodeResponse struct {
	GeocodeResults GeocodeResultList `json:"results"`
	Status         string            `json:"status"`
}

type GeocodeResultList []GeocodeResult

type GeocodeResult struct {
	FormattedAddress string   `json:"formatted_address"`
	Geometry         Geometry `json:"geometry"`
}

type Geometry struct {
	Location LocationLatLng `json:"location"`
}

type LocationLatLng struct {
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lng"`
}

func GeoCodingApi(req *GeocodeRequest) (geocodeResp *GeocodeResponse, err error) {
	log.Info("Inside GeoCoding api!!")

	if req.Address == "" {
		return nil, errors.New("Address parameter missing")
	}
	baseUrl, err := url.Parse("https://maps.googleapis.com/maps/api/geocode/json")
	if err != nil {
		log.Errorf("Malformed URL:", err.Error())
		return
	}
	params := url.Values{}
	params.Set("address", req.Address)
	params.Set("region", req.Region)
	params.Set("key", ApiKey)
	baseUrl.RawQuery = params.Encode()

	fmt.Println("url :", baseUrl.String())
	resp, err := http.Get(baseUrl.String())
	if err != nil {
		log.Errorf("Failed to get the geocode for location: %s with error: %s\n", req.Address, err.Error())
		return
	}
	//Defer the call to close the response body
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&geocodeResp)
	if err != nil {
		log.Errorf("Failed to decode the response with error: %s\n", err.Error())
		return
	}
	return geocodeResp, nil
}
