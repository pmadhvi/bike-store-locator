Bike Store Locator API
==================================

[![GoDoc](https://godoc.org/googlemaps.github.io/maps?status.svg)](https://godoc.org/googlemaps.github.io/maps)

## Description
The Application returns the list of bike stores (name and address) near SergelTorg and within radius of 2KM .

The routes for application includes:
-------------------------------------

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

[Google Places API - TextSearch](https://developers.google.com/places/web-service/search#TextSearchRequests)