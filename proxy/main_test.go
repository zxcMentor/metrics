package main

import (
	"geoservicemetrics/proxy/internal/service"
	"testing"
)

func TestAddressGetting(t *testing.T) {
	addrReq := &service.AddressDadata{
		Query: `{"query":"нижний советс 3"}`}

	resp, err := service.AddressGetting(addrReq)

	if err != nil {
		t.Errorf("no errors expected but got: %v", err)
	}

	if resp == nil {
		t.Errorf("non-empty respond expected but got nil one")
	}
}

func TestGeocodeGetting(t *testing.T) {
	geoSuggReq := &service.GeocodeDadata{Lat: "55.878", Lng: "37.653"}

	resp, err := service.GeoSuggGetting(geoSuggReq)

	if err != nil {
		t.Errorf("no errors expected but got: %v", err)
	}

	if resp == nil {
		t.Errorf("non-empty respond expected but got nil one")
	}
}

func TestRegister(t *testing.T) {

	if resp := Register("user", "password"); resp != "user registerred" {
		t.Errorf("register error occured, expected \"user registerred\", but got %v", resp)
	}

}

func TestLogin(t *testing.T) {

	if resp := Login("user", "password"); resp != "user user authenticated" {
		t.Errorf("register error occured, expected \"user registerred\", but got %v", resp)
	}

}
