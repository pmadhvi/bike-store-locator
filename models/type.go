package models

type BikeStore struct {
	StoreName    string `json:"store_name"`
	StoreAddress string `json:"store_address"`
}
type BikeStores []BikeStore

type Geocode struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
