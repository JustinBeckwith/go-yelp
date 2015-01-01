package yelp

import (
	"testing"

	"github.com/guregu/null"
)

// TestGeneralOptions will verify search with location and search term.
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
	result, err := client.DoSearch(options)
	check(t, err)
	assert(t, len(result.Businesses) > 0, containsResults)
}

// TestUnescapedCharactersInGeneralOptions verify URL escaped characters do not cause search to fail.
func TestUnescapedCharactersInGeneralOptions(t *testing.T) {
	client := getClient(t)
	options := SearchOptions{
		GeneralOptions: &GeneralOptions{
			Term: "Frimark Keller & Associates",
		},
		LocationOptions: &LocationOptions{
			Location: "60173",
		},
	}
	result, err := client.DoSearch(options)
	check(t, err)
	assert(t, len(result.Businesses) > 0, containsResults)
}

// TestMultipleCategories will perform a search with multiple categories on the general options filter.
func TestMultipleCategories(t *testing.T) {
	client := getClient(t)
	options := SearchOptions{
		GeneralOptions: &GeneralOptions{
			CategoryFilter: "climbing,bowling",
		},
		LocationOptions: &LocationOptions{
			Location: "Seattle",
		},
	}
	result, err := client.DoSearch(options)
	check(t, err)
	assert(t, len(result.Businesses) > 0, containsResults)
}

// TestLimit verify the limit parameter works as expected.
func TestLimit(t *testing.T) {
	client := getClient(t)
	options := SearchOptions{
		GeneralOptions: &GeneralOptions{
			Term:  "Coffee",
			Limit: null.IntFrom(15),
		},
		LocationOptions: &LocationOptions{
			Location: "Seattle",
		},
	}
	result, err := client.DoSearch(options)
	check(t, err)
	assert(t, len(result.Businesses) == 15, "There should be 15 results.")
}
