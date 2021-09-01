package repository_test

import (
	"testing"

	"github.com/dnogueir-org/international-pricebook-challenge/database"
	"github.com/dnogueir-org/international-pricebook-challenge/internal/models"
	"github.com/dnogueir-org/international-pricebook-challenge/repository"
	"github.com/stretchr/testify/require"
)

func TestProductRepositoriesDb(t *testing.T) {
	// given - a memory test database
	db := database.NewDbTest()
	defer db.Close()

	// and - an instance of a product
	productName := "Nike Sneakers"
	product, _ := models.NewProduct(productName)

	// when - inserting the product in the database
	productRepo := repository.NewProductRepository(db)
	_, err := productRepo.Insert(product)
	// then - no error should be returned
	require.Nil(t, err)

	// when - find the product
	p, err := productRepo.Find(product.ID)

	// then - no error should be returned and IDs should me matched
	require.Nil(t, err)
	require.Equal(t, product.ID, p.ID)
}
