package client

import "exporter/internal/entities"

type stockApiInteractor interface {
	GetStockInformationByTicker(ticker string) ([]entities.Stock, error)
}

type stockApiClient struct {
	saInteractor stockApiInteractor
}

func New(saInteractor stockApiInteractor) *stockApiClient {
	return &stockApiClient{
		saInteractor: saInteractor,
	}
}

func (saClient stockApiClient) GetStockInformationByTicker(ticker string) ([]entities.Stock, error) {
	return saClient.saInteractor.GetStockInformationByTicker(ticker)
}
