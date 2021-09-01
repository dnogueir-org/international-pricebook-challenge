package repository_test

import (
	"testing"

	"github.com/dnogueir-org/international-pricebook-challenge/database"
	"github.com/dnogueir-org/international-pricebook-challenge/internal/models"
	"github.com/dnogueir-org/international-pricebook-challenge/repository"
	"github.com/stretchr/testify/require"
)

func TestPriceRepositoriesDb(t *testing.T) {
	// given - a memory test database
	db := database.NewDbTest()
	defer db.Close()

	// and - an instance of a product
	productName := "Nike Sneakers"
	product, _ := models.NewProduct(productName)

	// and - inserting the product in the database
	productRepo := repository.NewProductRepository(db)
	productRepo.Insert(product)

	// and - an instance of a currency
	name, acronym := "Real", "BRL"
	currency, _ := models.NewCurrency(name, acronym)

	// and - inserting the currency in the database
	currencyRepo := repository.NewCurrencyRepository(db)
	currencyRepo.Insert(currency)

	// and - an instance of a price
	value := 999.98
	price, _ := models.NewPrice(product.ID, *currency, value)

	// when - insert the price in the database
	priceRepo := repository.NewPriceRepository(db)
	_, err := priceRepo.Insert(price)
	// then - no error should be returned
	require.Nil(t, err)

	// when - find the price
	p, err := priceRepo.Find(price.ID)

	// then - no error should be returned and IDs should me matched
	require.Nil(t, err)
	require.Equal(t, price.ID, p.ID)
	require.Equal(t, price.ProductID, p.ProductID)
	require.Equal(t, price.CurrencyID, p.CurrencyID)
}
