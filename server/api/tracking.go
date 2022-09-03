package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/philcantcode/stonks/server/tracking"
)

func HTTP_JSON_Tracking_GetAll(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(tracking.GetTracked())
}

// Find stock or return new
func HTTP_JSON_Tracking_ByTicker(w http.ResponseWriter, r *http.Request) {
	tickerParam := mux.Vars(r)["ticker"]
	tracked := tracking.GetTracked()

	if tickerParam == "" {
		return
	}

	for _, stock := range tracked {
		if stock.Ticker == tickerParam {
			json.NewEncoder(w).Encode(stock)
			return
		}
	}

	json.NewEncoder(w).Encode(tracking.Stock{})
}

func HTTP_NONE_PutTicker(w http.ResponseWriter, r *http.Request) {
	ticker := r.FormValue("Ticker")
	tracking.TrackNewTicker(ticker)
}

func HTTP_NONE_UpdateTracking(w http.ResponseWriter, r *http.Request) {
	ticker := r.FormValue("ticker")
	company := r.FormValue("company")
	industry := r.FormValue("industry")
	ownedStock := r.FormValue("ownedStock")
	group := r.FormValue("group")
	notes := r.FormValue("notes")

	tracking.UpdateTracking(ticker, company, industry, ownedStock, group, notes)
}
