package handler

import (
	"encoding/json"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/dnogueir-org/international-pricebook-challenge/internal/services"
	"github.com/gorilla/mux"
)

type PriceRequest struct {
	ProductId  string  `json:"product_id"`
	CurrencyId string  `json:"currency_id"`
	Price      float64 `json:"price"`
}

func MakePriceHandlers(r *mux.Router, n *negroni.Negroni, service services.PriceService) {
	r.Handle("/price", n.With(
		negroni.Wrap(createPrice(service)),
	)).Methods("POST", "OPTIONS")
}

func createPrice(service services.PriceService) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var PriceRequest PriceRequest
		err := json.NewDecoder(r.Body).Decode(&PriceRequest)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonError(err.Error()))
			return
		}
		product, err := service.Insert(PriceRequest.ProductId, PriceRequest.CurrencyId, PriceRequest.Price)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonError(err.Error()))
			return
		}
		err = json.NewEncoder(w).Encode(product)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonError(err.Error()))
			return
		}
	})
}
