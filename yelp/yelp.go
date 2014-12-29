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

	ERROR_UNSPECIFIED_LOCATION = "You must provide a location for the search."
	ERROR_BUSINESS_NOT_FOUND   = "The business could not be found."
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

/**
 * doSimpleSearch
 * Perform a simple search with a term and location.
 */
func (client *Client) doSimpleSearch(term, location string) (result SearchResult, err error) {

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

/**
 * doSearch
 * Perform a complex search with full search options.
 */
func (client *Client) doSearch(options SearchOptions) (result SearchResult, err error) {

	// get the options from the search provider
	params, err := options.GetParameters()
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

/**
 * getBusiness
 * Get a single business by name.
 */
func (client *Client) getBusiness(name string) (result Business, err error) {
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

/**
 * makeRequest
 * Internal API used to make underlying requests to the Yelp API.
 */
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
		client.options.ConsumerKey,
		client.options.ConsumerSecret,
		oauth.ServiceProvider{
			RequestTokenUrl:   "",
			AuthorizeTokenUrl: "",
			AccessTokenUrl:    "",
		})
	token := &oauth.AccessToken{
		client.options.AccessToken,
		client.options.AccessTokenSecret,
		make(map[string]string),
	}

	// make the request using the oauth lib
	response, err := c.Get(queryUri.String(), params, token)
	if err != nil {
		return nil, response.StatusCode, err
	}

	// ensure the request returned a 200
	if response.StatusCode != 200 {
		return nil, response.StatusCode, errors.New(response.Status)
	}

	fmt.Printf("%v\n", response.Request.URL.String())
	defer response.Body.Close()

	// read the body of the response
	bits, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, response.StatusCode, err
	}
	return bits, response.StatusCode, nil
}

// Create a new yelp search client.  All search operations should go through this API.
func createClient(options AuthOptions) Client {
	return Client{options}
}
