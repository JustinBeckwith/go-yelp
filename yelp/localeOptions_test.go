package yelp

import (
	"testing"
)

/**
 * Verify doing a search that includes locale options.
 */
func TestLocaleOptions(t *testing.T) {
	client := getClient()
	options := SearchOptions{
		GeneralOptions: &GeneralOptions{
			Term: "coffee",
		},
		LocaleOptions: &LocaleOptions{
			cc:   "US",
			lang: "en",
		},
	}
	result, err := client.doSearch(options)
	check(err)
	if len(result.Businesses) == 0 {
		t.Error("the query returned no results")
	}
}
