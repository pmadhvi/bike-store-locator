Bike Store Locator API
==================================

[![GoDoc](https://godoc.org/googlemaps.github.io/maps?status.svg)](https://godoc.org/googlemaps.github.io/maps)

## Description
The Application returns the list of bike stores (name and address) near SergelTorg  and within radius of 2KM . Application under the hood makes a request to google text search api and fetches the response(list of places with extra information like name, formatted_address, and other details like latitude and longitude) and caches the response under the key "bikestores".

Request parameters of application:
-----------------------------------
radius: Defines the distance (in meters), example: 2000


Request parameters for google TextSearch Requests :
-----------------------------------------------------
The application under the hood calls **google places textsearch api** with following params.

* query : The text string on which to search , example: "bicycle_store near Sergeltorg"
* region:  The region code, example: "se"
* radius: Defines the distance (in meters) within which to bias place results, example: 2000(2km)
* type: Restricts the results to places matching the specified type, example: "bicycle_store"



The routes(rest api) for application includes:
-----------------------------------------------

* HealthApi: "/bikestoresapi/health"
* BikeStoresApi: "/bikestoresapi/radius/2000"


The URLS the application supports :
------------------------------------
* [HealthAPI](http://localhost:9000/bikestoresapi/health) 
* [BikeStoresAPI](http://localhost:9000/bikestoresapi/radius/2000)

## Application Uses

- Go 1.14 or later.
- A Google Places API key.

## Running the application

* To install the Bike store locator:

```bash
    go install
    bike-locator-api
```

* To build the Bike store locator:
```bash
    go build
    ./bike-locator-api
```

* To run test:
```bash
    go test -v ./...
```

## Developer Documentation

[Google Places - TextSearch API](https://developers.google.com/places/web-service/search#TextSearchRequests)