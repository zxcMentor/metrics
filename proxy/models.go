package main

type Location struct {
	Location LocationClass `json:"location"`
}
type LocationClass struct {
	CityName   string             `json:"value"`
	CityPostal string             `json:"unrestricted_value"`
	Data       map[string]*string `json:"data"`
}

type RequestAddressGeocode struct {
	Lat string `json:"lat"`
	Lng string `json:"lng"`
}

// swagger:parameters RequestAddressSearch
type RequestAddressSearch struct {
	// in:query
	Query string `json:"query"`
}

// swagger:parameters ResponseAddress
type ResponseAddress struct {
	// in:addresses
	Addresses []*Address `json:"addresses"`
}

type Address string
