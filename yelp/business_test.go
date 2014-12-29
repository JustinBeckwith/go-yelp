package yelp

import (
	"testing"
)

/**
 * Perform a simple search for a business by name.
 */
func TestBusinessSearch(t *testing.T) {
	client := getClient(t)
	result, err := client.getBusiness("yelp-san-francisco")
	check(t, err)
	assert(t, result.Name != "", CONTAINS_RESULTS)
}
