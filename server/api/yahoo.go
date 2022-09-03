package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/philcantcode/stonks/server/system"
	"github.com/philcantcode/stonks/server/yahoo"
)

func HTTP_JSON_Yahoo_StockLookup(w http.ResponseWriter, r *http.Request) {
	ticker := r.FormValue("Ticker")
	tickerParam := mux.Vars(r)["ticker"]

	if ticker == "" {
		ticker = tickerParam
	}

	system.Log("Querying Yahoo Finance for "+ticker, true)

	stock := yahoo.QuerySymbol(ticker)
	stock.INSERT_Stock()

	json.NewEncoder(w).Encode(stock)
}
