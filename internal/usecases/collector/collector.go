package collector

import "exporter/internal/entities"

type client interface {
	GetStockInformationByTicker(ticker string) ([]entities.Stock, error)
}

type Controller interface {
	GetStockInformationByTicker(ticker string) ([]entities.Stock, error)
	GetAvailableCompanies() []entities.Company
}

type application struct {
	client client
}

func New(client client) *application {
	return &application{
		client: client,
	}
}

var availableTickers = map[string]string{
	"Apple":     "AAPL",
	"Microsoft": "MSFT",
	"IBM":       "IBM",
	"Tesla":     "TSLA",
	"Google":    "GOOGL",
}

func (app *application) GetStockInformationByTicker(ticker string) ([]entities.Stock, error) {
	return app.client.GetStockInformationByTicker(ticker)
}

func (app *application) GetAvailableCompanies() []entities.Company {
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
