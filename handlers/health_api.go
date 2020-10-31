package handlers

import (
	"net/http"
)

//HealthAPI returns the health status of the application server.
func HealthAPI(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(200)
	res.Write([]byte("Application server is up and running!"))
}
