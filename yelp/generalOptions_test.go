package yelp

import (
	"testing"
)

/**
 * Verify search with location and search term
 */
func TestGeneralOptions(t *testing.T) {
	client := getClient(t)
	options := SearchOptions{
		GeneralOptions: &GeneralOptions{
			Term: "coffee",
		},
		LocationOptions: &LocationOptions{
			Location: "seattle",
		},
	}
	result, err := client.doSearch(options)
	check(t, err)
	assert(t, len(result.Businesses) > 0, CONTAINS_RESULTS)
}
