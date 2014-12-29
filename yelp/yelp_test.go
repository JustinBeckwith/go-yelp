package yelp

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

const (
	CONTAINS_RESULTS        string = "The query returns at least one result"
	SHOULD_REQUIRE_LOCATION string = "The query should require a location"
)

func check(t *testing.T, e error) {
	if e != nil {
		t.Error(e)
	}
}

func assert(t *testing.T, condition bool, assertion string) {
	if !condition {
		t.Errorf("Assertion failed: %v", assertion)
	}
}

// helper function used to read the config from a json file,
// and create the oauth options
func getClient(t *testing.T) Client {
	data, err := ioutil.ReadFile("../config.json")
	check(t, err)
	var o AuthOptions
	err = json.Unmarshal(data, &o)
	check(t, err)
	client := createClient(o)
	return client
}

/**
 * Verify a simple search using a search term and
 * location returns a set of results.
 */
func TestSimpleSearch(t *testing.T) {
	client := getClient(t)
	result, err := client.doSimpleSearch("coffee", "seattle")
	check(t, err)
	assert(t, len(result.Businesses) > 0, CONTAINS_RESULTS)
}

/**
 * Verify a simple search using a search term and
 * location returns a set of results.
 */
func TestNoLocation(t *testing.T) {
	client := getClient(t)
	result, err := client.doSimpleSearch("coffee", "")
	assert(t, err.Error() == ERROR_UNSPECIFIED_LOCATION, SHOULD_REQUIRE_LOCATION)
}
