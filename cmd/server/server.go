package server

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/codegangsta/negroni"
	"github.com/dnogueir-org/international-pricebook-challenge/cmd/handler"
	"github.com/dnogueir-org/international-pricebook-challenge/internal/services"
	"github.com/gorilla/mux"
)

type Webserver struct {
	ProductService  services.ProductService
	PriceService    services.PriceService
	CountryService  services.CountryService
	CurrencyService services.CurrencyService
}

func MakeNewWebserver() *Webserver {
	return &Webserver{}
}

func (w Webserver) Serve() {
	r := mux.NewRouter()
	n := negroni.New(
		negroni.NewLogger(),
	)
	handler.MakeProductHandlers(r, n, w.ProductService)
	handler.MakeCurrencyHandlers(r, n, w.CurrencyService)
	handler.MakeCountryHandlers(r, n, w.CountryService)
	handler.MakePriceHandlers(r, n, w.PriceService)
	http.Handle("/", r)
	server := &http.Server{
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		Addr:              ":9000",
		Handler:           http.DefaultServeMux,
		ErrorLog:          log.New(os.Stderr, "log: ", log.Lshortfile),
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
