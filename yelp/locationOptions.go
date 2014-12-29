package yelp

import (
	"errors"
	"fmt"
)

type LocationOptionsBase interface {
}

/**
 * Specify Location by Neighborhood, Address, or City
 * Location is specified by a particular neighborhood, address or city.
 * The location format is defined: ?location=location
 */
type LocationOptions struct {
	Location          string             // Specifies the combination of "address, neighborhood, city, state or zip, optional country" to be used when searching for businesses. (required)
	CoordinateOptions *CoordinateOptions // An optional latitude, longitude parameter can also be specified as a hint to the geocoder to disambiguate the location text. The format for this is defined as:   ?cll=latitude,longitude
}

func (o *LocationOptions) GetParameters() (params map[string]string, err error) {
	params = make(map[string]string)

	// location is a required field
	if o.Location == "" {
		return params, errors.New("To perform a location based search, the location property must contain an area within to search.  For coordinate based searches, use the CoordinateOption class.")
	}
	params["location"] = o.Location

	// if coordinates are specified add those to the parameters hash
	if o.CoordinateOptions != nil &&
		o.CoordinateOptions.Latitude.Valid &&
		o.CoordinateOptions.Longitude.Valid {
		params["cll"] = fmt.Sprintf("%v,%v", o.CoordinateOptions.Latitude.Float64, o.CoordinateOptions.Longitude.Float64)
	}

	return params, nil
}
