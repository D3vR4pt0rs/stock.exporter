package finnhub

type Stock struct {
	Symbol     string   `json:"s"`
	LastPrice  float64  `json:"p"`
	Timestamp  int64    `json:"t"`
	Volume     float64  `json:"v"`
	Conditions []string `json:"c"`
}

type QuoteStock struct {
	CurrentPrice       float64 `json:"c"`
	Change             float64 `json:"d,omitempty"`
	PercentChange      float64 `json:"dp,omitempty"`
	HighPrice          float64 `json:"h"`
	LowPrice           float64 `json:"l"`
	OpenPrice          float64 `json:"o"`
	Timestamp          int64   `json:"t"`
	PreviousClosePrice float64 `json:"pc"`
}

type Message struct {
	Data []Stock `json:"data"`
	Type string  `json:"type"`
}

type CurrentStockValue struct {
	LastPrice float64 `json:"p"`
	Volume    float64 `json:"v"`
}
