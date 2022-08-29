package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/philcantcode/stonks/htmlproxy"
)

func HTTP_JSON_OpenInsiderFull(w http.ResponseWriter, r *http.Request) {
	ticker := r.FormValue("Ticker")
	tickerParam := mux.Vars(r)["ticker"]

	if ticker == "" {
		ticker = tickerParam
	}

	w.Write(htmlproxy.GetOpenInsiderPage(ticker))
}
