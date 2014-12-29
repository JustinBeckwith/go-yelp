package yelp

import (
	"testing"
)

/**
 * Verify search with location and search term
 */
func TestGeneralOptions(t *testing.T) {
	client := getClient()
	options := SearchOptions{
		GeneralOptions: &GeneralOptions{
			Term: "coffee",
		},
	}
	result, err := client.doSearch(options)
	check(err)
	if len(result.Businesses) == 0 {
		t.Error("the query returned no results")
	}
}
