package models

type Coordinate struct {
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
}

type Location struct {
	Coordinate
	Country string `json:"country"`
	City    string `json:"city"`
	Address string `json:"address"`
}
