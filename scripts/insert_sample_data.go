package scripts

import (
	"os"
	"strconv"

	"github.com/dnogueir-org/international-pricebook-challenge/internal/models"
	"github.com/jinzhu/gorm"
)

func InsertSampleData(dbConnection *gorm.DB) {

	shouldInsertSampleData, _ := strconv.ParseBool(os.Getenv("INSERT_SAMPLE_DATA"))

	if !shouldInsertSampleData {
		return
	}

	dollar, _ := models.NewCurrency("Dolar Americano", "USD")

	real, _ := models.NewCurrency("Real", "BRL")

	dbConnection.Create(dollar)
	dbConnection.Create(real)

	countryUs, _ := models.NewCountry("Estados Unidos", "USA", *dollar)

	dbConnection.Create(countryUs)
	countryBr, _ := models.NewCountry("Brasil", "BRA", *real)

	dbConnection.Create(countryBr)

	macbook, _ := models.NewProduct("Macbook Air M1")

	dbConnection.Create(macbook)

	americanMacPrice, _ := models.NewPrice(macbook.ID, *dollar, 999.39)

	dbConnection.Create(americanMacPrice)

	brazilianMacPrice, _ := models.NewPrice(macbook.ID, *real, 10000.39)

	dbConnection.Create(brazilianMacPrice)

	nikeSneakers, _ := models.NewProduct("Nike Sneakers")

	dbConnection.Create(nikeSneakers)

	brazilianSneakerPrice, _ := models.NewPrice(nikeSneakers.ID, *real, 539.90)

	dbConnection.Create(brazilianSneakerPrice)

}
