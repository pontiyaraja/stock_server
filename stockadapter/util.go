package stockadapter

import (
	"encoding/json"
	"log"
	"strings"
)

const (
	token         = "xCP7hR80w5FSbbpvuHUhRgqwfKPUoYGZerH3lwMUiuQ5WwnMv6n6l8XawJBb"
	baseURL       = "https://www.worldtradingdata.com/api/v1/stock?symbol="
	defultSymbol  = "AXP"
	stockExchange = "stock_exchange"
)

func (stockData *StockRespStruct) fillStockResponse(stockDetail StockDetails) {
	stockData.Symbol = stockDetail.Symbol
	stockData.Name = stockDetail.Name
	stockData.Price = stockDetail.Price
	stockData.MarketCapitalization = stockDetail.MarketCapitalization
	stockData.Currency = stockDetail.Currency
	stockData.CloseYesterDay = stockDetail.CloseYesterDay
	stockData.Volume = stockDetail.Volume
	stockData.TimeZone = stockDetail.TimeZone
	stockData.TimeZoneName = stockDetail.TimeZoneName
	stockData.LastTradeTime = stockDetail.LastTradeTime
}

func formResponse(stockData, stockExchng string) (string, error) {
	tradingData := &WorldTradingData{}
	err := json.NewDecoder(strings.NewReader(stockData)).Decode(tradingData)
	if err != nil {
		log.Println("error decoding json data", err)
		return "", err
	}
	byteAry, err := json.Marshal(tradingData.convertToTradeResp(stockExchng))
	if err != nil {
		return "", err
	}
	return string(byteAry), err
}
