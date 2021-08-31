package services

import (
	"github.com/dnogueir-org/international-pricebook-challenge/internal/models"
	"github.com/dnogueir-org/international-pricebook-challenge/repository"
)

type CurrencyService struct {
	CurrencyRepository repository.CurrencyRepository
}

func NewCurrencyService(currencyRepository repository.CurrencyRepository) *CurrencyService {
	return &CurrencyService{CurrencyRepository: currencyRepository}
}

func (cs *CurrencyService) Insert(name string, acronym string) (*models.Currency, error) {

	currency, err := models.NewCurrency(name, acronym)
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

func (cs *CurrencyService) FindAll() ([]*models.Currency, error) {

	currencies, err := cs.CurrencyRepository.FindAll()
	if err != nil {
		return nil, err
	}

	return currencies, nil
}
