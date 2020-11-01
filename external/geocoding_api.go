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

func (c Consumer) GeoCodingApi(req GeocodeRequest) (geocodeResp GeocodeResponse, err error) {
	log.Info("Inside GeoCoding api!!")

	if req.Address == "" {
		return GeocodeResponse{Status: "INVALID_REQUEST"}, errors.New("Address parameter missing")
	}
	reqURL := c.Host + c.Path

	parsedReqURL, err := url.Parse(reqURL)
	if err != nil {
		log.Error("Incorrect URL:\n", err.Error())
		return
	}

	params := url.Values{}
	params.Set("address", req.Address)
	params.Set("region", req.Region)
	params.Set("key", ApiKey)
	parsedReqURL.RawQuery = params.Encode()

	log.Infof("Request URL: %s", parsedReqURL.String())

	var resp *http.Response

	if resp, err = RunHTTP(parsedReqURL.String()); err != nil {
		log.Errorf("Failed to get the geocode for location: %s with error: %v\n", req.Address, err.Error())
		return
	}

	//Defer the call to close the response body
	defer resp.Body.Close()

	if err = json.NewDecoder(resp.Body).Decode(&geocodeResp); err != nil {
		log.Error("Failed to decode the response with error:\n", err.Error())
		return
	}

	fmt.Println("geocodeResp-->", geocodeResp)
	if geocodeResp.Status != "OK" && geocodeResp.Status != "ZERO_RESULTS" {
		log.Errorf("Failed to get the geocode for location: %s with status: %s", req.Address, geocodeResp.Status)
		return
	}

	if geocodeResp.Status == "ZERO_RESULTS" {
		log.Infof("The geocode request was successful, but returned no results, due to a non-existent address")
		return
	}
	return
}
