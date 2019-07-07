package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pontiyaraja/stock_server/stockadapter"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/stock/{symbol}", stockadapter.StockSymbolHandler).Methods(http.MethodGet)
	log.Println("stock server listening at 8060")
	http.ListenAndServe(":8060", r)
}
