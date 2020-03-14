package convert

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	bgsAPIURL         = "https://www.bgs.ac.uk/data/webservices/CoordConvert_LL_BNG.cfc"
	methodLatLngToBNG = "LatLongtoBNG"
	methodBNGToLatLng = "BNGtoLatLng"
	timeout           = 10 * time.Second
)

//DegMinSec coordinates
type DegMinSec struct {
	Degres  int     `json:"DEGREES"`
	Seconds float64 `json:"SECONDS"`
	Minutes int     `json:"MINUTES"`
}

//Response contains the conversion response
type Response struct {
	DegMinSecLng DegMinSec `json:"DEGMINSECLNG"`
	Easting      float64   `json:"EASTING"`
	Longitude    float64   `json:"LONGITUDE"`
	Error        bool      `json:"ERROR"`
	DegMinSecLat DegMinSec `json:"DEGMINSECLAT"`
	Northing     float64   `json:"NORTHING"`
	Latitude     float64   `json:"LATITUDE"`
}

//LatLnglToBNG converts the latitude and longitude (WGS84) to British National Grid (BNG) cordinates
//formerly known as the National Grid Reference (NGR)
func LatLnglToBNG(latitude, longitude float64) (response Response, err error) {
	// Build the URL
	urlWithParams := fmt.Sprintf("%v?method=%v&lat=%v&lon=%v", bgsAPIURL, methodLatLngToBNG, latitude, longitude)
	response, err = get(urlWithParams)
	return
}

//BNGToLatLng converts the British National Grid (BNG) formerly known as the National Grid Reference (NGR) to
//latitude and longitude (WGS84) cordinates
func BNGToLatLng(easting, northing float64) (response Response, err error) {
	// Build the URL
	urlWithParams := fmt.Sprintf("%v?method=%v&easting=%v&northing=%v", bgsAPIURL, methodBNGToLatLng, easting, northing)
	response, err = get(urlWithParams)
	return
}

func get(url string) (response Response, err error) {
	http.DefaultClient.Timeout = timeout
	resp, err := http.Get(url)
	if err != nil {
		err = fmt.Errorf("error in get request: %v", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		err = fmt.Errorf("error in reading response from API: %v", err)
		return
	}

	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("Response status code %v is not ok", resp.StatusCode)
		return
	}
	var r Response
	err = json.Unmarshal(body, &r)
	if err != nil {
		err = fmt.Errorf("Unable to Unmarshal JSON response: %v", err)
		return
	}
	return r, nil
}
