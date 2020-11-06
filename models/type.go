package models

//BikeStore struct
type BikeStore struct {
	StoreName    string `json:"store_name"`
	StoreAddress string `json:"store_address"`
}

//BikeStores list of BikeStore
type BikeStores []BikeStore
