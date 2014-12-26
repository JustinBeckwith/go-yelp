package yelp

type SearchResult struct {
	Region     Region     // Suggested bounds in a map to display results in
	Total      int        // Total number of business results
	Businesses []Business // The list of business entries (see Business)
}

type Region struct {
	Span   Span   // Span of suggested map bounds
	Center Center // Center position of map bounds
}

type Span struct {
	Latitude_delta  float32 // Latitude width of map bounds
	Longitude_delta float32 // Longitude height of map bounds
}

type Center struct {
	Latitude  float32 // Latitude position of map bounds center
	Longitude float32 // Longitude position of map bounds center
}

// Specify Location by Geographic Coordinate
// The geographic coordinate format is defined as:
// ll=latitude,longitude,accuracy,altitude,altitude_accuracy
type GeographicCoordinate struct {
	Latitude          float32 // Latitude of geo-point to search near
	Longitude         float32 // Longitude of geo-point to search near
	Accuracy          float32 // Accuracy of latitude, longitude
	Altitude          float32 // Altitude
	Altitude_accuracy float32 // Accuracy of altitude
}
