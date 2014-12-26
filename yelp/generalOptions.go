package yelp

import (
	"fmt"
	"github.com/guregu/null"
)

type SearchOptions interface {
	GetParameters() (params map[string]string, err error)
}

/**
 * Standard general search parameters
 */
type GeneralOptions struct {
	SearchOptions
	Term            string     // Search term (e.g. "food", "restaurants"). If term isn’t included we search everything.
	Limit           null.Int   // Number of business results to return
	Offset          null.Int   // Offset the list of returned business results by this amount
	Sort            null.Int   // Sort mode: 0=Best matched (default), 1=Distance, 2=Highest Rated. If the mode is 1 or 2 a search may retrieve an additional 20 businesses past the initial limit of the first 20 results. This is done by specifying an offset and limit of 20. Sort by distance is only supported for a location or geographic search. The rating sort is not strictly sorted by the rating value, but by an adjusted rating value that takes into account the number of ratings, similar to a bayesian average. This is so a business with 1 rating of 5 stars doesn’t immediately jump to the top.
	Category_filter string     // Category to filter search results with. See the list of supported categories. The category filter can be a list of comma delimited categories. For example, 'bars,french' will filter by Bars and French. The category identifier should be used (for example 'discgolf', not 'Disc Golf').
	Radius_filter   null.Float // Search radius in meters. If the value is too large, a AREA_TOO_LARGE error may be returned. The max value is 40000 meters (25 miles).
	Deals_filter    null.Bool  // Whether to exclusively search for businesses with deals
	// Location        string            // [required]	Specifies the combination of "address, neighborhood, city, state or zip, optional country" to be used when searching for businesses.
	// Coordinates     CoordinateOptions // An optional latitude, longitude parameter can also be specified as a hint to the geocoder to disambiguate the location text.
	// Bounds          BoundOptions      // Location is specified by a bounding box, defined by a southwest latitude/longitude and a northeast latitude/longitude geographic coordinate.
	// Cc              string            // ISO 3166-1 alpha-2 country code. Default country to use when parsing the location field. United States = US, Canada = CA, United Kingdom = GB (not UK).
	// Lang            string            // ISO 639 language code (default=en). Reviews and snippets written in the specified language will be shown.
}

func (o *GeneralOptions) GetParameters() (params map[string]string, err error) {
	ps := make(map[string]string)
	if o.Term != "" {
		ps["term"] = o.Term
	}
	if o.Limit.Valid {
		ps["limit"] = string(o.Limit.Int64)
	}
	if o.Offset.Valid {
		ps["offset"] = string(o.Offset.Int64)
	}
	if o.Sort.Valid {
		ps["sort"] = string(o.Sort.Int64)
	}
	if o.Category_filter != "" {
		ps["category_filter"] = o.Category_filter
	}
	if o.Radius_filter.Valid {
		ps["radius_filter"] = fmt.Sprintf("%v", o.Radius_filter.Float64)
	}
	if o.Deals_filter.Valid {
		ps["deals_filter"] = fmt.Sprintf("%v", o.Deals_filter.Bool)
	}
	return ps, nil
}
