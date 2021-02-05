package models

//BikeStore struct
type BikeStore struct {
	StoreName    string `json:"store_name"`
	StoreAddress string `json:"store_address"`
}

//BikeStores list of BikeStore
type BikeStores []BikeStore

// BikeStores response payload
// swagger:response bikeStoresRes
type swaggBikeStoreRes struct {
	// in:body
	Body BikeStores
}

// Health response payload
// swagger:response healthRes
type swaggHealthRes struct {
	// in:body
	Body struct {
		Message string `json:"message"`
	}
}

// Error Bad Request
// swagger:response badReq
type swaggReqBadRequest struct {
	// in:body
	Body struct {
		// HTTP status code 400 -  Bad Request
		Code int `json:"code"`
	}
}

// Error Not Found
// swagger:response notFoundReq
type swaggReqNotFound struct {
	// in:body
	Body struct {
		// HTTP status code 404 -  Not Found
		Code int `json:"code"`
	}
}
