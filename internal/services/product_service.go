package services

import (
	"github.com/dnogueir-org/international-pricebook-challenge/internal/models"
	"github.com/dnogueir-org/international-pricebook-challenge/repository"
)

type ProductService struct {
	ProductRepository  repository.ProductRepository
	PriceRepository    repository.PriceRepository
	CurrencyRepository repository.CurrencyRepository
}

func NewProductService(productRepository repository.ProductRepository, priceRepository repository.PriceRepository, currencyRepository repository.CurrencyRepository) *ProductService {
	return &ProductService{
		ProductRepository:  productRepository,
		PriceRepository:    priceRepository,
		CurrencyRepository: currencyRepository,
	}
}

func (ps *ProductService) Insert(name string, price float64, currencyId string) (*models.Product, error) {
	product, err := models.NewProduct(name)
	if err != nil {
		return nil, err
	}

	product, err = ps.ProductRepository.Insert(product)
	if err != nil {
		return nil, err
	}

	currency, err := ps.CurrencyRepository.Find(currencyId)
	if err != nil {
		return nil, err
	}

	productPrice, err := models.NewPrice(product.ID, *currency, price)
	if err != nil {
		return nil, err
	}

	_, err = ps.PriceRepository.Insert(productPrice)
	if err != nil {
		return nil, err
	}

	product.Prices = append(product.Prices, *productPrice)

	return product, nil
}

func (ps *ProductService) Find(id string) (*models.Product, error) {

	product, err := ps.ProductRepository.Find(id)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (ps *ProductService) FindAll() ([]*models.Product, error) {

	products, err := ps.ProductRepository.FindAll()
	if err != nil {
		return nil, err
	}

	return products, nil
}
