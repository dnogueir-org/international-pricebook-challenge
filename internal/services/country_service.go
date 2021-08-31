package services

import (
	"github.com/dnogueir-org/international-pricebook-challenge/internal/models"
	"github.com/dnogueir-org/international-pricebook-challenge/repository"
)

type CountryService struct {
	CountryRepository  repository.CountryRepository
	CurrencyRepository repository.CurrencyRepository
}

func NewCountryService(countryRepository repository.CountryRepository, currencyRepository repository.CurrencyRepository) *CountryService {
	return &CountryService{
		CountryRepository:  countryRepository,
		CurrencyRepository: currencyRepository,
	}
}

func (cs *CountryService) Insert(name string, abbreviation string, currencyId string) (*models.Country, error) {

	currency, err := cs.CurrencyRepository.Find(currencyId)
	if err != nil {
		return nil, err
	}

	country, err := models.NewCountry(name, abbreviation, *currency)
	if err != nil {
		return nil, err
	}

	country, err = cs.CountryRepository.Insert(country)
	if err != nil {
		return nil, err
	}

	return country, nil
}

func (cs *CountryService) Find(id string) (*models.Country, error) {

	country, err := cs.CountryRepository.Find(id)
	if err != nil {
		return nil, err
	}

	return country, nil
}

func (cs *CountryService) FindAll() ([]*models.Country, error) {
	countries, err := cs.CountryRepository.FindAll()
	if err != nil {
		return nil, err
	}

	return countries, nil
}
