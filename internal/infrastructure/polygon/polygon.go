package polygon

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"exporter/internal/entities"
	"github.com/D3vR4pt0rs/logger"
)

type client struct {
	token      string `json:"token"`
	httpClient *http.Client
}

func New(token string) *client {
	return &client{
		token:      token,
		httpClient: &http.Client{},
	}
}

const POLYGON_URL = "https://api.polygon.io"

type AggregateResponse struct {
	Adjusted     bool     `json:"adjusted"`
	QueryCount   int      `json:"queryCount"`
	RequestId    string   `json:"request_id"`
	Results      []Ticker `json:"results"`
	ResultsCount int      `json:"resultsCount"`
	Status       string   `json:"status"`
	Ticker       string   `json:"ticker"`
}

type Ticker struct {
	Close                float64 `json:"c"`
	High                 float64 `json:"h"`
	Low                  float64 `json:"l"`
	NumberTransactions   int     `json:"n"`
	Open                 float64 `json:"o"`
	Timestamp            int64   `json:"t"`
	Volume               int     `json:"v"`
	AveragePriceOfVolume float64 `json:"vw"`
}

func (c client) GetInformationForCandle(ticker string) ([]entities.CandlestickData, error) {
	url := fmt.Sprintf("%s/v2/aggs/ticker/%s/range/5/minute/2022-04-08/2022-04-08?adjusted=true&sort=asc", POLYGON_URL, ticker)

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		logger.Error.Println(err.Error())
		return []entities.CandlestickData{}, err
	}
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.token))

	resp, err := c.httpClient.Do(request)
	if err != nil {
		logger.Error.Println(err.Error())
		return []entities.CandlestickData{}, err
	}

	defer resp.Body.Close()

	var result AggregateResponse
	json.NewDecoder(resp.Body).Decode(&result)

	var stockInformations []entities.CandlestickData
	for _, rawStock := range result.Results {
		tUnixNanoRemainder := (rawStock.Timestamp % int64(time.Microsecond)) * int64(time.Millisecond)
		stock := entities.CandlestickData{
			Ticker: result.Ticker,
			Open:   rawStock.Open,
			Close:  rawStock.Close,
			High:   rawStock.High,
			Low:    rawStock.Low,
			Date:   time.Unix(rawStock.Timestamp/int64(time.Microsecond), tUnixNanoRemainder),
			Volume: rawStock.Volume,
		}
		stockInformations = append(stockInformations, stock)
	}
	return stockInformations, nil
}
