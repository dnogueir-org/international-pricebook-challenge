package repository

import (
	"fmt"

	"github.com/dnogueir-org/international-pricebook-challenge/internal/models"
	"github.com/jinzhu/gorm"
)

type ProductRepository interface {
	Insert(product *models.Product) (*models.Product, error)
	Find(id string) (*models.Product, error)
}

type ProductRepositoryDb struct {
	Db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepositoryDb {
	return &ProductRepositoryDb{Db: db}
}

func (repo ProductRepositoryDb) Insert(product *models.Product) (*models.Product, error) {

	err := repo.Db.Create(product).Error

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (repo ProductRepositoryDb) Find(id string) (*models.Product, error) {

	var product models.Product
	repo.Db.Preload("Prices").First(&product, "id = ?", id)

	if product.ID == "" {
		return nil, fmt.Errorf("product does not exist")
	}

	return &product, nil

}
