// go-yelp provides a lightweight wrapper around the Yelp REST API.  It supports authentication with
// OAuth 1.0, the Search API, and the Business API.
package yelp

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/url"

	"github.com/mrjones/oauth"
)

const (
	ROOT_URI      = "http://api.yelp.com/"
	BUSINESS_AREA = "/v2/business"
	SEARCH_AREA   = "/v2/search"

	ERROR_UNSPECIFIED_LOCATION = "you must provide a location for the search"
	ERROR_BUSINESS_NOT_FOUND   = "the business could not be found"
)

// AuthOptions provide keys required for using the Yelp API.  Find more
// information here:  http://www.yelp.com/developers/documentation.
type AuthOptions struct {
	ConsumerKey       string // Consumer Key from the yelp API access site.
	ConsumerSecret    string // Consumer Secret from the yelp API access site.
	AccessToken       string // Token from the yelp API access site.
	AccessTokenSecret string // Token Secret from the yelp API access site.
}

// All searches are performed from an instance of a client.  It is the top level
// object used to perform a search or business query.  Client objects should be
// created through the createClient API.
type Client struct {
	Options AuthOptions
}

// Perform a simple search with a term and location.
func (client *Client) DoSimpleSearch(term, location string) (result SearchResult, err error) {

	// verify the term and location are not empty
	if location == "" {
		return SearchResult{}, errors.New(ERROR_UNSPECIFIED_LOCATION)
	}

	// set up the query options
	params := map[string]string{
		"term":     term,
		"location": location,
	}

	// perform the search request
	rawResult, _, err := client.makeRequest(SEARCH_AREA, "", params)
	if err != nil {
		return SearchResult{}, err
	}

	// convert the result from json
	err = json.Unmarshal(rawResult, &result)
	if err != nil {
		return SearchResult{}, err
	}
	return result, nil
}

// Perform a complex search with full search options.
func (client *Client) DoSearch(options SearchOptions) (result SearchResult, err error) {

	// get the options from the search provider
	params, err := options.getParameters()
	if err != nil {
		return SearchResult{}, err
	}

	// perform the search request
	rawResult, _, err := client.makeRequest(SEARCH_AREA, "", params)
	if err != nil {
		return SearchResult{}, err
	}

	// convert the result from json
	err = json.Unmarshal(rawResult, &result)
	if err != nil {
		return SearchResult{}, err
	}
	return result, nil
}

// Get a single business by name.
func (client *Client) GetBusiness(name string) (result Business, err error) {
	rawResult, statusCode, err := client.makeRequest(BUSINESS_AREA, name, nil)
	if err != nil {
		if statusCode == 404 {
			return Business{}, errors.New(ERROR_BUSINESS_NOT_FOUND)
		}
		return Business{}, err
	}

	err = json.Unmarshal(rawResult, &result)
	if err != nil {
		return Business{}, err
	}
	return result, nil
}

// Internal/private API used to make underlying requests to the Yelp API.
func (client *Client) makeRequest(area string, id string, params map[string]string) (result []byte, statusCode int, err error) {

	// get the base url
	queryUri, err := url.Parse(ROOT_URI)
	if err != nil {
		return nil, 0, err
	}

	// add the type of request we're making (search|business)
	queryUri.Path = area

	if id != "" {
		queryUri.Path += "/" + id
	}

	// set up OAUTH
	c := oauth.NewConsumer(
		client.Options.ConsumerKey,
		client.Options.ConsumerSecret,
		oauth.ServiceProvider{
			RequestTokenUrl:   "",
			AuthorizeTokenUrl: "",
			AccessTokenUrl:    "",
		})
	token := &oauth.AccessToken{
		client.Options.AccessToken,
		client.Options.AccessTokenSecret,
		make(map[string]string),
	}

	// make the request using the oauth lib
	response, err := c.Get(queryUri.String(), params, token)

	// always log the url, and close the request when done
	fmt.Printf("%v\n", response.Request.URL.String())
	defer response.Body.Close()

	if err != nil {
		return nil, response.StatusCode, err
	}

	// ensure the request returned a 200
	if response.StatusCode != 200 {
		return nil, response.StatusCode, errors.New(response.Status)
	}

	// read the body of the response
	bits, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, response.StatusCode, err
	}
	return bits, response.StatusCode, nil
}

// Create a new yelp search client.  All search operations should go through this API.
func CreateClient(options AuthOptions) Client {
	return Client{options}
}
