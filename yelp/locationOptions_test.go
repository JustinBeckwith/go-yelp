package yelp

import (
	"testing"
)

/**
 * Check using location options with bounding coordinates
 */
func TestLocationOptions(t *testing.T) {
	client := getClient()
	options := SearchOptions{
		GeneralOptions: &GeneralOptions{
			Term: "coffee",
		},
		LocationOptions: &LocationOptions{
			Location: "seattle",
		},
	}

	result, err := client.doSearch(options)
	check(err)
	if len(result.Businesses) == 0 {
		t.Error("the query returned no results")
	}
}
