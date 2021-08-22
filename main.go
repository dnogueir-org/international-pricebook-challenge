package main

import (
	"encoding/json"
	"fmt"

	"github.com/dnogueir-org/international-pricebook-challenge/internal/models"
)

func main() {
	// create currency
	dollar, err := models.NewCurrency("Dolar Americano", "USD")
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonDollar, err := json.MarshalIndent(dollar, "", "\t")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Currency:", string(jsonDollar))
	// create country
	country, err := models.NewCountry("Estados Unidos", "USA", *dollar)
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonCountry, err := json.MarshalIndent(country, "", "\t")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Country:", string(jsonCountry))
	// create product
	product, err := models.NewProduct("Macbook Air M1")
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonProduct, err := json.MarshalIndent(product, "", "\t")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Product:", string(jsonProduct))
	// create price
	productPrice, err := models.NewPrice(999.98, *dollar)
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonPrice, err := json.MarshalIndent(productPrice, "", "\t")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Price:", string(jsonPrice))
	// add price to product
	product.Prices = append(product.Prices, *productPrice)
	jsonProductWithPrice, err := json.MarshalIndent(product, "", "\t")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Product with price:", string(jsonProductWithPrice))

	// convert json
}
