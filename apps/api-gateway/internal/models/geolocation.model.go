package models

type Coordinate struct {
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
}

type Geolocation struct {
	Coordinate
	Country string `json:"country"`
	City    string `json:"city"`
	Address string `json:"address"`
}

type GeolocationIPResponse struct {
	Result Geolocation `json:"result"`
}

type GeolocationAddressResponse struct {
	Results []Geolocation `json:"results"`
}
