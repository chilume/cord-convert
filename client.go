package convert

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	bgsAPIURL = "https://www.bgs.ac.uk/data/webservices/CoordConvert_LL_BNG.cfc"

//	methodLatLngToBNG = "LatLongtoBNG"
//	methodBNGToLatLng = "BNGtoLatLng"
//	timeout           = 10 * time.Second
)

// Client struct is used to create convert client to make conversion requests
type Client struct {
	// Base URL for API requests.
	BaseURL    *url.URL
	httpClient *http.Client
}

// Response is API response. This wraps the standard http.Response
// returned from API
/*
type Response struct {
	Response *http.Response
	Data     json.RawMessage `json:"data,omitempty"`
}
*/

//NewClient will return a new cleint, if no http client is provied a default is created.
func NewClient(client *http.Client) *Client {

	if client == nil {
		client = &http.Client{}
	}
	c := &Client{httpClient: client}

	c.BaseURL, _ = url.Parse(bgsAPIURL)

	return c
}

// NewRequest creates an API request. A relative URL can be provided in urlPath,
// in which case it is resolved relative to the BaseURL of the Client.
func (c *Client) NewRequest(method, urlPath string) (*http.Request, error) {

	u, err := c.BaseURL.Parse(urlPath)
	//fmt.Println(u.String())

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, u.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// Do sends an API request and returns the API response. The API response is
// JSON decoded and stored in the value pointed to by v, or returned as an
// error if an API error has occurred.
func (c *Client) Do(ctx context.Context, req *http.Request, v interface{}) (*http.Response, error) {
	req = req.WithContext(ctx)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if err != nil {
		return resp, err
	}
	// StatusNoContent, No Content.
	if resp.StatusCode != http.StatusNoContent {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil || body == nil {
			return resp, err
		}

		//fmt.Println("got response :", resp.StatusCode)

		if v != nil {
			err = json.Unmarshal(body, v)
			if err != nil {
				return &http.Response{}, err
			}
		}
	}
	return &http.Response{}, err
}
