package yelp

import (
	"github.com/guregu/null"
	"testing"
)

/**
 * Check using location options with bounding coordinates
 */
func TestCoordinateOptions(t *testing.T) {
	client := getClient()
	options := CoordinateOptions{
		null.FloatFrom(37.9),
		null.FloatFrom(-122.5),
		null.FloatFromPtr(nil),
		null.FloatFromPtr(nil),
		null.FloatFromPtr(nil),
	}
	result, err := client.doSearch(options)
	check(err)
	if len(result.Businesses) == 0 {
		t.Error("the query returned no results")
	}
}
