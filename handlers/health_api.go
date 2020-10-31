package handlers

import (
	"net/http"
)

//HealthAPI returns the health status of the application server.
func HealthAPI(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(200)
	res.Header().Set("Content-Type", "application/json")
	res.Write([]byte("Application server is up and running!"))
}
