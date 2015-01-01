package yelp

import (
	"fmt"
)

// BoundOptions provides a way to describe a location that uses a set of coordinates.
// Location is specified by a bounding box, defined by a southwest latitude/longitude and a
// northeast latitude/longitude geographic coordinate.
// The bounding box format is defined as:
// bounds=sw_latitude,sw_longitude|ne_latitude,ne_longitude
type BoundOptions struct {
	SwLatitude  float32 // Southwest latitude of bounding box
	SwLongitude float32 // Southwest longitude of bounding box
	NeLatitude  float32 // Northeast latitude of bounding box
	NeLongitude float32 // Northeast longitude of bounding box
}

// getParameters will reflect over the values of the given struct, and provide a type
// appropriate set of querystring parameters that match the defined values.
func (o BoundOptions) getParameters() (params map[string]string, err error) {
	return map[string]string{
		"bounds": fmt.Sprintf("%v,%v|%v,%v",
			o.SwLatitude,
			o.SwLongitude,
			o.NeLatitude,
			o.NeLongitude,
		),
	}, nil
}
