package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"exporter/internal/usecases/collector"
	"github.com/D3vR4pt0rs/logger"

	"github.com/gorilla/mux"
)

const (
	ticker = "ticker"
)

func getStockInformationByTicker(app collector.Controller) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Info.Printf("Got new request to get stock information by %s\n", mux.Vars(r)[ticker])
		errorMessage := "Error getting ticker"
		ticker := mux.Vars(r)[ticker]
		result, err := app.GetStockInformationByTicker(ticker)
		if err != nil {
			fmt.Println(err.Error())
			http.Error(w, errorMessage, http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		resp := make(map[string]interface{})
		resp["data"] = result

		json.NewEncoder(w).Encode(resp)
	})
}

func getAvailableCompanies(app collector.Controller) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Info.Printf("Got new request to get available companies")
		result := app.GetAvailableCompanies()
		w.Header().Set("Content-Type", "application/json")
		resp := make(map[string]interface{})
		resp["data"] = result
		json.NewEncoder(w).Encode(resp)
	})
}

func Make(r *mux.Router, app collector.Controller) {
	apiUri := "/api"
	serviceRouter := r.PathPrefix(apiUri).Subrouter()
	serviceRouter.Handle("/ticker/{ticker}", getStockInformationByTicker(app)).Methods("GET")
	serviceRouter.Handle("/ticker", getAvailableCompanies(app)).Methods("GET")
}
