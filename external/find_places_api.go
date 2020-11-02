package external

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

type FindPlaceRequest struct {
	Input              string
	InputType          string
	APIKey             string
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

//FindPlacesApi as a method with Consumer as receiver and FindPlaceRequest as parameter and returns response of type FindPlaceResponse and err of type error.
func (c Consumer) FindPlacesApi(req FindPlaceRequest) (findPlacesResp *FindPlaceResponse, err error) {
	//Validating required request parametersand APIKey.
	if req.Input == "" {
		log.Error("Input request parameter missing")
		err = fmt.Errorf("status_code: %d and error_message: %s", 400, "Input request parameter missing")
		return nil, AppError{Operation: "FindPlaceApi", Err: err}
	}

	if req.InputType == "" {
		log.Error("InputType request parameter missing")
		err = fmt.Errorf("status_code: %d and error_message: %s", 400, "InputType request parameter missing")
		return nil, AppError{Operation: "FindPlaceApi", Err: err}
	}

	if req.APIKey == "" {
		log.Error("APIKey is missing")
		err = fmt.Errorf("status_code: %d and error_message: %s", 400, "APIKey is missing")
		return nil, AppError{Operation: "FindPlaceApi", Err: err}
	}

	//Composing reqURL from Host and Path of Consumer struct.
	reqURL := c.Host + c.Path

	//Parsing the reqURL to convert it to type *url.Url and checking for error if any.
	parsedReqURL, err := url.Parse(reqURL)
	if err != nil {
		log.Error("Incorrect URL:", err.Error())
		err = fmt.Errorf("status_code: %d and error_message: %s", 400, "Incorrect URL")
		return nil, AppError{Operation: "FindPlaceApi", Err: err}
	}

	//Adding request params to url and encoding it.
	params := url.Values{}
	params.Set("input", req.Input)
	params.Set("inputtype", req.InputType)
	if len(req.Fields) > 0 {
		params.Set("fields", strings.Join(req.Fields, ","))
	}
	latlng := strconv.FormatFloat(req.LocationBiasLat, 'f', -1, 64) + "," + strconv.FormatFloat(req.LocationBiasLng, 'f', -1, 64)
	params.Set("locationbias", fmt.Sprintf("circle:%d@%s", req.LocationBiasRadius, latlng))
	params.Set("key", req.APIKey)
	parsedReqURL.RawQuery = params.Encode()

	log.Infof("Request URL: %s", parsedReqURL.String())

	var resp *http.Response
	//RunHTTP performs the http.Get under the hood and get's the response from server.
	resp, err = RunHTTP(parsedReqURL.String())
	if err != nil {
		log.Errorf("Failed to get places with error => %v\n", err.Error())
		return nil, AppError{Operation: "FindPlaceApi", Err: err}
	}

	//Defer the call to close the response body at the end of function execution.
	defer resp.Body.Close()

	//Read the response body into memory
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Errorf("Failed to read the response body with error => %v\n", err.Error())
		return nil, AppError{Operation: "FindPlaceApi", Err: err}
	}

	if resp.StatusCode == 404 {
		err = fmt.Errorf("status_code: %d and error_message: %s", resp.StatusCode, string(respBody))
		return nil, AppError{Operation: "FindPlaceApi", Err: err}
	}

	if resp.StatusCode >= 400 {
		err = fmt.Errorf("status_code: %d and error_message: %s", resp.StatusCode, string(respBody))
		return nil, AppError{Operation: "FindPlaceApi", Err: err}
	}

	//Decoding the response body to a type GeocodeResponse for further processing.
	err = json.Unmarshal(respBody, &findPlacesResp)
	if err != nil {
		log.Error("Failed to unmarshal the response body with error =>", err.Error())
		return nil, AppError{Operation: "FindPlaceApi", Err: err}
	}

	//Check if the findPlacesResp.Status is other than "OK" or "ZERO_RESULTS", then the request failed and so log the error message and return
	if findPlacesResp.Status != "OK" && findPlacesResp.Status != "ZERO_RESULTS" {
		log.Errorf("Failed to get places with with status: %v", findPlacesResp.Status)
		return
	}

	if findPlacesResp.Status == "ZERO_RESULTS" {
		log.Infof("The findPlaces request was successful, but returned no results, due to a non-existent address")
		return
	}

	return
}
