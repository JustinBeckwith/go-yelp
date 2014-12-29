package yelp

/**
 * Options for locale
 */
type LocaleOptions struct {
	cc   string // ISO 3166-1 alpha-2 country code. Default country to use when parsing the location field. United States = US, Canada = CA, United Kingdom = GB (not UK).
	lang string // ISO 639 language code (default=en). Reviews written in the specified language will be shown.
}

func (o *LocaleOptions) GetParameters() (params map[string]string, err error) {
	params = make(map[string]string)
	if o.cc != "" {
		params["cc"] = o.cc
	}
	if o.lang != "" {
		params["lang"] = o.lang
	}
	return params, nil
}
