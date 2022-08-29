package api

import (
	"encoding/json"
	"net/http"

	"github.com/philcantcode/stonks/tracking"
)

func HTTP_JSON_Tracking_GetAll(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(tracking.GetTracked())
}

func HTTP_NONE_PutTicker(w http.ResponseWriter, r *http.Request) {
	ticker := r.FormValue("Ticker")
	tracking.TrackNewTicker(ticker)
}
