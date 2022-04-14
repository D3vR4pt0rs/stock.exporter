package entities

import "time"

type CandlestickData struct {
	Ticker string    `json:"ticker"`
	Date   time.Time `json:"date"`
	Open   float64   `json:"open"`
	High   float64   `json:"high"`
	Low    float64   `json:"low"`
	Close  float64   `json:"close"`
	Volume int       `json:"volume"`
}

type Company struct {
	Name   string `json:"name"`
	Ticker string `json:"ticker"`
}

type Stock struct {
	Symbol    string    `json:"symbol"`
	LastPrice float64   `json:"last_price"`
	Timestamp time.Time `json:"timestamp"`
}
