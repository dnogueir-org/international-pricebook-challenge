package repository_test

import (
	"testing"

	"github.com/dnogueir-org/international-pricebook-challenge/database"
	"github.com/dnogueir-org/international-pricebook-challenge/internal/models"
	"github.com/dnogueir-org/international-pricebook-challenge/repository"
	"github.com/stretchr/testify/require"
)

func TestCountryRepositoryDb(t *testing.T) {
	// given - a memory test database
	db := database.NewDbTest()
	defer db.Close()

	// and - an instance of a currency
	name, acronym := "Real", "BRL"
	currency, _ := models.NewCurrency(name, acronym)

	// and - inserting the currency in the database
	currencyRepo := repository.NewCurrencyRepository(db)
	currencyRepo.Insert(currency)

	// and - an instance of a country
	countryName, abbreviation := "Brasil", "BRA"
	country, _ := models.NewCountry(countryName, abbreviation, *currency)

	// when - inserting the country in the database
	countryRepo := repository.NewCountryRepository(db)
	_, err := countryRepo.Insert(country)
	// then - no error should be returned
	require.Nil(t, err)

	// when - finding the country in the database
	c, err := countryRepo.Find(country.ID)

	// then - no error should be returned and IDs should me matched
	require.Nil(t, err)
	require.Equal(t, country.ID, c.ID)
}
