package models

//BikeStore struct
type BikeStore struct {
	StoreName    string `json:"name"`
	StoreAddress string `json:"address"`
}

//BikeStores list of BikeStore
type BikeStores []BikeStore

//Geocode struct
type Geocode struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
