package yelp

import (
	"testing"

	"github.com/guregu/null"
)

// TestLocationOptions will check using location options by term.
func TestLocationOptions(t *testing.T) {
	client := getClient(t)
	options := SearchOptions{
		GeneralOptions: &GeneralOptions{
			Term: "coffee",
		},
		LocationOptions: &LocationOptions{
			Location: "seattle",
		},
	}
	result, err := client.DoSearch(options)
	check(t, err)
	assert(t, len(result.Businesses) > 0, containsResults)
}

// TestLocationWithCoordinates will check using location options with bounding coordinates.
func TestLocationWithCoordinates(t *testing.T) {
	client := getClient(t)
	options := SearchOptions{
		GeneralOptions: &GeneralOptions{
			Term: "food",
		},
		LocationOptions: &LocationOptions{
			"bellevue",
			&CoordinateOptions{
				Latitude:  null.FloatFrom(37.788022),
				Longitude: null.FloatFrom(-122.399797),
			},
		},
	}
	result, err := client.DoSearch(options)
	check(t, err)
	assert(t, len(result.Businesses) > 0, containsResults)
}
