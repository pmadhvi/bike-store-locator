package models

//BikeStore struct
type BikeStore struct {
	StoreName    string `json:"name"`
	StoreAddress string `json:"address"`
}

//BikeStores list of BikeStore
type BikeStores []BikeStore
