package yelp

import "net/url"

type Deal struct {
	id                      string       // Deal identifier
	title                   string       // Deal title
	url                     url.URL      // Deal url
	image_url               url.URL      // Deal image url
	currency_code           string       // ISO_4217 Currency Code
	time_start              float32      // Deal start time (Unix timestamp)
	time_end                float32      // Deal end time (optional: this field is present only if the Deal ends)
	is_popular              bool         // Whether the Deal is popular (optional: this field is present only if true)
	what_you_get            string       // Additional details for the Deal, separated by newlines
	important_restrictions  string       // Important restrictions for the Deal, separated by newlines
	additional_restrictions string       // Deal additional restrictions
	options                 []DealOption //Deal options

}

type DealOption struct {
	title                    string  // Deal option title
	purchase_url             url.URL // Deal option url for purchase
	price                    float32 // Deal option price (in cents)
	formatted_price          string  // Deal option price (formatted, e.g. "$6")
	original_price           float32 // Deal option original price (in cents)
	formatted_original_price string  // Deal option original price (formatted, e.g. "$12")
	is_quantity_limited      bool    // Whether the deal option is limited or unlimited
	remaining_count          float32 // The remaining deal options available for purchase (optional: this field is only present if the deal is limited)
}

type GiftCertificate struct {
	id              string                   // Gift certificate identifier
	url             url.URL                  // Gift certificate landing page url
	image_url       url.URL                  //	Gift certificate image url
	currency_code   string                   // ISO_4217 Currency Code
	unused_balances string                   // Whether unused balances are returned as cash or store credit
	options         []GiftCertificateOptions //	Gift certificate options
}

type GiftCertificateOptions struct {
	price           float32 //	Gift certificate option price (in cents)
	formatted_price string  //	Gift certificate option price (formatted, e.g. "$50")
}

type Review struct {
	id                     string  // Review identifier
	rating                 float32 // Rating from 1-5
	rating_image_url       url.URL // URL to star rating image for this business (size = 84x17)
	rating_image_small_url url.URL // URL to small version of rating image for this business (size = 50x10)
	rating_image_large_url url.URL // URL to large version of rating image for this business (size = 166x30)
	excerpt                string  // Review excerpt
	time_created           float32 // Time created (Unix timestamp)
	user                   User    // User who wrote the review
}

type User struct {
	id        string  // User identifier
	image_url url.URL // User profile image url
	name      string  // User name
}

type Category struct {
	name  string
	alias string
}

type Coordinate struct {
	latitude  float32 // Latitude of current location
	longitude float32 // Longitude of current location
}

type Location struct {
	coordinate      Coordinate // Address for this business formatted for display. Includes all address fields, cross streets and city, state_code, etc.
	address         []string   // Address for this business. Only includes address fields.
	display_address []string
	city            string   // City for this business
	state_code      string   // ISO 3166-2 state code for this business
	postal_code     string   // Postal code for this business
	country_code    string   // ISO 3166-1 country code for this business
	cross_streets   string   // Cross streets for this business
	neighborhoods   []string // List that provides neighborhood(s) information for business
}

type Business struct {
	id                   string            // Yelp ID for this business
	name                 string            // Name of this business
	image_url            string            // URL of photo for this business
	url                  string            // URL for business page on Yelp
	mobile_url           string            // URL for mobile business page on Yelp
	phone                string            // Phone number for this business with international dialing code (e.g. +442079460000)
	display_phone        string            // Phone number for this business formatted for display
	review_count         int               // Number of reviews for this business
	categories           []Category        // Provides a list of category name, alias pairs that this business is associated with. The alias is provided so you can search with the category_filter.
	distance             float32           // Distance that business is from search location in meters, if a latitude/longitude is specified.
	rating               float32           // Rating for this business (value ranges from 1, 1.5, ... 4.5, 5)
	rating_img_url       string            // URL to star rating image for this business (size = 84x17)
	rating_img_url_small string            // URL to small version of rating image for this business (size = 50x10)
	rating_img_url_large string            // URL to large version of rating image for this business (size = 166x30)
	snippet_text         string            // Snippet text associated with this business
	snippet_image_url    string            // URL of snippet image associated with this business
	location             Location          // Location data for this business
	is_claimed           bool              // Whether business has been claimed by a business owner
	is_closed            bool              // Whether business has been (permanently) closed
	menu_provider        string            // Provider of the menu for this business
	menu_date_updated    float32           // Last time this menu was updated on Yelp (Unix timestamp)
	deals                []Deal            // Deal info for this business (optional: this field is present only if there’s a Deal)
	gift_certificates    []GiftCertificate // Gift certificate info for this business (optional: this field is present only if there are gift certificates available)
	reviews              []Review          // Contains one review associated with business
}
