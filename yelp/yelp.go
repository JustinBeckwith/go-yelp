package yelp

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	ROOT_URI      = "http://api.yelp.com/"
	BUSINESS_AREA = "/v2/business"
	SEARCH_AREA   = "/v2/search"
)

type AuthOptions struct {
	ConsumerKey       string
	ConsumerSecret    string
	AccessToken       string
	AccessTokenSecret string
}

type Client struct {
	options AuthOptions
}

// Perform a simple search with a term and location.
func (client *Client) doSimpleSearch(term, location string) string {
	params := map[string]string{
		"term":     term,
		"location": location,
	}

	return client.makeRequest(SEARCH_AREA, "", params)
}

// Get a single business by name.
func (client *Client) getBusiness(name string) string {
	return client.makeRequest(BUSINESS_AREA, name, nil)
}

// Internal API used to make underlying requests to the Yelp API.
func (client *Client) makeRequest(area string, id string, params map[string]string) string {

	// get the base url
	queryUri, err := url.Parse(ROOT_URI)
	if err != nil {
		fmt.Println(err)
	}

	// add the type of request we're making (search|business)
	queryUri.Path = area

	if id != "" {
		queryUri.Path += "/" + id
	}

	// add querystring parameters
	values := queryUri.Query()
	for key, value := range params {
		values.Add(key, value)
	}
	queryUri.RawQuery = values.Encode()

	// set up OAUTH

	// make the actual request
	fmt.Printf("making http request to:\n%v\n", queryUri.String())
	resp, err := http.Get(queryUri.String())
	if err != nil {
		fmt.Println("unable to herp my derp: %v", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return string(body)
}

// Create a new yelp search client.  All search operations should go through this API.
func createClient(options AuthOptions) Client {
	return Client{options}
}
