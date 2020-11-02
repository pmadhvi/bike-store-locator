package external

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	log "github.com/sirupsen/logrus"
)

//GeocodeRequest request params structure.
type GeocodeRequest struct {
	Address string `json:"address"`
	Region  string `json:"status"`
	APIKey  string
}

//GeocodeResponse response structure.
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

//GeoCodingApi as a method with Consumer as receiver and GeocodeRequest as parameter and returns response of type GeocodeResponse and err of type error.
func (c Consumer) GeoCodingApi(req GeocodeRequest) (geocodeResp *GeocodeResponse, err error) {
	//Validating required request parameter and APIKey.
	if req.Address == "" {
		log.Error("Address request parameter missing")
		err = fmt.Errorf("status_code: %d and error_message: %s", 400, "Address request parameter missing")
		return nil, AppError{Operation: "GeoCodingApi", Err: err}
	}

	if req.APIKey == "" {
		log.Error("APIKey is missing")
		err = fmt.Errorf("status_code: %d and error_message: %s", 400, "APIKey is missing")
		return nil, AppError{Operation: "GeoCodingApi", Err: err}
	}

	//Composing reqURL from Host and Path of Consumer struct.
	reqURL := c.Host + c.Path

	//Parsing the reqURL to convert it to type *url.Url and checking for error if any.
	parsedReqURL, err := url.Parse(reqURL)
	if err != nil {
		log.Error("Incorrect URL:", err.Error())
		err = fmt.Errorf("status_code: %d and error_message: %s", 400, "Incorrect URL")
		return nil, AppError{Operation: "GeoCodingApi", Err: err}
	}

	//Adding request params to url and encoding it.
	params := url.Values{}
	params.Set("address", req.Address)
	params.Set("region", req.Region)
	params.Set("key", req.APIKey)
	parsedReqURL.RawQuery = params.Encode()

	log.Infof("Request URL: %s", parsedReqURL.String())

	var resp *http.Response
	//RunHTTP performs the http.Get under the hood and get's the response from server.
	resp, err = RunHTTP(parsedReqURL.String())
	if err != nil {
		log.Errorf("Failed to get the geocode for location: %s with error => %v\n", req.Address, err.Error())
		return nil, AppError{Operation: "GeoCodingApi", Err: err}
	}

	//Defer the call to close the response body at the end of function execution.
	defer resp.Body.Close()

	//Read the response body into memory
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Errorf("Failed to read the response body with error => %v\n", err.Error())
		return nil, AppError{Operation: "GeoCodingApi", Err: err}
	}

	if resp.StatusCode == 404 {
		err = fmt.Errorf("status_code: %d and error_message: %s", resp.StatusCode, string(respBody))
		return nil, AppError{Operation: "GeoCodingApi", Err: err}
	}

	if resp.StatusCode >= 400 {
		err = fmt.Errorf("status_code: %d and error_message: %s", resp.StatusCode, string(respBody))
		return nil, AppError{Operation: "GeoCodingApi", Err: err}
	}

	//Decoding the response body to a type GeocodeResponse for further processing.
	err = json.Unmarshal(respBody, &geocodeResp)
	if err != nil {
		log.Error("Failed to unmarshal the response body with error =>", err.Error())
		return nil, AppError{Operation: "GeoCodingApi", Err: err}
	}

	//Check if the geocodeResp.Status is other than "OK" or "ZERO_RESULTS", then the request failed and so log the error message and return
	if geocodeResp.Status != "OK" && geocodeResp.Status != "ZERO_RESULTS" {
		log.Errorf("Failed to get the geocode for location: %s with status: %v\n", req.Address, geocodeResp.Status)
		return
	}

	if geocodeResp.Status == "ZERO_RESULTS" {
		log.Info("The geocode request was successful, but returned no results, due to a non-existent address")
		return
	}

	return
}
