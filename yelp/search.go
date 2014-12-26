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

type SearchParameters struct {
	Term            string     // Search term (e.g. "food", "restaurants"). If term isn’t included we search everything.
	Limit           int        // Number of business results to return
	Offset          int        // Offset the list of returned business results by this amount
	Sort            int        // Sort mode: 0=Best matched (default), 1=Distance, 2=Highest Rated. If the mode is 1 or 2 a search may retrieve an additional 20 businesses past the initial limit of the first 20 results. This is done by specifying an offset and limit of 20. Sort by distance is only supported for a location or geographic search. The rating sort is not strictly sorted by the rating value, but by an adjusted rating value that takes into account the number of ratings, similar to a bayesian average. This is so a business with 1 rating of 5 stars doesn’t immediately jump to the top.
	Category_filter string     // Category to filter search results with. See the list of supported categories. The category filter can be a list of comma delimited categories. For example, 'bars,french' will filter by Bars and French. The category identifier should be used (for example 'discgolf', not 'Disc Golf').
	Radius_filter   float32    // Search radius in meters. If the value is too large, a AREA_TOO_LARGE error may be returned. The max value is 40000 meters (25 miles).
	Deals_filter    bool       // Whether to exclusively search for businesses with deals
	Location        string     // [required]	Specifies the combination of "address, neighborhood, city, state or zip, optional country" to be used when searching for businesses.
	Coordinates     Coordinate // An optional latitude, longitude parameter can also be specified as a hint to the geocoder to disambiguate the location text.
	Bounds          Bounds     // Location is specified by a bounding box, defined by a southwest latitude/longitude and a northeast latitude/longitude geographic coordinate.
	Cc              string     // ISO 3166-1 alpha-2 country code. Default country to use when parsing the location field. United States = US, Canada = CA, United Kingdom = GB (not UK).
	Lang            string     // ISO 639 language code (default=en). Reviews and snippets written in the specified language will be shown.
}

// Location is specified by a bounding box, defined by a southwest latitude/longitude and a northeast latitude/longitude geographic coordinate.
// The bounding box format is defined as:
// bounds=sw_latitude,sw_longitude|ne_latitude,ne_longitude
type Bounds struct {
	Sw_latitude  float32 //	Southwest latitude of bounding box
	Sw_longitude float32 //	Southwest longitude of bounding box
	Ne_latitude  float32 //	Northeast latitude of bounding box
	Ne_longitude float32 //	Northeast longitude of bounding box
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
