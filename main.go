package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/dnogueir-org/international-pricebook-challenge/cmd/server"
	"github.com/dnogueir-org/international-pricebook-challenge/database"
	"github.com/dnogueir-org/international-pricebook-challenge/internal"
	"github.com/dnogueir-org/international-pricebook-challenge/internal/services"
	"github.com/dnogueir-org/international-pricebook-challenge/repository"
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
	/*
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

		americanMacPrice, err := models.NewPrice(999.39, *dollar, *macbook)
		if err != nil {
			fmt.Println(err.Error())
		}
		dbConnection.Create(americanMacPrice)
		brazilianMacPrice, err := models.NewPrice(10000.39, *real, *macbook)
		if err != nil {
			fmt.Println(err.Error())
		}
		dbConnection.Create(brazilianMacPrice)

		nikeSneakers, err := models.NewProduct("Nike Sneakers")
		if err != nil {
			fmt.Println(err.Error())
		}
		dbConnection.Create(nikeSneakers)

		brazilianSneakerPrice, err := models.NewPrice(539.90, *real, *nikeSneakers)
		if err != nil {
			fmt.Println(err.Error())
		}
		dbConnection.Create(brazilianSneakerPrice)
	*/

	productRepository := repository.NewProductRepository(dbConnection)
	productService := services.NewProductService(productRepository)
	webServer := server.MakeNewWebserver()
	webServer.ProductService = *productService
	fmt.Println("Webserver has been started")
	webServer.Serve()

}
