package external

import (
	"fmt"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

//Consumer is used to differnetiate between actual server call or mock server call
type Consumer struct {
	Host string
	Path string
}

//AppError is a custom error type to return status code and message
type AppError struct {
	Operation string
	Err       error
}

//RunHTTP under the hood performs http.Get call and returns response or error
func RunHTTP(url string) (resp *http.Response, err error) {
	client := http.Client{
		Timeout: time.Duration(5 * time.Second),
	}
	start := time.Now()
	if resp, err = client.Get(url); err != nil {
		log.Error("Request failed with error\n", err.Error())
		return
	}

	finish := time.Since(start)
	log.Infof("The total time to serve the request from Google api: %v", finish.Seconds())

	return
}

func (e AppError) Error() string {
	return fmt.Sprintf("Error in %v with error details => %+v ", e.Operation, e.Err.Error())
}
