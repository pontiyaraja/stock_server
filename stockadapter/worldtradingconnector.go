package stockadapter

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func doRequest(symbol string) (string, error) {
	//https://www.worldtradingdata.com/api/v1/stock?symbol=AAPL,MSFT,HSBA.L&api_token=xCP7hR80w5FSbbpvuHUhRgqwfKPUoYGZerH3lwMUiuQ5WwnMv6n6l8XawJBb
	resp, err := http.Get(baseURL + symbol + "&api_token=" + token)
	if err != nil {
		return "", err
	}
	respbyte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	respString := string(respbyte)
	if strings.Contains(respString, "Error") {
		return "", errors.New(respString)
	}
	return respString, err
}

func (input *InputStruct) getStockData() (string, error) {
	respData, err := doRequest(input.Symbol)
	if err != nil {
		log.Println("error calling trading api", err)
		return "", err
	}
	return formResponse(respData, input.StockExng)
}

func (tradeData *WorldTradingData) convertToTradeResp(stockExng string) map[string]*StockRespStruct {
	respMap := make(map[string]*StockRespStruct)
	var stockData *StockRespStruct
	for _, data := range tradeData.StockData {
		stockData = new(StockRespStruct)
		stockData.fillStockResponse(data)
		if len(stockExng) > 0 && strings.Contains(stockExng, data.StockExchangeShort) {
			respMap[data.StockExchangeShort] = stockData
		} else if len(stockExng) == 0 {
			respMap[data.StockExchangeShort] = stockData
		}
	}
	return respMap
}
