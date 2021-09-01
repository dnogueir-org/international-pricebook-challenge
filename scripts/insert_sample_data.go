package scripts

import (
	"fmt"
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

	dollar, err := models.NewCurrency("Dolar Americano", "USD")
	if err != nil {
		fmt.Println(err.Error())
	}
	real, err := models.NewCurrency("Real", "BRL")
	if err != nil {
		fmt.Println(err.Error())
	}

	dbConnection.Create(dollar)
	dbConnection.Create(real)

	countryUs, err := models.NewCountry("Estados Unidos", "USA", *dollar)
	if err != nil {
		fmt.Println(err.Error())
	}
	dbConnection.Create(countryUs)
	countryBr, err := models.NewCountry("Brasil", "BRA", *real)
	if err != nil {
		fmt.Println(err.Error())
	}
	dbConnection.Create(countryBr)

	macbook, err := models.NewProduct("Macbook Air M1")
	if err != nil {
		fmt.Println(err.Error())
	}
	dbConnection.Create(macbook)

	americanMacPrice, err := models.NewPrice(macbook.ID, *dollar, 999.39)
	if err != nil {
		fmt.Println(err.Error())
	}
	dbConnection.Create(americanMacPrice)

	brazilianMacPrice, err := models.NewPrice(macbook.ID, *real, 10000.39)
	if err != nil {
		fmt.Println(err.Error())
	}
	dbConnection.Create(brazilianMacPrice)

	nikeSneakers, err := models.NewProduct("Nike Sneakers")
	if err != nil {
		fmt.Println(err.Error())
	}
	dbConnection.Create(nikeSneakers)

	brazilianSneakerPrice, err := models.NewPrice(nikeSneakers.ID, *real, 539.90)
	if err != nil {
		fmt.Println(err.Error())
	}
	dbConnection.Create(brazilianSneakerPrice)

}
