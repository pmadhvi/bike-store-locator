package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetBikeStoresAPI(t *testing.T) {
	//os.Setenv("PORT", "1234")
	var expected, _ = readResponse("../testdata/bike_stores_response.json")

	t.Run("returns list of bike stores with correct input and request url", func(t *testing.T) {
		//Mock handler for BikeStoresHandler
		server := getBikeStoresMockhandler()
		defer server.Close()
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

	t.Run("returns error when called with wrong request url", func(t *testing.T) {
		//Mock handler for BikeStoresHandler
		server := getBikeStoresIncorrectURLMockhandler()
		defer server.Close()
		res, err := http.Get(server.URL)
		if err != nil {
			log.Fatal(err)
		}

		respBody, err := ioutil.ReadAll(res.Body)
		res.Body.Close()

		//Check the response status and response body
		assert.Equal(t, http.StatusNotFound, res.StatusCode)
		assert.Equal(t, string("Incorrect URL"), string(respBody))
	})

	t.Run("returns error when called with wrong request params", func(t *testing.T) {
		//Mock handler for BikeStoresHandler
		server := getBikeStoresIncorrectRequestMockhandler()
		defer server.Close()
		res, err := http.Get(server.URL)
		if err != nil {
			log.Fatal(err)
		}

		respBody, err := ioutil.ReadAll(res.Body)
		res.Body.Close()

		//Check the response status and response body
		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
		assert.Equal(t, string("Incorrect request param"), string(respBody))
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

//Mock handler for GetBikeStoresHandler
func getBikeStoresIncorrectURLMockhandler() *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Incorrect URL"))
	}))
	return server
}

//Mock handler for GetBikeStoresHandler
func getBikeStoresIncorrectRequestMockhandler() *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Incorrect request param"))
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
