package entities

import "time"

type Stock struct {
	Ticker string    `json:"ticker"`
	Date   time.Time `json:"date"`
	Open   float64   `json:"open"`
	High   float64   `json:"high"`
	Low    float64   `json:"low"`
	Close  float64   `json:"close"`
	Volume int       `json:"volume"`
}
