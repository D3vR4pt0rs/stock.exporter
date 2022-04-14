package client

import (
	"exporter/internal/entities"
)

var availableTickers = map[string]string{
	"Apple":     "AAPL",
	"Microsoft": "MSFT",
	"IBM":       "IBM",
	"Tesla":     "TSLA",
	"Google":    "GOOGL",
}

type StockClient interface {
	GetQuoteStock(ticker string) (entities.Stock, error)
	OpenWebSocketConnection(symbols []string)
}

type stockApiClient struct {
	stockClient StockClient
}

func New(stockClient StockClient) *stockApiClient {
	// go stockClient.OpenWebSocketConnection(getTickers())
	return &stockApiClient{
		stockClient: stockClient,
	}
}

func getTickers() []string {
	var tickers []string
	for _, value := range availableTickers {
		tickers = append(tickers, value)
	}
	return tickers
}

func (saClient stockApiClient) GetAvailableCompanies() []entities.Company {
	var companies []entities.Company
	for key, value := range availableTickers {
		company := entities.Company{
			Name:   key,
			Ticker: value,
		}
		companies = append(companies, company)
	}
	return companies
}

func (saClient stockApiClient) GetInformationAboutStock(ticker string) (entities.Stock, error) {
	return saClient.stockClient.GetQuoteStock(ticker)
}
