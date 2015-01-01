package yelp

import (
	"errors"
	"fmt"
)

// LocationOptions enable specifing a Location by Neighborhood, Address, or City.
// The location format is defined: ?location=location
type LocationOptions struct {
	Location          string             // Specifies the combination of "address, neighborhood, city, state or zip, optional country" to be used when searching for businesses. (required)
	CoordinateOptions *CoordinateOptions // An optional latitude, longitude parameter can also be specified as a hint to the geocoder to disambiguate the location text. The format for this is defined as:   ?cll=latitude,longitude
}

// getParameters will reflect over the values of the given
// struct, and provide a type appropriate set of querystring parameters
// that match the defined values.
func (o *LocationOptions) getParameters() (params map[string]string, err error) {
	params = make(map[string]string)

	// location is a required field
	if o.Location == "" {
		return params, errors.New("to perform a location based search, the location property must contain an area within to search.  For coordinate based searches, use the CoordinateOption class")
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
