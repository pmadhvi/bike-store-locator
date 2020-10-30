package handlers

import (
	"io"
	"net/http"
)

func HealthApi(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(200)
	res.Header().Set("Content-Type", "application/json")
	io.WriteString(res, `{"alive": true}`)
}
