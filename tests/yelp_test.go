package yelp

import (
	"fmt"
	"testing"
)

func TestMakeRequest(t *testing.T) {
	result := makeRequest(SEARCH_AREA, "", map[string]string{
		"term":     "coffee",
		"location": "seattle",
	})
	fmt.Println(result)
}
