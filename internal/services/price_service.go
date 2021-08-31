package services

import (
	"github.com/dnogueir-org/international-pricebook-challenge/internal/models"
	"github.com/dnogueir-org/international-pricebook-challenge/repository"
)

type PriceService struct {
	PriceRepository    repository.PriceRepository
	CurrencyRepository repository.CurrencyRepository
}

func NewPriceService(priceRepository repository.PriceRepository, currencyRepository repository.CurrencyRepository) *PriceService {
	return &PriceService{
		PriceRepository:    priceRepository,
		CurrencyRepository: currencyRepository,
	}
}

func (ps *PriceService) Insert(productId string, currencyId string, price float64) (*models.Price, error) {

	currency, err := ps.CurrencyRepository.Find(currencyId)
	if err != nil {
		return nil, err
	}

	productPrice, err := models.NewPrice(productId, *currency, price)
	if err != nil {
		return nil, err
	}

	productPrice, err = ps.PriceRepository.Insert(productPrice)
	if err != nil {
		return nil, err
	}

	return productPrice, nil
}

func (ps *PriceService) Find(id string) (*models.Price, error) {

	price, err := ps.PriceRepository.Find(id)
	if err != nil {
		return nil, err
	}

	return price, nil
}
