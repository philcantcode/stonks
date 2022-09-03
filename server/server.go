package main

import (
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/philcantcode/stonks/server/api"
	"github.com/philcantcode/stonks/server/system"
)

func initServer() {
	system.Log("Setting up server", true)
	router := mux.NewRouter()

	router.HandleFunc("/get/tracking/all", api.HTTP_JSON_Tracking_GetAll)
	router.HandleFunc("/get/tracking/ticker/{ticker}", api.HTTP_JSON_Tracking_ByTicker)
	router.HandleFunc("/put/tracking/ticker", api.HTTP_NONE_PutTicker)
	router.HandleFunc("/put/update/tracking", api.HTTP_NONE_UpdateTracking)

	router.HandleFunc("/get/yahoo/basic/ticker-lookup", api.HTTP_JSON_Yahoo_StockLookup)
	router.HandleFunc("/get/yahoo/basic/ticker-lookup/{ticker}", api.HTTP_JSON_Yahoo_StockLookup)

	router.HandleFunc("/get/open-insider/part", api.HTTP_JSON_OpenInsiderPart)
	router.HandleFunc("/get/open-insider/part/{ticker}/{target}", api.HTTP_JSON_OpenInsiderPart)

	router.HandleFunc("/get/open-insider/page", api.HTTP_JSON_OpenInsiderFull)
	router.HandleFunc("/get/open-insider/page/{ticker}", api.HTTP_JSON_OpenInsiderFull)

	router.HandleFunc("/get/optiononics/part", api.HTTP_JSON_OptiononicsPart)
	router.HandleFunc("/get/optiononics/part/{ticker}/{target}", api.HTTP_JSON_OptiononicsPart)

	router.HandleFunc("/get/optiononics/page", api.HTTP_JSON_OptionicsFull)
	router.HandleFunc("/get/optiononics/page/{ticker}", api.HTTP_JSON_OptionicsFull)

	cors := handlers.CORS(
		handlers.AllowedHeaders([]string{"content-type"}),
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowCredentials(),
		handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"}),
	)

	fileServer := http.FileServer(http.Dir("/"))

	router.PathPrefix("/").Handler(http.StripPrefix("/", fileServer))

	http.ListenAndServe(":"+system.GetConfig("server-port"), cors(router))
}
