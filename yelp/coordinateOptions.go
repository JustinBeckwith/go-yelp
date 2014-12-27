package yelp

import (
	"errors"
	"fmt"
	"github.com/guregu/null"
)

/**
 * The geographic coordinate format is defined as:
 * ll=latitude,longitude,accuracy,altitude,altitude_accuracy
 */
type CoordinateOptions struct {
	Latitude          null.Float // Latitude of geo-point to search near (required)
	Longitude         null.Float // Longitude of geo-point to search near (required)
	Accuracy          null.Float // Accuracy of latitude, longitude (optional)
	Altitude          null.Float // Altitude (optional)
	Altitude_accuracy null.Float // Accuracy of altitude (optional)
}

func (o CoordinateOptions) GetParameters() (params map[string]string, err error) {
	// coordinate requires at least a latitude and longitude - others are option
	if !o.Latitude.Valid || !o.Longitude.Valid {
		return nil, errors.New("latitude and longitude are required fields for a coordinate based search")
	}

	ll := fmt.Sprintf("%v,%v", o.Latitude.Float64, o.Longitude.Float64)
	if o.Accuracy.Valid {
		ll += fmt.Sprintf(",%v", o.Accuracy.Float64)
	}
	if o.Altitude.Valid {
		ll += fmt.Sprintf(",%v", o.Altitude.Float64)
	}
	if o.Altitude_accuracy.Valid {
		ll += fmt.Sprintf(",%v", o.Altitude_accuracy.Float64)
	}

	return map[string]string{
		"ll": ll,
	}, nil
}
