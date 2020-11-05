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

var _ = Describe("GeoCodingAPI ", func() {
	var (
		geocodeReq GeocodeRequest
		consumer   Consumer
		server     *httptest.Server
	)

	BeforeEach(func() {
		//Setting up the geocodemock server in order to test the GeoCodingApi
		server = geocodingMockServer()

		//Consumer takes in actual or test host, here host is localhost, so that the test does not hit the real google places server.
		consumer = Consumer{Host: server.URL, Path: "/maps/api/geocode/json"}

		//Pass in the parameters to GeocodeRequest struct.
		geocodeReq.Address = "Toledo"
		geocodeReq.APIKey = "XYZ123"
	})

	AfterEach(func() {
		server.Close()
	})

	It("Get geocode for a location and a region", func() {
		//Adding region parameter to request
		geocodeReq.Region = "es"

		//GeoCodingAPI being called to test the response.
		geocodeResp, err := consumer.GeoCodingAPI(geocodeReq)

		//Checking the expected response
		Expect(err).To(BeNil())
		Expect(geocodeResp.Longitude).To(Equal(-4.027323099999999))
		Expect(geocodeResp.Latitude).To(Equal(39.8628316))
	})

	It("Get geocode for a location only", func() {
		//GeoCodingAPI being called to test the response.
		geocodeResp, err := consumer.GeoCodingAPI(geocodeReq)

		//Checking the expected response
		Expect(err).To(BeNil())
		Expect(geocodeResp.Longitude).To(Equal(-4.027323099999999))
		Expect(geocodeResp.Latitude).To(Equal(39.8628316))
	})

	It("Get geocode for an empty location", func() {
		//Override address parameter
		geocodeReq.Address = ""

		//GeoCodingAPI being called to test the response.
		geocodeResp, err := consumer.GeoCodingAPI(geocodeReq)

		//Checking the expected response
		Expect(err).ToNot(BeNil())
		Expect(geocodeResp).To(BeNil())
	})

	It("Get geocode for missing APIKey", func() {
		//Override APIKey parameter
		geocodeReq.APIKey = ""

		//GeoCodingAPI being called to test the response.
		geocodeResp, err := consumer.GeoCodingAPI(geocodeReq)

		//Checking the expected response
		Expect(err).ToNot(BeNil())
		Expect(geocodeResp).To(BeNil())
	})

	It("Get geocode for an incorrect url", func() {
		//Overriding the consumer with incorrect url
		consumer = Consumer{Host: server.URL, Path: "/maps/api/"}

		//GeoCodingAPI being called to test the response.
		geocodeResp, err := consumer.GeoCodingAPI(geocodeReq)

		//Checking the expected response
		Expect(err).ToNot(BeNil())
		Expect(geocodeResp).To(BeNil())
	})
})

func geocodingMockServer() *httptest.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/maps/api/geocode/json", geocodeMockApi)
	mux.HandleFunc("/maps/api/", geocodeMockApiIncorrectURL)
	srv := httptest.NewServer(mux)

	return srv
}

func geocodeMockApi(res http.ResponseWriter, r *http.Request) {
	file, err := os.Open("../testdata/geocode_response.json")
	if err != nil {
		fmt.Errorf("Unable to open file: %v", err.Error())
	}
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Errorf("Unable to read the file content: %v", err.Error())
	}
	res.Write(bytes)
}

func geocodeMockApiIncorrectURL(res http.ResponseWriter, r *http.Request) {
	res.WriteHeader(http.StatusNotFound)
	res.Write([]byte("URL not found"))
}
