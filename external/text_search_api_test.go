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

var _ = Describe("TextSearchAPI", func() {
	var (
		TextSearchReq TextSearchRequest
		consumer      Consumer
		server        *httptest.Server
	)

	BeforeEach(func() {
		//Setting up the geocodemock server in order to test the TextSearchAPI
		server = TextSearchMockServer()

		//Consumer takes in actual or test host, here host is localhost, so that the test does not hit the real google places server.
		consumer = Consumer{Host: server.URL, Path: "/maps/api/place/textsearch/json"}

		//Pass in the parameters to TextSearchRequest struct.
		TextSearchReq.Query = "restaurants in Sydney"
		TextSearchReq.APIKey = "XYZ123"
		TextSearchReq.Radius = "2000"
		TextSearchReq.Region = "au"
		TextSearchReq.PlaceType = "restaurant"
	})

	AfterEach(func() {
		server.Close()
	})

	It("Gets the list of bike stores", func() {
		//TextSearchAPI being called to test the response.
		TextSearchResp, err := consumer.TextSearchAPI(TextSearchReq)

		//Checking the expected response
		Expect(err).To(BeNil())
		Expect(len(TextSearchResp)).To(Equal(4))
		Expect(TextSearchResp[0].StoreAddress).To(Equal("Pyrmont Bay Wharf Darling Dr, Sydney"))
		Expect(TextSearchResp[0].StoreName).To(Equal("Rhythmboat Cruises"))
		Expect(TextSearchResp[1].StoreAddress).To(Equal("Australia"))
		Expect(TextSearchResp[1].StoreName).To(Equal("Private Charter Sydney Habour Cruise"))
		Expect(TextSearchResp[2].StoreAddress).To(Equal("37 Bank St, Pyrmont"))
		Expect(TextSearchResp[2].StoreName).To(Equal("Bucks Party Cruise"))
		Expect(TextSearchResp[3].StoreAddress).To(Equal("32 The Promenade, King Street Wharf 5, Sydney"))
		Expect(TextSearchResp[3].StoreName).To(Equal("Australian Cruise Group"))
	})

	It("Gets the list of bike stores with empty query", func() {
		//Override input parameter
		TextSearchReq.Query = ""

		//TextSearchAPI being called to test the response.
		TextSearchResp, err := consumer.TextSearchAPI(TextSearchReq)

		//Checking the expected response
		Expect(err).ToNot(BeNil())
		Expect(TextSearchResp).To(BeNil())
	})

	It("Gets the list of bike stores with empty APIKey", func() {
		//Override APIKey parameter
		TextSearchReq.APIKey = ""

		//TextSearchAPI being called to test the response.
		TextSearchResp, err := consumer.TextSearchAPI(TextSearchReq)

		//Checking the expected response
		Expect(err).ToNot(BeNil())
		Expect(TextSearchResp).To(BeNil())
	})

	It("Gets the list of places for an incorrect url", func() {
		//Overriding the consumer with incorrect url
		consumer = Consumer{Host: server.URL, Path: "/maps/api/places/xxxx"}

		//TextSearchAPI being called to test the response.
		TextSearchResp, err := consumer.TextSearchAPI(TextSearchReq)

		//Checking the expected response
		Expect(err).ToNot(BeNil())
		Expect(TextSearchResp).To(BeNil())
	})
})

func TextSearchMockServer() *httptest.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/maps/api/place/textsearch/json", TextSearchMockApi)
	mux.HandleFunc("/maps/api/places/xxxx", TextSearchMockApiIncorrectURL)
	srv := httptest.NewServer(mux)

	return srv
}

func TextSearchMockApi(res http.ResponseWriter, r *http.Request) {
	file, err := os.Open("../testdata/text_search_response.json")
	if err != nil {
		fmt.Errorf("Unable to open file: %v", err.Error())
	}
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Errorf("Unable to read the file content: %v", err.Error())
	}
	res.Write(bytes)
}

func TextSearchMockApiIncorrectURL(res http.ResponseWriter, r *http.Request) {
	res.WriteHeader(http.StatusNotFound)
	res.Write([]byte("URL not found"))
}
