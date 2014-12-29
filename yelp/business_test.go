package yelp

import (
	"fmt"
	"testing"
)

/**
 * Perform a simple search for a business by name.
 */
func TestBusinessSearch(t *testing.T) {
	client := getClient()
	result, err := client.getBusiness("yelp-san-francisco")
	check(err)
	if result.Name != "" {
		t.Error("the query returned no results")
	}
	fmt.Println(result)
}
