package yelp

import (
	"errors"
	"reflect"
)

// OptionProvider provides a base level interface which all search option interfaces implement.
// It provides information that allows for easily mapping to querystring parameters for the search query.
type OptionProvider interface {
	getParameters() (params map[string]string, err error)
}

// SearchOptions are the top level search parameters used for performing searches.
// You can define multiple sets of options, and use them together. One (and only one) of
// LocationOptions, CoordinateOptions, or BoundOptions can be used at the same time.
type SearchOptions struct {
	GeneralOptions    *GeneralOptions    // standard general search options (filters, terms, etc)
	LocaleOptions     *LocaleOptions     // Results will be localized in the region format and language if supported.
	LocationOptions   *LocationOptions   // Use a location term and potentially coordinates to define the location
	CoordinateOptions *CoordinateOptions // Use coordinate options to define the location.
	BoundOptions      *BoundOptions      // Use bound options (an area) to define the location.
}

// Generate a map that contains the querystring parameters for
// all of the defined options.
func (o *SearchOptions) getParameters() (params map[string]string, err error) {

	// ensure only one loc option provider is being used
	locOptionsCnt := 0
	if o.LocationOptions != nil {
		locOptionsCnt++
	}
	if o.CoordinateOptions != nil {
		locOptionsCnt++
	}
	if o.BoundOptions != nil {
		locOptionsCnt++
	}

	if locOptionsCnt == 0 {
		return params, errors.New("a single location search options type (Location, Coordinate, Bound) must be used")
	}
	if locOptionsCnt > 1 {
		return params, errors.New("only a single location search options type (Location, Coordinate, Bound) can be used at a time")
	}

	// create an empty map of options
	params = make(map[string]string)

	// reflect over the properties in o, adding parameters to the global map
	val := reflect.ValueOf(o).Elem()
	for i := 0; i < val.NumField(); i++ {
		if !val.Field(i).IsNil() {
			o := val.Field(i).Interface().(OptionProvider)
			fieldParams, err := o.getParameters()
			if err != nil {
				return params, err
			}
			for k, v := range fieldParams {
				params[k] = v
			}
		}
	}
	return params, nil
}
