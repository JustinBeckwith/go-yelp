package yelp

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"testing"
)

const (
	containsResults       string = "The query returns at least one result."
	shouldRequireLocation string = "The query should require a location."
)

// Check an error result for a value.  If present, fail the test with
// an error written to the console.
func check(t *testing.T, e error) {
	if e != nil {
		t.Error(e)
	}
}

// assert provides a simple way to verify an assertion, and fail the test
// if that assertion fails.
func assert(t *testing.T, condition bool, assertion string) {
	if !condition {
		t.Errorf("Assertion failed: %v", assertion)
	}
}

// getClient creates a client with keys in a json file, making it possible to run the
// tests against the public Yelp API.
func getClient(t *testing.T) *Client {

	var o *AuthOptions

	// start by looking for the keys in config.json
	data, err := ioutil.ReadFile("../config.json")
	if err != nil {
		// if the file isn't there, check environment variables
		o = &AuthOptions{
			ConsumerKey:       os.Getenv("CONSUMER_KEY"),
			ConsumerSecret:    os.Getenv("CONSUMER_SECRET"),
			AccessToken:       os.Getenv("ACCESS_TOKEN"),
			AccessTokenSecret: os.Getenv("ACCESS_TOKEN_SECRET"),
		}
		if o.ConsumerKey == "" || o.ConsumerSecret == "" || o.AccessToken == "" || o.AccessTokenSecret == "" {
			check(t, errors.New("to run tests, keys must be provided either in a config.json file at the root of the repo, or in environment variables"))
		}
	} else {
		err = json.Unmarshal(data, &o)
		check(t, err)
	}
	client := New(o, nil)
	return client
}

//
// TESTS
//

// TestSimpleSearch verifies a simple search using a search term and location returns a set of results.
func TestSimpleSearch(t *testing.T) {
	client := getClient(t)
	result, err := client.DoSimpleSearch("coffee", "seattle")
	check(t, err)
	assert(t, len(result.Businesses) > 0, containsResults)

	// verify basic fields are returned
	assert(t, result.Region.Span.LatitudeDelta != 0, "latitude is returned")
	assert(t, result.Region.Span.LongitudeDelta != 0, "longitude is returned")

}

// TestNoLocation ensures validation for a missing location in the search.
func TestNoLocation(t *testing.T) {
	client := getClient(t)
	_, err := client.DoSimpleSearch("coffee", "")
	assert(t, err == errUnspecifiedLocation, shouldRequireLocation)
}

// TestNoTerm ensures you can query with no term defined and only a location.
func TestNoTerm(t *testing.T) {
	client := getClient(t)
	result, err := client.DoSimpleSearch("", "Seattle")
	check(t, err)
	assert(t, len(result.Businesses) > 0, containsResults)
}
