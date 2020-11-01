package external

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

type FindPlaceRequest struct {
	Input              string
	InputType          string
	Fields             []string
	LocationBiasType   string
	LocationBiasRadius int
	LocationBiasLat    float64
	LocationBiasLng    float64
}

type FindPlaceResponse struct {
	Places PlaceList `json:"candidates"`
	Status string    `json:"status"`
}
type PlaceList []Place

type Place struct {
	FormattedAddress string `json:"formatted_address"`
	Name             string `json:"name"`
}

func (c Consumer) FindPlaceApi(req FindPlaceRequest) (findPlacesResp FindPlaceResponse, err error) {
	log.Info("Inside Find places api!!")
	if req.Input == "" {
		return FindPlaceResponse{Status: "INVALID_REQUEST"}, errors.New("Input parameter missing")
	}

	if req.InputType == "" {
		return FindPlaceResponse{Status: "INVALID_REQUEST"}, errors.New("InputType parameter missing")
	}
	reqURL := c.Host + c.Path

	parsedReqURL, err := url.Parse(reqURL)
	if err != nil {
		log.Error("Incorrect URL:\n", err.Error())
		return
	}

	params := url.Values{}
	params.Set("input", req.Input)
	params.Set("inputtype", req.InputType)
	if len(req.Fields) > 0 {
		params.Set("fields", strings.Join(req.Fields, ","))
	}
	latlng := strconv.FormatFloat(req.LocationBiasLat, 'f', -1, 64) + "," + strconv.FormatFloat(req.LocationBiasLng, 'f', -1, 64)
	params.Set("locationbias", fmt.Sprintf("circle:%d@%s", req.LocationBiasRadius, latlng))
	params.Set("key", ApiKey)
	parsedReqURL.RawQuery = params.Encode()

	log.Infof("Request URL: %s", parsedReqURL.String())

	var resp *http.Response
	if resp, err = RunHTTP(parsedReqURL.String()); err != nil {
		log.Error("Failed to get places with error:\n", err.Error())
		return
	}

	//Defer the call to close the response body
	defer resp.Body.Close()

	if err = json.NewDecoder(resp.Body).Decode(&findPlacesResp); err != nil {
		log.Error("Failed to decode the response with error:\n", err.Error())
		return
	}

	if findPlacesResp.Status != "OK" && findPlacesResp.Status != "ZERO_RESULTS" {
		log.Errorf("Failed to get places with with status: %s", findPlacesResp.Status)
		return
	}

	if findPlacesResp.Status == "ZERO_RESULTS" {
		log.Infof("The findPlaces request was successful, but returned no results, due to a non-existent address")
		return
	}

	return
}
