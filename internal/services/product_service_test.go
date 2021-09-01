package services_test

import (
	"testing"

	"github.com/dnogueir-org/international-pricebook-challenge/database"
	"github.com/dnogueir-org/international-pricebook-challenge/internal/models"
	"github.com/dnogueir-org/international-pricebook-challenge/internal/services"
	"github.com/dnogueir-org/international-pricebook-challenge/repository"
	"github.com/stretchr/testify/require"
)

func TestProductService(t *testing.T) {
	// given - a memory test database
	db := database.NewDbTest()
	defer db.Close()

	// and - an instance of a currency
	name, acronym := "Real", "BRL"
	currency, _ := models.NewCurrency(name, acronym)

	// and - inserting the currency in the database
	currencyRepo := repository.NewCurrencyRepository(db)
	currencyRepo.Insert(currency)

	// when - inserting product via service
	productName := "Nike Sneakers"
	productPrice := 255.55
	productRepo := repository.NewProductRepository(db)
	priceRepo := repository.NewPriceRepository(db)
	productService := services.NewProductService(productRepo, priceRepo, currencyRepo)
	product, err := productService.Insert(productName, productPrice, currency.ID)
	insertedPrices := product.Prices

	// then - there should be no error returned and fields should match
	require.Nil(t, err)
	require.Equal(t, product.Name, productName)
	require.Equal(t, productPrice, insertedPrices[0].Value)

}
