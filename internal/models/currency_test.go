package models_test

import (
	"testing"

	"github.com/dnogueir-org/international-pricebook-challenge/internal/models"
	"github.com/stretchr/testify/require"
)

func TestNewCurrency(t *testing.T) {
	name, acronym := "Real", "BRL"
	currency, err := models.NewCurrency(name, acronym)

	require.Nil(t, err)
	require.Equal(t, name, currency.Name)
	require.Equal(t, acronym, currency.Acronym)
}

func TestValidateCurrency(t *testing.T) {
	name, acronym := "Real", "BRL"
	currency := &models.Currency{
		ID:      "123",
		Name:    name,
		Acronym: acronym,
	}

	err := currency.Validate()
	require.NotNil(t, err)
}
