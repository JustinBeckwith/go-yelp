package yelp

import (
	"testing"
)

// TestBoundOptions will check using location options with bounding coordinates
func TestBoundOptions(t *testing.T) {
	client := getClient(t)
	options := SearchOptions{
		BoundOptions: &BoundOptions{37.9, -122.5, 37.788022, -122.399797},
	}
	result, err := client.DoSearch(options)
	check(t, err)
	assert(t, len(result.Businesses) > 0, containsResults)
}
