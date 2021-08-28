package repository

import (
	"fmt"

	"github.com/dnogueir-org/international-pricebook-challenge/internal/models"
	"github.com/jinzhu/gorm"
)

type PriceRepository interface {
	Insert(price *models.Price) (*models.Price, error)
	Find(id string) (*models.Price, error)
}

type PriceRepositoryDb struct {
	Db *gorm.DB
}

func NewPriceRepository(db *gorm.DB) *PriceRepositoryDb {
	return &PriceRepositoryDb{Db: db}
}

func (repo PriceRepositoryDb) Insert(price *models.Price) (*models.Price, error) {

	err := repo.Db.Create(price).Error

	if err != nil {
		return nil, err
	}

	return price, nil
}

func (repo PriceRepositoryDb) Find(id string) (*models.Price, error) {

	var price models.Price
	repo.Db.Preload("Currency").First(&price, "id = ?", id)

	if price.ID == "" {
		return nil, fmt.Errorf("price does not exist")
	}

	return &price, nil

}
