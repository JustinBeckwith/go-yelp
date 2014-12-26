package yelp

import "net/url"

type Deal struct {
	Id                      string       // Deal identifier
	Title                   string       // Deal title
	Url                     url.URL      // Deal url
	Image_url               url.URL      // Deal image url
	Currency_code           string       // ISO_4217 Currency Code
	Time_start              float32      // Deal start time (Unix timestamp)
	Time_end                float32      // Deal end time (optional: this field is present only if the Deal ends)
	Is_popular              bool         // Whether the Deal is popular (optional: this field is present only if true)
	What_you_get            string       // Additional details for the Deal, separated by newlines
	Important_restrictions  string       // Important restrictions for the Deal, separated by newlines
	Additional_restrictions string       // Deal additional restrictions
	Options                 []DealOption //Deal options

}

type DealOption struct {
	Title                    string  // Deal option title
	Purchase_url             url.URL // Deal option url for purchase
	Price                    float32 // Deal option price (in cents)
	Formatted_price          string  // Deal option price (formatted, e.g. "$6")
	Original_price           float32 // Deal option original price (in cents)
	Formatted_original_price string  // Deal option original price (formatted, e.g. "$12")
	Is_quantity_limited      bool    // Whether the deal option is limited or unlimited
	Remaining_count          float32 // The remaining deal options available for purchase (optional: this field is only present if the deal is limited)
}

type GiftCertificate struct {
	Id              string                   // Gift certificate identifier
	Url             url.URL                  // Gift certificate landing page url
	Image_url       url.URL                  //	Gift certificate image url
	Currency_code   string                   // ISO_4217 Currency Code
	Unused_balances string                   // Whether unused balances are returned as cash or store credit
	Options         []GiftCertificateOptions //	Gift certificate options
}

type GiftCertificateOptions struct {
	Price           float32 //	Gift certificate option price (in cents)
	Formatted_price string  //	Gift certificate option price (formatted, e.g. "$50")
}

type Review struct {
	Id                     string  // Review identifier
	Rating                 float32 // Rating from 1-5
	Rating_image_url       url.URL // URL to star rating image for this business (size = 84x17)
	Rating_image_small_url url.URL // URL to small version of rating image for this business (size = 50x10)
	Rating_image_large_url url.URL // URL to large version of rating image for this business (size = 166x30)
	Excerpt                string  // Review excerpt
	Time_created           float32 // Time created (Unix timestamp)
	User                   User    // User who wrote the review
}

type User struct {
	Id        string  // User identifier
	Image_url url.URL // User profile image url
	Name      string  // User name
}

type Category struct {
	Name  string
	Alias string
}

type Coordinate struct {
	Latitude  float32 // Latitude of current location
	Longitude float32 // Longitude of current location
}

type Location struct {
	Coordinate      Coordinate // Address for this business formatted for display. Includes all address fields, cross streets and city, state_code, etc.
	Address         []string   // Address for this business. Only includes address fields.
	Display_address []string
	City            string   // City for this business
	State_code      string   // ISO 3166-2 state code for this business
	Postal_code     string   // Postal code for this business
	Country_code    string   // ISO 3166-1 country code for this business
	Cross_streets   string   // Cross streets for this business
	Neighborhoods   []string // List that provides neighborhood(s) information for business
	Geo_accuracy    float32
}

type Business struct {
	Id                   string            // Yelp ID for this business
	Name                 string            // Name of this business
	Image_url            string            // URL of photo for this business
	Url                  string            // URL for business page on Yelp
	Mobile_url           string            // URL for mobile business page on Yelp
	Phone                string            // Phone number for this business with international dialing code (e.g. +442079460000)
	Display_phone        string            // Phone number for this business formatted for display
	Review_count         int               // Number of reviews for this business
	Categories           [][]string        // Provides a list of category name, alias pairs that this business is associated with. The alias is provided so you can search with the category_filter.
	Distance             float32           // Distance that business is from search location in meters, if a latitude/longitude is specified.
	Rating               float32           // Rating for this business (value ranges from 1, 1.5, ... 4.5, 5)
	Rating_img_url       string            // URL to star rating image for this business (size = 84x17)
	Rating_img_url_small string            // URL to small version of rating image for this business (size = 50x10)
	Rating_img_url_large string            // URL to large version of rating image for this business (size = 166x30)
	Snippet_text         string            // Snippet text associated with this business
	Snippet_image_url    string            // URL of snippet image associated with this business
	Location             Location          // Location data for this business
	Is_claimed           bool              // Whether business has been claimed by a business owner
	Is_closed            bool              // Whether business has been (permanently) closed
	Menu_provider        string            // Provider of the menu for this business
	Menu_date_updated    float32           // Last time this menu was updated on Yelp (Unix timestamp)
	Deals                []Deal            // Deal info for this business (optional: this field is present only if there’s a Deal)
	Gift_certificates    []GiftCertificate // Gift certificate info for this business (optional: this field is present only if there are gift certificates available)
	Reviews              []Review          // Contains one review associated with business
}
