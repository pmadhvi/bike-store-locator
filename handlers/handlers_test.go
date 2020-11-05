package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHealthAPI(t *testing.T) {
	var expected = "{\"message\": \"alive\"}"
	var expectedNotFound = "{\"message\": \"not found\"}"
	t.Run("test healthapi with correct url ", func(t *testing.T) {
		//Create a request for checking the health of the application
		req := getRequest("/bikestoresapi/health")

		//Create a ResponseRecorder to record the response
		response := httptest.NewRecorder()

		//Calls the Healthhandler with recorded response and request
		HealthHandler(response, req)

		//Check the response status and response body
		assert.Equal(t, http.StatusOK, response.Code)
		assert.Equal(t, expected, response.Body.String())
	})

	t.Run("test healthapi with wrong url ", func(t *testing.T) {
		//Create a request for checking the health of the application with wrong url
		req := getRequest("/bikestoresapi/healthzz")

		//Create a ResponseRecorder to record the response
		response := httptest.NewRecorder()

		//calls the NotFoundHandler with recorded response and request
		NotFoundHandler(response, req)

		//Check the response status and response body
		assert.Equal(t, http.StatusNotFound, response.Code)
		assert.Equal(t, expectedNotFound, response.Body.String())
	})

	t.Run("test notfound ", func(t *testing.T) {
		//Create a request for checking the health of the application with wrong url
		req := getRequest("/")

		//Create a ResponseRecorder to record the response
		response := httptest.NewRecorder()

		//calls the NotFoundHandler with recorded response and request
		NotFoundHandler(response, req)

		//Check the response status and response body
		assert.Equal(t, http.StatusNotFound, response.Code)
		assert.Equal(t, expectedNotFound, response.Body.String())
	})
}

func getRequest(path string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, path, nil)
	return req
}
