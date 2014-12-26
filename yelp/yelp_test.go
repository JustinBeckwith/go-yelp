package yelp

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// helper function used to read the config from a json file,
// and create the oauth options
func getClient() Client {
	data, err := ioutil.ReadFile("../config.json")
	check(err)
	var o AuthOptions
	err = json.Unmarshal(data, &o)
	check(err)
	client := createClient(o)
	return client
}

func TestSimpleSearch(t *testing.T) {
	client := getClient()
	result, err := client.doSimpleSearch("coffee", "seattle")
	check(err)
	if len(result.Businesses) == 0 {
		t.Error("the query returned no results")
	}
	fmt.Println(result)
}

// func TestBusinessSearch(t *testing.T) {
// 	client := getClient()
// 	result, err := client.getBusiness("yelp-san-francisco")
// 	check(err)
// 	// if len(result.businesses) == 0 {
// 	// 	t.Error("the query returned no results")
// 	// }
// 	fmt.Println(result)
// }
