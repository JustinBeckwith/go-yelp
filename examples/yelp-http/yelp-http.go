// This example demonstrates using querystring parameters to perform
// a simple query with the yelp API.
// Example url:  http://localhost:8000/?term=coffee&location=seattle
package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/JustinBeckwith/go-yelp/yelp"
)

func main() {
	http.HandleFunc("/", res)
	http.ListenAndServe(":8000", nil)
}

func res(w http.ResponseWriter, r *http.Request) {

	// get the keys either from config file
	options, err := getOptions(w)
	if err != nil {
		fmt.Println(err)
	}

	// create a new yelp client with the auth keys
	client := yelp.New(options)

	// make a simple query
	term := r.URL.Query().Get("term")
	location := r.URL.Query().Get("location")
	results, err := client.DoSimpleSearch(term, location)
	if err != nil {
		fmt.Println(err)
	}

	// print the results
	io.WriteString(w, fmt.Sprintf("<div>Found a total of %v results for \"%v\" in \"%v\".</div>", results.Total, term, location))
	io.WriteString(w, "<div>-----------------------------</div>")
	for i := 0; i < len(results.Businesses); i++ {
		io.WriteString(w, fmt.Sprintf("<div>%v, %v</div>", results.Businesses[i].Name, results.Businesses[i].Rating))
	}
}

// getOptions obtains the keys required to use the Yelp API from a config file
// or from environment variables.
func getOptions(w http.ResponseWriter) (options yelp.AuthOptions, err error) {

	var o yelp.AuthOptions

	// start by looking for the keys in config.json
	data, err := ioutil.ReadFile("../../config.json")
	if err != nil {
		// if the file isn't there, check environment variables
		o = yelp.AuthOptions{
			ConsumerKey:       os.Getenv("CONSUMER_KEY"),
			ConsumerSecret:    os.Getenv("CONSUMER_SECRET"),
			AccessToken:       os.Getenv("ACCESS_TOKEN"),
			AccessTokenSecret: os.Getenv("ACCESS_TOKEN_SECRET"),
		}
		if o.ConsumerKey == "" || o.ConsumerSecret == "" || o.AccessToken == "" || o.AccessTokenSecret == "" {
			return o, errors.New("to use the sample, keys must be provided either in a config.json file at the root of the repo, or in environment variables")
		}
	} else {
		err = json.Unmarshal(data, &o)
		return o, err
	}
	return o, nil
}
