package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/dnogueir-org/international-pricebook-challenge/cmd/web/server"
	"github.com/dnogueir-org/international-pricebook-challenge/database"
	"github.com/dnogueir-org/international-pricebook-challenge/internal"
	"github.com/dnogueir-org/international-pricebook-challenge/internal/services"
	"github.com/dnogueir-org/international-pricebook-challenge/repository"
	"github.com/dnogueir-org/international-pricebook-challenge/scripts"
	"github.com/joho/godotenv"
)

var db *database.Database

func init() {
	err := godotenv.Load()
	if err != nil {
		internal.Logger.Fatal("error loading .env file")
	}

	autoMmigrateDb, err := strconv.ParseBool(os.Getenv("AUTO_MIGRATE_DB"))
	if err != nil {
		internal.Logger.Fatal("error parsing boolean env var")
	}

	debug, err := strconv.ParseBool(os.Getenv("DEBUG"))
	if err != nil {
		internal.Logger.Fatal("error parsing boolean env var")
	}

	db = database.NewDb(
		os.Getenv("DSN"),
		os.Getenv("DB_TYPE"),
		debug,
		autoMmigrateDb,
		os.Getenv("ENV"))
}

func main() {
	dbConnection, err := db.Connect()
	if err != nil {
		internal.Logger.Fatal("error connecting to DB")
	}
	defer dbConnection.Close()

	scripts.InsertSampleData(dbConnection)

	productRepository := repository.NewProductRepository(dbConnection)
	priceRepository := repository.NewPriceRepository(dbConnection)
	currencyRepository := repository.NewCurrencyRepository(dbConnection)
	countryRepository := repository.NewCountryRepository(dbConnection)
	productService := services.NewProductService(productRepository, priceRepository, currencyRepository)
	currencyService := services.NewCurrencyService(currencyRepository)
	countryService := services.NewCountryService(countryRepository, currencyRepository)
	priceService := services.NewPriceService(priceRepository, currencyRepository)
	webServer := server.MakeNewWebserver()
	webServer.CurrencyService = *currencyService
	webServer.ProductService = *productService
	webServer.CountryService = *countryService
	webServer.PriceService = *priceService
	fmt.Println("Webserver has been started")
	webServer.Serve()

}
