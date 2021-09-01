package handler

import (
	"encoding/json"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/dnogueir-org/international-pricebook-challenge/internal/services"
	"github.com/gorilla/mux"
)

type CountryRequest struct {
	Name         string `json:"name"`
	Abbreviation string `json:"abbreviation"`
	CurrencyId   string `json:"currency_id"`
}

func MakeCountryHandlers(r *mux.Router, n *negroni.Negroni, service services.CountryService) {
	r.Handle("/country/{id}", n.With(
		negroni.Wrap(getCountry(service)),
	)).Methods("GET", "OPTIONS")

	r.Handle("/countries", n.With(
		negroni.Wrap(getCountries(service)),
	)).Methods("GET", "OPTIONS")

	r.Handle("/country", n.With(
		negroni.Wrap(createCountry(service)),
	)).Methods("POST", "OPTIONS")
}

func getCountry(service services.CountryService) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		vars := mux.Vars(r)
		id := vars["id"]
		country, err := service.Find(id)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		err = json.NewEncoder(w).Encode(country)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	})
}

func getCountries(service services.CountryService) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		countries, err := service.FindAll()
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		err = json.NewEncoder(w).Encode(countries)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	})
}

func createCountry(service services.CountryService) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var CountryRequest CountryRequest
		err := json.NewDecoder(r.Body).Decode(&CountryRequest)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonError(err.Error()))
			return
		}
		currency, err := service.Insert(CountryRequest.Name, CountryRequest.Abbreviation, CountryRequest.CurrencyId)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonError(err.Error()))
			return
		}
		err = json.NewEncoder(w).Encode(currency)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonError(err.Error()))
			return
		}
	})
}
