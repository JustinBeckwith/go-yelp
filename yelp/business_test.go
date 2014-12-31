package yelp

import (
	"testing"
)

// Perform a simple search for a business by name.
func TestBusinessSearch(t *testing.T) {
	client := getClient(t)
	result, err := client.GetBusiness("yelp-san-francisco")
	check(t, err)
	assert(t, result.Name != "", CONTAINS_RESULTS)
}

// Verify searching for a non-existent business throws the right error.
func TestNonExistingBusinessSearch(t *testing.T) {
	client := getClient(t)
	_, err := client.GetBusiness("place-that-doesnt-exist")
	assert(t, err.Error() == ERROR_BUSINESS_NOT_FOUND, "Searching for a non-existent business should return a 404 error")
}
