package services

import (
	"github.com/dnogueir-org/international-pricebook-challenge/internal/models"
	"github.com/dnogueir-org/international-pricebook-challenge/repository"
)

type CurrencyService struct {
	CurrencyRepository repository.CurrencyRepository
}

func (cs *CurrencyService) Insert(currency *models.Currency) (*models.Currency, error) {

	err := currency.Validate()
	if err != nil {
		return nil, err
	}

	currency, err = cs.CurrencyRepository.Insert(currency)
	if err != nil {
		return nil, err
	}

	return currency, nil
}

func (cs *CurrencyService) Find(id string) (*models.Currency, error) {

	currency, err := cs.CurrencyRepository.Find(id)
	if err != nil {
		return nil, err
	}

	return currency, nil
}
