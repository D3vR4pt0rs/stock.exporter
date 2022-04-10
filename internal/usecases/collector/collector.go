package collector

import "exporter/internal/entities"

type client interface {
	GetStockInformationByTicker(ticker string) ([]entities.Stock, error)
}

type application struct {
	client client
}

func New(client client) *application {
	return &application{
		client: client,
	}
}

func (app *application) GetStockInformationById(ticker string) ([]entities.Stock, error) {
	return app.client.GetStockInformationByTicker(ticker)
}
