package external

import (
	"encoding/json"
	"errors"
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

func GeoCodingApi(req GeocodeRequest) (geocodeResp GeocodeResponse, err error) {
	log.Info("Inside GeoCoding api!!")

	if req.Address == "" {
		return GeocodeResponse{Status: "INVALID_REQUEST"}, errors.New("Address parameter missing")
	}
	reqURL, err := url.Parse("https://maps.googleapis.com/maps/api/geocode/json")
	if err != nil {
		log.Error("Incorrect URL:\n", err.Error())
		return
	}

	params := url.Values{}
	params.Set("address", req.Address)
	params.Set("region", req.Region)
	params.Set("key", ApiKey)
	reqURL.RawQuery = params.Encode()

	log.Infof("Request URL: %s", reqURL.String())

	var resp *http.Response

	if resp, err = RunHTTP(reqURL.String()); err != nil {
		log.Errorf("Failed to get the geocode for location: %s with error: %v\n", req.Address, err.Error())
		return
	}

	//Defer the call to close the response body
	defer resp.Body.Close()

	if err = json.NewDecoder(resp.Body).Decode(&geocodeResp); err != nil {
		log.Error("Failed to decode the response with error:\n", err.Error())
		return
	}

	if geocodeResp.Status != "OK" {
		log.Errorf("Failed to get the geocode for location: %s with status: %s\n", req.Address, geocodeResp.Status)
		return
	}
	return
}
