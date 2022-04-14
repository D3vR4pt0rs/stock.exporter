package collector

import "exporter/internal/entities"

type client interface {
	GetInformationAboutStock(ticker string) (entities.Stock, error)
	GetAvailableCompanies() []entities.Company
}

type Controller interface {
	GetInformationAboutStock(ticker string) (entities.Stock, error)
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

func (app *application) GetInformationAboutStock(ticker string) (entities.Stock, error) {
	return app.client.GetInformationAboutStock(ticker)
}

func (app *application) GetAvailableCompanies() []entities.Company {
	return app.client.GetAvailableCompanies()
}
