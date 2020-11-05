package external

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/pmadhvi/tech-test/bike-locator-api/models"
	log "github.com/sirupsen/logrus"
)

//TextSearchRequest request parameters.
type TextSearchRequest struct {
	Query     string
	APIKey    string
	Radius    string
	Region    string
	PlaceType string
}

//TextSearchResponse response structure.
type TextSearchResponse struct {
	Places PlaceList `json:"results"`
	Status string    `json:"status"`
}

//PlaceList is slice of Place
type PlaceList []Place

//Place structure
type Place struct {
	Address string `json:"formatted_address"`
	Name    string `json:"name"`
}

//TextSearchAPI as a method with Consumer as receiver and TextSearchRequest as parameter and returns response of type TextSearchResponse and err of type error.
func (c Consumer) TextSearchAPI(req TextSearchRequest) (models.BikeStores, error) {

	//Validating required request parametersand APIKey.
	err := validateTextSearchRequestParams(req)
	if err != nil {
		return nil, err
	}

	//Composing reqURL from Host and Path of Consumer struct.
	reqURL, err := c.composeTextSearchReqURL(req)
	if err != nil {
		return nil, err
	}

	//RunHTTP performs the http.Get under the hood and get's the response from server.
	var resp *http.Response
	resp, err = RunHTTP(reqURL.String())
	if err != nil {
		log.Errorf("Failed to get places with error => %v\n", err.Error())
		return nil, AppError{Operation: "TextSearchApi", Err: err}
	}

	respBody, err := checkResponseBody(resp)
	if err != nil {
		return nil, AppError{Operation: "TextSearchApi", Err: err}
	}

	//Decoding the response body to a type TextSearchResponse for further processing.
	var TextSearchsResp TextSearchResponse
	err = json.Unmarshal(respBody, &TextSearchsResp)
	if err != nil {
		log.Error("Failed to unmarshal the response body with error =>", err.Error())
		return nil, AppError{Operation: "TextSearchApi", Err: err}
	}

	//TODO: remove these prints
	fmt.Println("TextSearchsResp ====>", TextSearchsResp)

	//Check if the TextSearchsResp.Status is other than "OK" or "ZERO_RESULTS", then the request failed and so log the error message and return
	if TextSearchsResp.Status != "OK" && TextSearchsResp.Status != "ZERO_RESULTS" {
		log.Errorf("Failed to get places with with status: %v", TextSearchsResp.Status)

		message := "Request failed to get places"
		err = fmt.Errorf("TextSearchResponse status_code: %v and error_message: %v", TextSearchsResp.Status, message)
		return nil, AppError{Operation: "TextSearchApi", Err: err}
	}

	if TextSearchsResp.Status == "ZERO_RESULTS" {
		message := "The request was successful, but returned no results, due to a non-existent address"
		log.Info(message)
		err = fmt.Errorf("TextSearchResponse status_code: %v and message: %v", TextSearchsResp.Status, message)
		return nil, AppError{Operation: "TextSearchApi", Err: err}
	}

	//Fetching the response and putting it to a type of models.BikeStores consisting of []model.BikeStore
	bikeStores := make([]models.BikeStore, len(TextSearchsResp.Places))

	if len(TextSearchsResp.Places) >= 1 {
		for i, place := range TextSearchsResp.Places {
			bikeStores[i] = models.BikeStore{
				StoreName:    place.Name,
				StoreAddress: place.Address,
			}
		}
	}
	return bikeStores, nil
}

//validateTextSearchRequestParams validates required request parameters and APIKey.
func validateTextSearchRequestParams(req TextSearchRequest) (err error) {
	if req.Query == "" {
		message := "Query request parameter missing"
		log.Error(message)
		err := fmt.Errorf("status_code: %d and error_message: %s", 400, message)
		return AppError{Operation: "TextSearchApi", Err: err}
	}

	if req.PlaceType == "" {
		message := "PlaceType request parameter missing"
		log.Error(message)
		err := fmt.Errorf("status_code: %d and error_message: %s", 400, message)
		return AppError{Operation: "TextSearchApi", Err: err}
	}

	if req.APIKey == "" {
		message := "APIKey is missing"
		log.Error(message)
		err := fmt.Errorf("status_code: %d and error_message: %s", 400, message)
		return AppError{Operation: "TextSearchApi", Err: err}
	}
	return
}

func (c Consumer) composeTextSearchReqURL(req TextSearchRequest) (parsedReqURL *url.URL, err error) {
	reqURL := c.Host + c.Path

	//Parsing the reqURL to convert it to type *url.Url and checking for error if any.
	parsedReqURL, err = url.Parse(reqURL)
	if err != nil {
		message := "Incorrect URL"
		log.Error(message, err.Error())
		err = fmt.Errorf("status_code: %d and error_message: %s", 400, message)
		return nil, AppError{Operation: "TextSearchApi", Err: err}
	}

	//Adding request params to url and encoding it.
	params := url.Values{}
	params.Set("query", req.Query)

	params.Set("radius", req.Radius)
	params.Set("region", req.Region)
	params.Set("type", req.PlaceType)
	params.Set("key", req.APIKey)
	parsedReqURL.RawQuery = params.Encode()

	log.Infof("Request URL: %s", parsedReqURL.String())
	return
}

func checkResponseBody(resp *http.Response) (respBody []byte, err error) {
	//Defer the call to close the response body at the end of function execution.
	defer resp.Body.Close()

	//Start reading the response bosy into memory
	respBody, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Errorf("Failed to read the response body with error => %v\n", err.Error())
		return nil, AppError{Operation: "TextSearchApi", Err: err}
	}

	if resp.StatusCode == 404 {
		err = fmt.Errorf("status_code: %d and error_message: %s", resp.StatusCode, string(respBody))
		return nil, AppError{Operation: "TextSearchApi", Err: err}
	}

	if resp.StatusCode >= 400 {
		err = fmt.Errorf("status_code: %d and error_message: %s", resp.StatusCode, string(respBody))
		return nil, AppError{Operation: "TextSearchApi", Err: err}
	}
	return
}
