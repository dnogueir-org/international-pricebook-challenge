package services

import (
	"github.com/dnogueir-org/international-pricebook-challenge/internal/models"
	"github.com/dnogueir-org/international-pricebook-challenge/repository"
)

type PriceService struct {
	PriceRepository repository.PriceRepository
}

func (ps *PriceService) Insert(price *models.Price) (*models.Price, error) {

	err := price.Validate()
	if err != nil {
		return nil, err
	}

	price, err = ps.PriceRepository.Insert(price)
	if err != nil {
		return nil, err
	}

	return price, nil
}

func (ps *PriceService) Find(id string) (*models.Price, error) {

	price, err := ps.PriceRepository.Find(id)
	if err != nil {
		return nil, err
	}

	return price, nil
}
