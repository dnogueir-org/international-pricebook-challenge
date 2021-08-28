package services

import (
	"github.com/dnogueir-org/international-pricebook-challenge/internal/models"
	"github.com/dnogueir-org/international-pricebook-challenge/repository"
)

type CountryService struct {
	CountryRepository repository.CountryRepository
}

func (cs *CountryService) Insert(country *models.Country) (*models.Country, error) {

	err := country.Validate()
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
