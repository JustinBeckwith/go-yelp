package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/JustinBeckwith/go-yelp/yelp"
)

func main() {

	// get the keys either from config file
	var o yelp.AuthOptions
	data, err := ioutil.ReadFile("../../config.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, &o)
	if err != nil {
		panic(err)
	}

	// create a new yelp client with the auth keys
	client := yelp.New(o)

	// make a simple query
	term := os.Args[1]
	location := os.Args[2]
	results, err := client.DoSimpleSearch(term, location)
	if err != nil {
		panic(err)
	}

	// print the results
	fmt.Printf("\nFound a total of %v results for \"%v\" in \"%v\".\n", results.Total, term, location)
	fmt.Println("-----------------------------")
	for i := 0; i < len(results.Businesses); i++ {
		fmt.Printf("%v\t\t%v\n", results.Businesses[i].Name, results.Businesses[i].Rating)
	}
}
