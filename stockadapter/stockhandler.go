package stockadapter

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func StockSymbolHandler(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	symbol := vars["symbol"]
	if len(symbol) == 0 {
		symbol = defultSymbol
	}
	stockExng := req.URL.Query()[stockExchange]
	var exchange string
	if len(stockExng) > 0 {
		log.Println(stockExng[0])
		exchange = stockExng[0]
	}
	input := InputStruct{Symbol: symbol, StockExng: exchange}
	stockData, err := input.getStockData()
	//var resp Response
	if err != nil {
		fmt.Fprint(w, err.Error())
		return
		// resp.StatusCode = http.StatusInternalServerError
		// resp.Body = err
	} //else {
	// 	resp.StatusCode = http.StatusOK
	// 	resp.Body = stockData
	// }
	//respByte, err := json.Marshal(stockData)
	fmt.Fprint(w, stockData)
	return
}
