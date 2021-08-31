package handler

import (
	"encoding/json"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/dnogueir-org/international-pricebook-challenge/internal/services"
	"github.com/gorilla/mux"
)

type CurrencyRequest struct {
	Name    string `json:"name"`
	Acronym string `json:"acronym"`
}

func MakeCurrencyHandlers(r *mux.Router, n *negroni.Negroni, service services.CurrencyService) {
	r.Handle("/currency/{id}", n.With(
		negroni.Wrap(getCurrency(service)),
	)).Methods("GET", "OPTIONS")

	r.Handle("/currencies", n.With(
		negroni.Wrap(getCurrencies(service)),
	)).Methods("GET", "OPTIONS")

	r.Handle("/currency", n.With(
		negroni.Wrap(createCurrency(service)),
	)).Methods("POST", "OPTIONS")
}

func getCurrency(service services.CurrencyService) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		vars := mux.Vars(r)
		id := vars["id"]
		currency, err := service.Find(id)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		err = json.NewEncoder(w).Encode(currency)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	})
}

func getCurrencies(service services.CurrencyService) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		currencies, err := service.FindAll()
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		err = json.NewEncoder(w).Encode(currencies)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	})
}

func createCurrency(service services.CurrencyService) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var CurrencyRequest CurrencyRequest
		err := json.NewDecoder(r.Body).Decode(&CurrencyRequest)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonError(err.Error()))
			return
		}
		currency, err := service.Insert(CurrencyRequest.Name, CurrencyRequest.Acronym)
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
