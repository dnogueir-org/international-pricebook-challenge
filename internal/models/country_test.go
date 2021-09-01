package models_test

import (
	"testing"

	"github.com/dnogueir-org/international-pricebook-challenge/internal/models"
	"github.com/stretchr/testify/require"
)

func TestNewCountry(t *testing.T) {
	// given - a currency
	currencyName, acronym := "Real", "BRL"
	currency, _ := models.NewCurrency(currencyName, acronym)

	// when - create a crounty
	countryName, abbreviation := "Brasil", "BRA"
	country, err := models.NewCountry(countryName, abbreviation, *currency)

	// then - should not return err and fields should be filled
	require.Nil(t, err)
	require.Equal(t, countryName, country.Name)
	require.Equal(t, abbreviation, country.Abbreviation)
}

func TestValidateCountry(t *testing.T) {

	// given - currency
	currencyName, acronym := "Real", "BRL"
	currency, _ := models.NewCurrency(currencyName, acronym)

	// and country with no abbreviation
	countryName, abbreviation := "Brasil", "BRA"
	country, _ := models.NewCountry(countryName, abbreviation, *currency)
	country.Abbreviation = ""

	// when - country is validated
	err := country.Validate()

	// then - there should be an error
	require.NotNil(t, err)
}
