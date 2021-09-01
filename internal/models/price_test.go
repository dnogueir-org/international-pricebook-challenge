package models_test

import (
	"testing"

	"github.com/dnogueir-org/international-pricebook-challenge/internal/models"
	"github.com/stretchr/testify/require"
)

func TestNewPrice(t *testing.T) {
	// given - a product
	productName := "Nike Sneakers"
	product, _ := models.NewProduct(productName)

	// and a currency
	currencyName, acronym := "Real", "BRL"
	currency, _ := models.NewCurrency(currencyName, acronym)

	// when - creating new price
	value := 999.98
	price, err := models.NewPrice(product.ID, *currency, value)

	// then - should not return error and fields should be filled
	require.Nil(t, err)
	require.Equal(t, product.ID, price.ProductID)
	require.Equal(t, *currency, price.Currency)
	require.Equal(t, value, price.Value)
}

func TestValidatePrice(t *testing.T) {
	// given - a product
	productName := "Nike Sneakers"
	product, _ := models.NewProduct(productName)

	// and a currency
	currencyName, acronym := "Real", "BRL"
	currency, _ := models.NewCurrency(currencyName, acronym)

	// when - creating new price
	value := 999.98
	price, _ := models.NewPrice(product.ID, *currency, value)

	// when - remove the product and validate
	price.ProductID = ""
	err := price.Validate()
	// then - there should be an error
	require.NotNil(t, err)
}
