package yelp

import (
	"testing"
)

// TestBusinessSearch will perform a simple search for a business by name.
func TestBusinessSearch(t *testing.T) {
	client := getClient(t)
	result, err := client.GetBusiness("yelp-san-francisco")
	check(t, err)
	assert(t, result.Name != "", containsResults)
}

// TestNonExistingBusinessSearch will verify searching for a non-existent business throws the right error.
func TestNonExistingBusinessSearch(t *testing.T) {
	client := getClient(t)
	_, err := client.GetBusiness("place-that-doesnt-exist")
	assert(t, err == errBusinessNotFound, "Searching for a non-existent business should return a 404 error")
}
