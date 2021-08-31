package repository

import (
	"fmt"

	"github.com/dnogueir-org/international-pricebook-challenge/internal/models"
	"github.com/jinzhu/gorm"
)

type CurrencyRepository interface {
	Insert(currency *models.Currency) (*models.Currency, error)
	Find(id string) (*models.Currency, error)
	FindAll() ([]*models.Currency, error)
}

type CurrencyRepositoryDb struct {
	Db *gorm.DB
}

func NewCurrencyRepository(db *gorm.DB) *CurrencyRepositoryDb {
	return &CurrencyRepositoryDb{Db: db}
}

func (repo CurrencyRepositoryDb) Insert(currency *models.Currency) (*models.Currency, error) {

	err := repo.Db.Create(currency).Error

	if err != nil {
		return nil, err
	}

	return currency, nil
}

func (repo CurrencyRepositoryDb) Find(id string) (*models.Currency, error) {

	var currency models.Currency
	repo.Db.Preload("Countries").First(&currency, "id = ?", id)

	if currency.ID == "" {
		return nil, fmt.Errorf("currency does not exist")
	}

	return &currency, nil

}

func (repo CurrencyRepositoryDb) FindAll() ([]*models.Currency, error) {

	var currencies []*models.Currency
	repo.Db.Find(&currencies)

	return currencies, nil

}
