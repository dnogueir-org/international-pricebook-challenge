package services_test

import (
	"testing"

	"github.com/dnogueir-org/international-pricebook-challenge/database"
	"github.com/dnogueir-org/international-pricebook-challenge/internal/services"
	"github.com/dnogueir-org/international-pricebook-challenge/repository"
	"github.com/stretchr/testify/require"
)

func TestCurrencyService(t *testing.T) {

	// given - a memory test database
	db := database.NewDbTest()
	defer db.Close()

	// and - currency service
	currencyRepo := repository.NewCurrencyRepository(db)
	currencyService := services.NewCurrencyService(currencyRepo)

	// when - insert currency via service
	name, acronym := "Real", "BRL"
	currency, err := currencyService.Insert(name, acronym)

	// then - no error should be returned
	require.Nil(t, err)

	// when - find currency via service
	c, err := currencyService.Find(currency.ID)

	// then - no error should be returned and fields should match
	require.Nil(t, err)
	require.Equal(t, name, c.Name)
	require.Equal(t, acronym, c.Acronym)

}
