package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"fmt"
)

func TestGetBikeStoresAPI(t *testing.T) {
	t.Run("returns list of bike stores", func(t *testing.T) {
		//Create a request for getting list of bike stores
		req := getBikeRequest("es", "Todelo", 2000)
		
		//Create a ResponseRecorder to record the response
		response := httptest.NewRecorder()

		//Call the HealthApi with response and request
		GetBikeStoresAPI(response, req)

		//Check the response status
		assertResponseStatusCode(t, response.Code, http.StatusOK)

		//Check the response body
		expected := "Hello World"
		assertResponseBody(t, response.Body.String(), expected)
	})

	// t.Run("test healthapi with wrong url ", func(t *testing.T) {
	// 	//Create a request for checking the health of the application with wrong url
	// 	req, err := http.NewRequest("GET", "/bike-locator-api/healthapi", nil)
	// 	if err != nil {
	// 		t.Error(err)
	// 	}
	// 	//Create a ResponseRecorder to record the response
	// 	response := httptest.NewRecorder()

	// //call the HealthApi with response and request
	// 	HealthAPI(response, req)

	// 	//Check the response status
	// 	assertResponseStatusCode(t, response.Code, http.StatusOK)

	// 	//Check the response body
	// 	expected := "Application server is up and running!"
	// 	assertResponseBody(t, response.Body.String(), expected)
	// })
}

// func assertResponseBody(t *testing.T, got, want string) {
// 	t.Helper()
// 	if got != want {
// 			t.Errorf("Health Api returning incorrect response, got %q want %q", got, want)
// 	}
// }

// func assertResponseStatusCode(t *testing.T, got, want int) {
// 	t.Helper()
// 	if got != want {
// 			t.Errorf("Health Api returning incorrect status code, got %q want %q", got, want)
// 	}
// }

func getBikeRequest(region string, location string, radius int) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/bike-locator-api/region/%s/location/%s/radius/%d", region, location, radius), nil)
	return req
}