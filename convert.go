package convert

import (
	"context"
	"fmt"
	"net/http"
)

const (
	methodLatLngToBNG = "LatLongtoBNG"
	methodBNGToLatLng = "BNGtoLatLng"
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
func (c *Client) LatLnglToBNG(ctx context.Context, latitude, longitude float64) (*Response, *http.Response, error) {
	// Build the URL path
	urlPath := fmt.Sprintf("?method=%v&lat=%v&lon=%v", methodLatLngToBNG, latitude, longitude)

	req, err := c.NewRequest(http.MethodGet, urlPath)
	if err != nil {
		return nil, nil, err
	}
	r := new(Response)
	resp, err := c.Do(ctx, req, r)

	return r, resp, err

}

//BNGToLatLng converts the British National Grid (BNG) formerly known as the National Grid Reference (NGR) to
//latitude and longitude (WGS84) cordinates
func (c *Client) BNGToLatLng(ctx context.Context, easting, northing float64) (*Response, *http.Response, error) {
	// Build the URL path
	urlPath := fmt.Sprintf("?method=%v&easting=%v&northing=%v", methodBNGToLatLng, easting, northing)
	req, err := c.NewRequest(http.MethodGet, urlPath)
	if err != nil {
		return nil, nil, err
	}
	r := new(Response)
	resp, err := c.Do(ctx, req, r)

	return r, resp, err
}
