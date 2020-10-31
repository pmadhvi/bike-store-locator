package external

import (
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

const (
	ApiKey = "XYZ"
)

func RunHTTP(url string) (resp *http.Response, err error) {
	client := http.Client{
		Timeout: time.Duration(5 * time.Second),
	}
	start := time.Now()
	if resp, err = client.Get(url); err != nil {
		log.Errorf("Request failed with error :\n", err.Error())
		return
	}
	defer resp.Body.Close()

	finish := time.Since(start)
	log.Infof("The total time to serve the request from Google places api:  %d\n", finish.Seconds())

	// var result map[string]interface{}
	// if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
	// 	log.Errorf("Failed to decode the response body with error :\n", err.Error())
	// 	return
	// }
	return
}

// func CheckError(err Error, message string) {
//     if err != nil {
// 		log.Errorf(message + " with error :", err.Error())
// 		return
// 	}
// }
