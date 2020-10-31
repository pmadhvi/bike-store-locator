package handlers

import (
	"io"
	"net/http"
)

//HealthAPI returns the health status of the application server.
func HealthAPI(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(200)
	io.WriteString(res, `{"alive": true}`)
}
