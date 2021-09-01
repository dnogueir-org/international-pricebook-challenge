package repository_test

import (
	"testing"

	"github.com/dnogueir-org/international-pricebook-challenge/database"
	"github.com/dnogueir-org/international-pricebook-challenge/internal/models"
	"github.com/dnogueir-org/international-pricebook-challenge/repository"
	"github.com/stretchr/testify/require"
)

func TestCurrencyRepositoryDbInsert(t *testing.T) {
	// given - a memory test database
	db := database.NewDbTest()
	defer db.Close()

	// and - an instance of a currency
	name, acronym := "Real", "BRL"
	currency, _ := models.NewCurrency(name, acronym)

	// when - inserting the currency in the database
	repo := repository.NewCurrencyRepository(db)
	_, err := repo.Insert(currency)

	// then - no error should be returned
	require.Nil(t, err)

	// when - finding the currency in the database
	c, err := repo.Find(currency.ID)

	// then - no error should be returned and IDs should me matched
	require.Nil(t, err)
	require.Equal(t, currency.ID, c.ID)
}

func TestCurrencyRepositoryFindAll(t *testing.T) {

	// given - a memory test database
	db := database.NewDbTest()
	defer db.Close()

	// and - an instance of a currency
	name, acronym := "Real", "BRL"
	currency, _ := models.NewCurrency(name, acronym)

	name2, acronym2 := "Dolar", "USD"
	currency2, _ := models.NewCurrency(name2, acronym2)

	// and - currencies inserted in the database
	repo := repository.NewCurrencyRepository(db)
	repo.Insert(currency)
	repo.Insert(currency2)

	// when - find all currencies
	currencies, err := repo.FindAll()

	// then - there should be no error and two currencies should be returned
	require.Nil(t, err)
	require.Equal(t, 2, len(currencies))

}
