package yelp

import (
	"encoding/json"
	"fmt"
	"github.com/mrjones/oauth"
	"io/ioutil"
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

/**
 * doSimpleSearch
 * Perform a simple search with a term and location.
 */
func (client *Client) doSimpleSearch(term, location string) (result SearchResult, err error) {
	// params := map[string]string{
	// 	"term":     term,
	// 	"location": location,
	// }
	// rawResult, err := client.makeRequest(SEARCH_AREA, "", params)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return SearchResult{}, err
	// }
	rawResult, err := ioutil.ReadFile("results.txt")
	fmt.Println(string(rawResult))
	err = json.Unmarshal(rawResult, &result)
	if err != nil {
		fmt.Println(err)
		return SearchResult{}, err
	}
	return result, nil
}

/**
 * getBusiness
 * Get a single business by name.
 */
func (client *Client) getBusiness(name string) (result Business, err error) {
	rawResult, err := client.makeRequest(BUSINESS_AREA, name, nil)
	if err != nil {
		fmt.Println(err)
		return Business{}, err
	}

	err = json.Unmarshal(rawResult, &result)
	if err != nil {
		fmt.Println(err)
		return Business{}, err
	}
	return result, nil
}

// Internal API used to make underlying requests to the Yelp API.
func (client *Client) makeRequest(area string, id string, params map[string]string) (result []byte, err error) {

	// get the base url
	queryUri, err := url.Parse(ROOT_URI)
	if err != nil {
		fmt.Println(err)
		return nil, err
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
		fmt.Println(err)
		return nil, err
	}
	defer response.Body.Close()

	bits, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return bits, nil

}

// Create a new yelp search client.  All search operations should go through this API.
func createClient(options AuthOptions) Client {
	return Client{options}
}
