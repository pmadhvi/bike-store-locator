package external_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/pmadhvi/tech-test/bike-locator-api/external"
)

//Setup the test suite for testing Geocoding Api external request
func TestGeoCodingApi(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Gecoding Suite")
}

var _ = Describe("External", func() {

	Describe("GeoCodingApi", func() {

		// BeforeEach(func() {

		// })
		// AfterEach(func() {

		// })
		It("Get geocode for a location and a region", func() {
			srv := geocodingMockServer()
			defer srv.Close()
			consumer := Consumer{Host: srv.URL, Path: "/maps/api/geocode/json"}
			var geocodeReq GeocodeRequest
			var geocodeResponse GeocodeResponse
			geocodeReq.Address = "Toledo"
			geocodeReq.Region = "es"
			geocodeResponse, err := consumer.GeoCodingApi(geocodeReq)

			fmt.Println("geocoderesp from test::", geocodeResponse)

			Expect(err).To(BeNil())
			Expect(geocodeResponse.Status).To(Equal("OK"))

			// resp := external.GeocodeResponse{
			// 	GeocodeResults: []GeocodeResult{
			// 		FormattedAddress: "Toledo, Spain",
			// 		Geometry: Geometry{
			// 			Location: LocationLatLng{Latitude: 39.8628316, Latitude: -4.027323099999999},
			// 		},
			// 	},
			// 	Status: "OK",
			// }
			// Expect(geocodeResponse).To(Equal(resp))
		})

	})
})

func geocodingMockServer() *httptest.Server {
	handler := http.NewServeMux()
	handler.HandleFunc("/maps/api/geocode/json", getGeocodeMockApi)

	srv := httptest.NewServer(handler)

	return srv
}

func getGeocodeMockApi(w http.ResponseWriter, r *http.Request) {
	response := `{
		"results" : [
			 {
					"address_components" : [
						 {
								"long_name" : "Toledo",
								"short_name" : "Toledo",
								"types" : [ "locality", "political" ]
						 },
						 {
								"long_name" : "Toledo",
								"short_name" : "TO",
								"types" : [ "administrative_area_level_2", "political" ]
						 },
						 {
								"long_name" : "Castile-La Mancha",
								"short_name" : "CM",
								"types" : [ "administrative_area_level_1", "political" ]
						 },
						 {
								"long_name" : "Spain",
								"short_name" : "ES",
								"types" : [ "country", "political" ]
						 }
					],
					"formatted_address" : "Toledo, Spain",
					"geometry" : {
						 "bounds" : {
								"northeast" : {
									 "lat" : 39.88605099999999,
									 "lng" : -3.9192423
								},
								"southwest" : {
									 "lat" : 39.8383676,
									 "lng" : -4.0796176
								}
						 },
						 "location" : {
								"lat" : 39.8628316,
								"lng" : -4.027323099999999
						 },
						 "location_type" : "APPROXIMATE",
						 "viewport" : {
								"northeast" : {
									 "lat" : 39.88605099999999,
									 "lng" : -3.9192423
								},
								"southwest" : {
									 "lat" : 39.8383676,
									 "lng" : -4.0796176
								}
						 }
					},
					"place_id" : "ChIJ8f21C60Lag0R_q11auhbf8Y",
					"types" : [ "locality", "political" ]
			 }
		],
		"status" : "OK"
 }`
	_, _ = w.Write([]byte(response))
}
