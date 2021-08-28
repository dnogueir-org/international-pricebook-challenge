package repository

import (
	"fmt"

	"github.com/dnogueir-org/international-pricebook-challenge/internal/models"
	"github.com/jinzhu/gorm"
)

type CountryRepository interface {
	Insert(country *models.Country) (*models.Country, error)
	Find(id string) (*models.Country, error)
}

type CountryRepositoryDb struct {
	Db *gorm.DB
}

func NewCountryRepository(db *gorm.DB) *CountryRepositoryDb {
	return &CountryRepositoryDb{Db: db}
}

func (repo CountryRepositoryDb) Insert(country *models.Country) (*models.Country, error) {

	err := repo.Db.Create(country).Error

	if err != nil {
		return nil, err
	}

	return country, nil
}

func (repo CountryRepositoryDb) Find(id string) (*models.Country, error) {

	var country models.Country
	repo.Db.Preload("Currency").First(&country, "id = ?", id)

	if country.ID == "" {
		return nil, fmt.Errorf("country does not exist")
	}

	return &country, nil

}
