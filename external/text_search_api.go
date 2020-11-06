package external

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/patrickmn/go-cache"
	"github.com/pmadhvi/tech-test/bike-locator-api/models"
	log "github.com/sirupsen/logrus"
)

const (
	cacheKey       = "bikestores"
	testSearchPath = "/maps/api/place/textsearch/json"
)

var (
	storeCache     = cache.New(5*time.Minute, 2*time.Minute)
	textSearchResp TextSearchResponse
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
func (c Consumer) TextSearchAPI(req TextSearchRequest) (bikeStores models.BikeStores, err error) {

	//Validating required request parametersand APIKey.
	err = validateTextSearchRequestParams(req)
	if err != nil {
		return nil, err
	}

	//Composing reqURL from Host and Path of Consumer struct.
	reqURL, err := c.composeTextSearchReqURL(req)
	if err != nil {
		return nil, err
	}

	//Check the cache first to find the list of stores
	bikeStores, found := readFromCache(cacheKey)
	if !found {
		//Since the data is not found in cache, an http request is made.
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
		err = json.Unmarshal(respBody, &textSearchResp)
		if err != nil {
			log.Error("Failed to unmarshal the response body with error =>", err.Error())
			return nil, AppError{Operation: "TextSearchApi", Err: err}
		}

		//validate the textSearchResp status
		err = validatesTextSearchResponseStatus(textSearchResp)
		if err != nil {
			return nil, err
		}

		//Fetching the response and putting it to a type of models.BikeStores consisting of []model.BikeStore
		bikeStores = make([]models.BikeStore, len(textSearchResp.Places))

		if len(textSearchResp.Places) >= 1 {
			for i, place := range textSearchResp.Places {
				bikeStores[i] = models.BikeStore{
					StoreName:    place.Name,
					StoreAddress: place.Address,
				}
			}
		}
		if set := setCache(cacheKey, bikeStores); !set {
			log.Error("Bikestores data did not cached")
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

//composeTextSearchReqURL composes the request URL.
func (c Consumer) composeTextSearchReqURL(req TextSearchRequest) (parsedReqURL *url.URL, err error) {
	//Check if request url path is correct
	if c.Path != testSearchPath {
		message := "Incorrect URL"
		log.Error(message)
		err = fmt.Errorf("status_code: %d and error_message: %s", 400, message)
		return nil, AppError{Operation: "TextSearchApi", Err: err}
	}

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

//checkResponseBody reads the response body in memeory and also checks the http status code of the response.
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

//setCache is used to set data with key in cache
func setCache(key string, data interface{}) bool {
	storeCache.Set(key, data, cache.DefaultExpiration)
	return true
}

//readFromCache is used to fetch data from cache
func readFromCache(key string) (bikeStores models.BikeStores, found bool) {
	data, found := storeCache.Get(key)
	if found {
		bikeStores = data.(models.BikeStores)
	}
	return bikeStores, found
}

func validatesTextSearchResponseStatus(textSearchResp TextSearchResponse) (err error) {
	//Check if the textSearchResp.Status is other than "OK" or "ZERO_RESULTS", then the request failed and so log the error message and return
	if textSearchResp.Status != "OK" && textSearchResp.Status != "ZERO_RESULTS" {
		log.Errorf("Failed to get places with with status: %v", textSearchResp.Status)
		message := "Request failed to get places"
		err = fmt.Errorf("TextSearchResponse status_code: %v and error_message: %v", textSearchResp.Status, message)
		return AppError{Operation: "TextSearchApi", Err: err}
	}

	if textSearchResp.Status == "ZERO_RESULTS" {
		message := "The request was successful, but returned no results, due to a non-existent address"
		log.Info(message)
		err = fmt.Errorf("TextSearchResponse status_code: %v and message: %v", textSearchResp.Status, message)
		return AppError{Operation: "TextSearchApi", Err: err}
	}

	if textSearchResp.Status == "OVER_QUERY_LIMIT" {
		message := "Going above your quota limits"
		log.Info(message)
		err = fmt.Errorf("TextSearchResponse status_code: %v and message: %v", textSearchResp.Status, message)
		return AppError{Operation: "TextSearchApi", Err: err}
	}
	return
}
