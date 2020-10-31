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

const (
	ApiKey = "XYZ"
)

type FindPlaceFromTextRequest struct {
	Input              string
	InputType          string
	Fields             []string
	LocationBiasType   string
	LocationBiasRadius int
	LocationBiasLat    float64
	LocationBiasLng    float64
}

type FindPlaceFromTextResponse struct {
	Places PlaceList `json:"candidates"`
	Status string    `json:"status"`
}
type PlaceList []Place

type Place struct {
	FormattedAddress string `json:"formatted_address"`
	Name             string `json:"name"`
}

func FindPlaceApi(req *FindPlaceFromTextRequest) (placesResp *FindPlaceFromTextResponse, err error) {
	log.Info("Inside Find places api!!")
	if req.Input == "" {
		return nil, errors.New("Input parameter missing")
	}

	if req.InputType == "" {
		return nil, errors.New("InputType parameter missing")
	}
	baseUrl, err := url.Parse("https://maps.googleapis.com/maps/api/place/findplacefromtext/json")
	if err != nil {
		log.Errorf("Malformed URL: ", err.Error())
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
	baseUrl.RawQuery = params.Encode()

	fmt.Println("url :", baseUrl.String())
	resp, err := http.Get(baseUrl.String())
	if err != nil {
		log.Errorf("Failed to get places with error: %s\n", err.Error())
		return
	}
	//Defer the call to close the response body
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&placesResp)
	if err != nil {
		log.Errorf("Failed to decode the response with error: %s\n", err.Error())
		return
	}
	return
}
