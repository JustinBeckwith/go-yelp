# go-yelp
go-yelp is a #golang wrapper for the Yelp REST API. It lets you do all kinds of interesting things like searching for businesses, getting user comments and ratings, and handling common errors. The library is written Go.

[![GoDoc](https://godoc.org/github.com/JustinBeckwith/go-yelp/yelp?status.svg)](https://godoc.org/github.com/JustinBeckwith/go-yelp/yelp)

[![Build Status](https://travis-ci.org/JustinBeckwith/go-yelp.svg?branch=master)](https://travis-ci.org/JustinBeckwith/go-yelp)

[![Coverage Status](https://coveralls.io/repos/JustinBeckwith/go-yelp/badge.png)](https://coveralls.io/r/JustinBeckwith/go-yelp)

For more information, visit the [Yelp REST API](http://www.yelp.com/developers/documentation/v2/overview).

## Getting Started
To install go-yelp, just use the `go get` command:

```sh
go get github.com/JustinBeckwith/go-yelp/yelp
```

When you're ready to start using the API, import the reference:

```go
import "github.com/JustinBeckwith/go-yelp/yelp"
```

### Authentication

All searches are performed through a client. To create a new client, you need provide a set of access keys necessary to use the V2 Yelp API. You can sign up for a Yelp developer account, and access your keys here:

[Yelp | Manage Keys](http://www.yelp.com/developers/manage_api_keys)

Keep these keys safe! There are a variety of ways to store them. I chose to store them in a config.json file which is not checked into the repository. To run the tests, you can create your own `config.json` file:

```json
{
	"ConsumerKey": "MY_CONSUMER_KEY",
	"ConsumerSecret":	"MY_CONSUMER_SECRET",
	"AccessToken":	"MY_ACCESS_TOKEN",
	"AccessTokenSecret":	"MY_ACCESS_TOKEN_SECRET"
}
```

### The Search API

The simple search API enables searching for businesses with a term and a location (ex: coffee, Seattle). After you have your keys, create a client, and make a simple query:

```go
import "github.com/JustinBeckwith/go-yelp/yelp"

client := yelp.New(options)
result, err := client.DoSimpleSearch("coffee", "seattle")
```

For more complex searches, the `DoSearch` method allows for searching based on a combination of general search criteria, and advanced location options:

```go
// Build an advanced set of search criteria that include 
// general options, and location specific options.
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

// Perform the search using the search options
result, err := client.DoSearch(options)
```

### The Business API
To directly search for a business by name, use the `client.GetBusiness(...)` method on the client:

```go
client := yelp.New(options)
result, err := client.GetBusiness("yelp-san-francisco")
```


## License
This library is distributed under the [MIT License](http://opensource.org/licenses/MIT) found in the LICENSE file.


## Questions?
Feel free to submit an issue on the repository, or find me at [@JustinBeckwith](http://twitter.com/JustinBeckwith)
