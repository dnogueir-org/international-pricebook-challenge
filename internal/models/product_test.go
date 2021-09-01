package models_test

import (
	"testing"

	"github.com/dnogueir-org/international-pricebook-challenge/internal/models"
	"github.com/stretchr/testify/require"
)

func TestNewProduct(t *testing.T) {
	// when - create a product with standard function
	productName := "Nike Sneakers"
	product, err := models.NewProduct(productName)

	// then - should not return err and fields should be filled
	require.Nil(t, err)
	require.Equal(t, productName, product.Name)
}

func TestValidateProduct(t *testing.T) {
	// given - a product without UUID on Id field
	productName := "Nike Sneakers"
	product := &models.Product{
		ID:   "123",
		Name: productName,
	}

	// when - validate the product
	err := product.Validate()
	// then - there should be an error
	require.NotNil(t, err)
}
