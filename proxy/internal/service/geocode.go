package service

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

type Geocoder interface {
	GetGeocode() ([]byte, error)
}

type GeocodeDadata struct {
	Lat string `json:"lat"`
	Lng string `json:"lng"`
}

func GeocodeRequestUnmarshal(req []byte) (*GeocodeRequest, error) {
	reqUnmrsh := GeocodeRequest{}

	err := json.Unmarshal(req, &reqUnmrsh)

	return &reqUnmrsh, err
}

func UnmarshalGeocode(data []byte) (Geocode, error) {
	var r Geocode
	err := json.Unmarshal(data, &r)
	return r, err
}

func (inp *GeocodeDadata) GetGeocode() ([]byte, error) {

	client := &http.Client{}
	lat := inp.Lat
	long := inp.Lng
	var data = strings.NewReader(fmt.Sprintf("{ \"lat\":%s, \"lon\":%s }", lat, long)) //(`{ "lat": 55.878, "lon": 37.653 }`)
	req, err := http.NewRequest("POST", "https://suggestions.dadata.ru/suggestions/api/4_1/rs/geolocate/address", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Token dadff476c93c94605f6bb06c4c3ba8e9d6ecaa8f")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return bodyText, err
}

func UnmarshalSuggestionsByGeo(data []byte) (*SuggestionsByGeo, error) {
	var suggs SuggestionsByGeo
	err := json.Unmarshal(data, &suggs)
	return &suggs, err
}

func getResponseAddressByGeo(suggs *SuggestionsByGeo) ([]byte, error) {

	addrResp := &GeocodeResponse{}

	for _, sugg := range suggs.Suggestions {

		addr := Address(sugg.Address)
		addrResp.Addresses = append(addrResp.Addresses, &addr)

	}

	return json.Marshal(addrResp)

}

func GeoSuggGetting(inp *GeocodeDadata) (*SuggestionsByGeo, error) {

	//georeq, err := GeocodeRequestUnmarshal(inp)
	//if err != nil {
	//	fmt.Println("request unmarshaling error occured", err)
	//}

	resp, err := inp.GetGeocode()
	if err != nil {
		fmt.Println("respond error occured", err)
	}

	respUnmr, err := UnmarshalSuggestionsByGeo(resp)
	if err != nil {
		fmt.Println("respond unmarshalling error occured", err)
	}

	//outp, err := getResponseAddressByGeo(respUnmr)
	//if err != nil {
	//	fmt.Println("respond marshalling error occured", err)
	//}

	return respUnmr, err

}

type GeocodeRequest struct {
	Lat string `json:"lat"`
	Lng string `json:"lng"`
}

type SuggestionsByGeo struct {
	Suggestions []Suggestion `json:"suggestions"`
}

type Suggestion struct {
	Address     string             `json:"value"`
	FullAddress string             `json:"unrestricted_value"`
	Data        map[string]*string `json:"data"`
}

type Geocode []GeocodeElement

type GeocodeElement struct {
	Source       string `json:"source"`
	Result       string `json:"result"`
	PostalCode   string `json:"postal_code"`
	Country      string `json:"country"`
	Region       string `json:"region"`
	CityArea     string `json:"city_area"`
	CityDistrict string `json:"city_district"`
	Street       string `json:"street"`
	House        string `json:"house"`
	GeoLat       string `json:"geo_lat"`
	GeoLon       string `json:"geo_lon"`
	QcGeo        int64  `json:"qc_geo"`
}

type GeocodeResponse struct {
	Addresses []*Address `json:"addresses"`
}

type Address string
