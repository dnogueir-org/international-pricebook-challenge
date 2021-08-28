package services

import (
	"github.com/dnogueir-org/international-pricebook-challenge/internal/models"
	"github.com/dnogueir-org/international-pricebook-challenge/repository"
)

type ProductService struct {
	ProductRepository repository.ProductRepository
}

func NewProductService(repository repository.ProductRepository) *ProductService {
	return &ProductService{ProductRepository: repository}
}

func (ps *ProductService) Insert(name string) (*models.Product, error) {
	product, err := models.NewProduct(name)
	if err != nil {
		return nil, err
	}

	product, err = ps.ProductRepository.Insert(product)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (ps *ProductService) Find(id string) (*models.Product, error) {

	product, err := ps.ProductRepository.Find(id)
	if err != nil {
		return nil, err
	}

	return product, nil
}
