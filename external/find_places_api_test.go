package external_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/pmadhvi/tech-test/bike-locator-api/external"
)

var _ = Describe("FindPlacesAPI", func() {
	var (
		findPlaceReq FindPlaceRequest
		consumer     Consumer
		server       *httptest.Server
	)

	BeforeEach(func() {
		//Setting up the geocodemock server in order to test the FindPlacesAPI
		server = findPlacesMockServer()

		//Consumer takes in actual or test host, here host is localhost, so that the test does not hit the real google places server.
		consumer = Consumer{Host: server.URL, Path: "/maps/api/place/findplacefromtext/json"}

		//Pass in the parameters to findPlaceRequest struct.
		findPlaceReq.Input = "Toledo"
		findPlaceReq.InputType = "textquery"
		findPlaceReq.APIKey = "XYZ123"
		findPlaceReq.Fields = []string{"name, formatted_address"}
		findPlaceReq.LocationBiasType = "circle"
		findPlaceReq.LocationBiasRadius = 2000
		findPlaceReq.LocationBiasLat = 39.8628316
		findPlaceReq.LocationBiasLng = -4.027323099999999
	})

	AfterEach(func() {
		server.Close()
	})

	It("Gets the list of places", func() {
		//FindPlacesAPI being called to test the response.
		findPlaceResp, err := consumer.FindPlacesAPI(findPlaceReq)

		//Checking the expected response
		Expect(err).To(BeNil())
		Expect(len(findPlaceResp)).To(Equal(1))
		Expect(findPlaceResp[0].StoreAddress).To(Equal("Toledo, Spain"))
		Expect(findPlaceResp[0].StoreName).To(Equal("Museum of Contemporary Art Australia"))
	})

	It("Gets the list of places with empty input", func() {
		//Override input parameter
		findPlaceReq.Input = ""

		//FindPlacesAPI being called to test the response.
		findPlaceResp, err := consumer.FindPlacesAPI(findPlaceReq)

		//Checking the expected response
		Expect(err).ToNot(BeNil())
		Expect(findPlaceResp).To(BeNil())
	})

	It("Gets the list of places with empty input type", func() {
		//Override inputType parameter
		findPlaceReq.InputType = ""

		//FindPlacesAPI being called to test the response.
		findPlaceResp, err := consumer.FindPlacesAPI(findPlaceReq)

		//Checking the expected response
		Expect(err).ToNot(BeNil())
		Expect(findPlaceResp).To(BeNil())
	})

	It("Gets the list of places with empty APIKey", func() {
		//Override APIKey parameter
		findPlaceReq.APIKey = ""

		//FindPlacesAPI being called to test the response.
		findPlaceResp, err := consumer.FindPlacesAPI(findPlaceReq)

		//Checking the expected response
		Expect(err).ToNot(BeNil())
		Expect(findPlaceResp).To(BeNil())
	})

	It("Gets the list of places for an incorrect url", func() {
		//Overriding the consumer with incorrect url
		consumer = Consumer{Host: server.URL, Path: "/maps/api/places/"}

		//FindPlacesAPI being called to test the response.
		findPlaceResp, err := consumer.FindPlacesAPI(findPlaceReq)

		//Checking the expected response
		Expect(err).ToNot(BeNil())
		Expect(findPlaceResp).To(BeNil())
	})
})

func findPlacesMockServer() *httptest.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/maps/api/place/findplacefromtext/json", findPlacesMockApi)
	mux.HandleFunc("/maps/api/place/", findPlacesMockApiIncorrectURL)
	srv := httptest.NewServer(mux)

	return srv
}

func findPlacesMockApi(res http.ResponseWriter, r *http.Request) {
	file, err := os.Open("../testdata/findplaces_response.json")
	if err != nil {
		fmt.Errorf("Unable to open file: %v", err.Error())
	}
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Errorf("Unable to read the file content: %v", err.Error())
	}
	res.Write(bytes)
}

func findPlacesMockApiIncorrectURL(res http.ResponseWriter, r *http.Request) {
	res.WriteHeader(http.StatusNotFound)
	res.Write([]byte("URL not found"))
}
