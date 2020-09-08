![Go](https://github.com/chilume/cord-convert/workflows/Go/badge.svg?branch=master)

# Cord-convert 

[Coordinate converter](https://www.bgs.ac.uk/data/webservices/convertForm.cfm#decimalLatLng) is a webservice to Convert British National Grid (BNG) formerly known as the National Grid Reference (NGR) to latitude and longitude (lat/long WGS84) or vice versa.

This is a Golang HTTP client libary for conversions.


# Usage

 Please look into the [example](example_test.go) code on how to use the client library.


# Testing 
  Run the below command to test.

  ```
  go test ./... -v -cover
  ```

## TODO

- Add support to convert degrees, seconds and minutes to BNG
