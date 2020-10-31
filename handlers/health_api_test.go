package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthAPI(t *testing.T) {
	//Create a request for checking the health of the application
	req, err := http.NewRequest("GET", "/bike-locator-api/healthz", nil)
	if err != nil {
		t.Error(err)
	}
	//Create a ResponseRecorder to record the response
	recorder := httptest.NewRecorder()

	//pass the handler function and call serverHTTP with response and request
	handler := http.HandlerFunc(HealthAPI)
	handler.ServeHTTP(recorder, req)

	//Check the response status
	if status := recorder.Code; status != 200 {
		t.Errorf("Health Api returned wrong status code: %v and expected is %v\n", status, http.StatusOK)
	}

	//Check the response body
	expected := "Application server is up and running!"
	if recorder.Body.String() != expected {
		t.Errorf("Health Api returning incorrect response body: %s and expected is :%s\n", recorder.Body.String(), expected)
	}
}

func TestHealthAPIWithWrongURL(t *testing.T) {
	//Create a request for checking the health of the application with wrong url
	req, err := http.NewRequest("GET", "/bike-locator-api/healthapi", nil)
	if err != nil {
		t.Error(err)
	}
	//Create a ResponseRecorder to record the response
	recorder := httptest.NewRecorder()

	//pass the handler function and call serverHTTP with response and request
	handler := http.HandlerFunc(HealthAPI)
	handler.ServeHTTP(recorder, req)

	//Check the response status
	if status := recorder.Code; status != 200 {
		t.Errorf("Status code got: %v, expected status code : %v\n", status, http.StatusOK)
	}

	//Check the response body
	expected := "Application server is up and running!"
	if recorder.Body.String() != expected {
		t.Errorf("Response body got: %s, expected is: %s\n", recorder.Body.String(), expected)
	}
}
