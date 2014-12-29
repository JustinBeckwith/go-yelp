package yelp

import (
	"errors"
)

type OptionProvider interface {
	GetParameters() (params map[string]string, err error)
}

/**
 * Top level search object used for doing searches.  You can define multiple
 * sets of options, and use them together.  Only one of LocationOptions,
 * CoordinateOptions, or BoundOptions can be used at the same time.
 */
type SearchOptions struct {
	GeneralOptions    *GeneralOptions    // standard general search options (filters, terms, etc)
	LocaleOptions     *LocaleOptions     // Results will be localized in the region format and language if supported.
	LocationOptions   *LocationOptions   // Use a location term and potentially coordinates to define the location
	CoordinateOptions *CoordinateOptions // Use coordinate options to define the location.
	BoundOptions      *BoundOptions      // Use bound options (an area) to define the location.
}

/**
 * Generate a map that contains the querystring parameters for
 * all of the defined options.
 */
func (o *SearchOptions) GetParameters() (params map[string]string, err error) {

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
	if locOptionsCnt > 1 {
		return params, errors.New("Only a single location search options type (Location, Coordinate, Bound) can be used at a time.")
	}

	// create an empty map of options
	params = make(map[string]string)

	err = appendOptions(o.GeneralOptions, &params)
	if err != nil {
		return params, err
	}
	err = appendOptions(o.LocaleOptions, &params)
	if err != nil {
		return params, err
	}
	err = appendOptions(o.LocationOptions, &params)
	if err != nil {
		return params, err
	}
	err = appendOptions(o.CoordinateOptions, &params)
	if err != nil {
		return params, err
	}
	err = appendOptions(o.BoundOptions, &params)
	if err != nil {
		return params, err
	}
	return params, nil
}

func appendOptions(optionProvider OptionProvider, output *map[string]string) (err error) {
	if optionProvider != nil {
		params, err := optionProvider.GetParameters()
		if err != nil {
			return err
		}
		appendToMap(&params, output)
	}
	return nil
}

func appendToMap(src, dst *map[string]string) {
	for k, v := range *src {
		(*dst)[k] = v
	}
}
