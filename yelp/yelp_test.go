package yelp

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

const (
	CONTAINS_RESULTS        string = "The query returns at least one result."
	SHOULD_REQUIRE_LOCATION string = "The query should require a location."
)

// Check an error result for a value.  If present, fail the test with
// an error written to the console.
func check(t *testing.T, e error) {
	if e != nil {
		t.Error(e)
	}
}

// Provide a simple way to verify an assertion, and fail the test
// if that assertion fails.
func assert(t *testing.T, condition bool, assertion string) {
	if !condition {
		t.Errorf("Assertion failed: %v", assertion)
	}
}

// Creates a client with keys in a json file, making it possible to run the
// tests against the public Yelp API.
func getClient(t *testing.T) Client {
	data, err := ioutil.ReadFile("../config.json")
	check(t, err)
	var o AuthOptions
	err = json.Unmarshal(data, &o)
	check(t, err)
	client := CreateClient(o)
	return client
}

//
// TESTS
//

// Verify a simple search using a search term and location returns a set of results.
func TestSimpleSearch(t *testing.T) {
	client := getClient(t)
	result, err := client.DoSimpleSearch("coffee", "seattle")
	check(t, err)
	assert(t, len(result.Businesses) > 0, CONTAINS_RESULTS)
}

// Ensure validation for a missing location in the search.
func TestNoLocation(t *testing.T) {
	client := getClient(t)
	_, err := client.DoSimpleSearch("coffee", "")
	assert(t, err.Error() == ERROR_UNSPECIFIED_LOCATION, SHOULD_REQUIRE_LOCATION)
}

// Ensure you can query with no term defined and only a location.
func TestNoTerm(t *testing.T) {
	client := getClient(t)
	result, err := client.DoSimpleSearch("", "Seattle")
	check(t, err)
	assert(t, len(result.Businesses) > 0, CONTAINS_RESULTS)
}
