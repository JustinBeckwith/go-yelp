package yelp

import (
	"testing"

	"github.com/guregu/null"
)

// TestCoordinateOptions will check using location options with bounding coordinates.
func TestCoordinateOptions(t *testing.T) {
	client := getClient(t)
	options := SearchOptions{
		CoordinateOptions: &CoordinateOptions{
			null.FloatFrom(37.9),
			null.FloatFrom(-122.5),
			null.FloatFromPtr(nil),
			null.FloatFromPtr(nil),
			null.FloatFromPtr(nil),
		},
	}
	result, err := client.DoSearch(options)
	check(t, err)
	assert(t, len(result.Businesses) > 0, containsResults)
}
