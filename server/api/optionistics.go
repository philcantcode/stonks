package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/philcantcode/stonks/server/htmlproxy"
)

func HTTP_JSON_OptionicsFull(w http.ResponseWriter, r *http.Request) {
	ticker := r.FormValue("ticker")
	tickerParam := mux.Vars(r)["ticker"]

	if ticker == "" {
		ticker = tickerParam
	}

	w.Write(htmlproxy.GetOptiononicsPage(ticker))
}

func HTTP_JSON_OptiononicsPart(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Part request")

	ticker := r.FormValue("ticker")
	tickerParam := mux.Vars(r)["ticker"]
	target := r.FormValue("target")
	targetParam := mux.Vars(r)["target"]

	if ticker == "" {
		ticker = tickerParam
	}

	if target == "" {
		target = targetParam
	}

	w.Write(htmlproxy.GetOptiononicsPart(ticker, target))
}
