package yelp

import (
	"fmt"

	"github.com/guregu/null"
)

// GeneralOptions includes a set of standard query options for using the search API.
// They are used along with a location based option to complete a search.
type GeneralOptions struct {
	Term           string     // Search term (e.g. "food", "restaurants"). If term isn’t included we search everything.
	Limit          null.Int   // Number of business results to return
	Offset         null.Int   // Offset the list of returned business results by this amount
	Sort           null.Int   // Sort mode: 0=Best matched (default), 1=Distance, 2=Highest Rated. If the mode is 1 or 2 a search may retrieve an additional 20 businesses past the initial limit of the first 20 results. This is done by specifying an offset and limit of 20. Sort by distance is only supported for a location or geographic search. The rating sort is not strictly sorted by the rating value, but by an adjusted rating value that takes into account the number of ratings, similar to a bayesian average. This is so a business with 1 rating of 5 stars doesn’t immediately jump to the top.
	CategoryFilter string     // Category to filter search results with. See the list of supported categories. The category filter can be a list of comma delimited categories. For example, 'bars,french' will filter by Bars and French. The category identifier should be used (for example 'discgolf', not 'Disc Golf').
	RadiusFilter   null.Float // Search radius in meters. If the value is too large, a AREA_TOO_LARGE error may be returned. The max value is 40000 meters (25 miles).
	DealsFilter    null.Bool  // Whether to exclusively search for businesses with deals
}

// getParameters will reflect over the values of the given
// struct, and provide a type appropriate set of querystring parameters
// that match the defined values.
func (o *GeneralOptions) getParameters() (params map[string]string, err error) {
	ps := make(map[string]string)
	if o.Term != "" {
		ps["term"] = o.Term
	}
	if o.Limit.Valid {
		ps["limit"] = fmt.Sprintf("%v", o.Limit.Int64)
	}
	if o.Offset.Valid {
		ps["offset"] = fmt.Sprintf("%v", o.Offset.Int64)
	}
	if o.Sort.Valid {
		ps["sort"] = fmt.Sprintf("%v", o.Sort.Int64)
	}
	if o.CategoryFilter != "" {
		ps["category_filter"] = o.CategoryFilter
	}
	if o.RadiusFilter.Valid {
		ps["radius_filter"] = fmt.Sprintf("%v", o.RadiusFilter.Float64)
	}
	if o.DealsFilter.Valid {
		ps["deals_filter"] = fmt.Sprintf("%v", o.DealsFilter.Bool)
	}
	return ps, nil
}
