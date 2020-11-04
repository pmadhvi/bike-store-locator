package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHealthAPI(t *testing.T) {
	var expected = "Application server is up and running!"
	t.Run("test healthapi with correct url ", func(t *testing.T) {
		//Create a request for checking the health of the application
		req := getHealthRequest("/bike-locator-api/healthz")

		//Create a ResponseRecorder to record the response
		response := httptest.NewRecorder()

		//Call the HealthApi with recorded response and request
		HealthAPI(response, req)

		//Check the response status and response body
		assert.Equal(t, http.StatusOK, response.Code)
		assert.Equal(t, expected, response.Body.String())
	})

	t.Run("test healthapi with wrong url ", func(t *testing.T) {
		//Create a request for checking the health of the application with wrong url
		req := getHealthRequest("/bike-locator-api/healthapi")

		//Create a ResponseRecorder to record the response
		response := httptest.NewRecorder()

		//call the HealthApi with recorded response and request
		HealthAPI(response, req)

		//Check the response status and response body
		assert.Equal(t, http.StatusOK, response.Code)
		assert.Equal(t, expected, response.Body.String())
	})
}

func getHealthRequest(path string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, path, nil)
	return req
}
