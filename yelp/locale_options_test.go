package yelp

import (
	"testing"
)

// TestLocaleOptions will verify doing a search that includes locale options.
func TestLocaleOptions(t *testing.T) {
	client := getClient(t)
	options := SearchOptions{
		GeneralOptions: &GeneralOptions{
			Term: "coffee",
		},
		LocationOptions: &LocationOptions{
			Location: "seattle",
		},
		LocaleOptions: &LocaleOptions{
			cc:   "US",
			lang: "en",
		},
	}
	result, err := client.DoSearch(options)
	check(t, err)
	assert(t, len(result.Businesses) > 0, containsResults)
}
