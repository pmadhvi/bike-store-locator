package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestGetBikeStoresAPI(t *testing.T) {
	os.Setenv("PORT", "1234")
	var expected, _ = readResponse("../testdata/bike_stores_response.json")

	//Mock handler for BikeStoresHandler
	server := getBikeStoresMockhandler()
	defer server.Close()

	t.Run("returns list of bike stores", func(t *testing.T) {
		res, err := http.Get(server.URL)
		if err != nil {
			log.Fatal(err)
		}

		stores, err := ioutil.ReadAll(res.Body)
		res.Body.Close()

		//Check the response status and response body
		assert.Equal(t, http.StatusOK, res.StatusCode)
		assert.Equal(t, string(expected), string(stores))
	})
}

//Mock handler for GetBikeStoresHandler
func getBikeStoresMockhandler() *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data, err := readResponse("../testdata/bike_stores_response.json")
		if err != nil {
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(data))
	}))
	return server
}

//readResponse reads the data from filepath
func readResponse(filepath string) (data []byte, err error) {
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Errorf("Unable to open file: %v", err.Error())
	}

	data, err = ioutil.ReadAll(file)
	if err != nil {
		fmt.Errorf("Unable to read the file content: %v", err.Error())
	}
	return
}

//getBikeRequest helps to create a request url with radius
func getBikeRequest(radius string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/bikestoresapi/radius/%s", radius), nil)

	//Since gorilla mux is being used for serving the request, that's why we need to set the request params in test using mux.SetURLVars, else request params will be not set and mux.Vars(req) returns map[].
	req = mux.SetURLVars(req, map[string]string{"radius": radius})
	return req
}
