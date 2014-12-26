package yelp

import (
	"fmt"
	"testing"
)

/**
 * Check using location options with bounding coordinates
 */
func TestCoordinateOptions(t *testing.T) {
	client := getClient()
	options := CoordinateOptions{37.9, -122.5, 37.788022, -122.399797}
	result, err := client.doSearch(options)
	check(err)
	if len(result.Businesses) == 0 {
		t.Error("the query returned no results")
	}
	fmt.Println(result)
}
