package stockadapter

type WorldTradingData struct {
	SymbolRequested int            `json:"symbols_requested"`
	SymbolReturned  int            `json:"symbols_returned"`
	StockData       []StockDetails `json:"data"`
}

type StockDetails struct {
	Symbol               string `json:"symbol"`
	Name                 string `json:"name"`
	Currency             string `json:"currency"`
	Price                string `json:"price"`
	PriceOpen            string `json:"price_open"`
	DayHigh              string `json:"day_high"`
	DayLow               string `json:"day_low"`
	WeekHigh52           string `json:"52_week_high"`
	WeekLow52            string `json:"52_week_low"`
	DayChange            string `json:"day_change"`
	ChangePercentage     string `json:"change_pct"`
	CloseYesterDay       string `json:"close_yesterday"`
	MarketCapitalization string `json:"market_cap"`
	Volume               string `json:"volume"`
	VolumeAverage        string `json:"volume_avg"`
	Shares               string `json:"shares"`
	StockExchangeLong    string `json:"stock_exchange_long"`
	StockExchangeShort   string `json:"stock_exchange_short"`
	TimeZone             string `json:"timezone"`
	TimeZoneName         string `json:"timezone_name"`
	GMTOffset            string `json:"gmt_offset"`
	LastTradeTime        string `json:"last_trade_time"`
}

type StockRespStruct struct {
	Symbol               string `json:"symbol"`
	Name                 string `json:"name"`
	Price                string `json:"price"`
	CloseYesterDay       string `json:"close_yesterday"`
	Currency             string `json:"currency"`
	MarketCapitalization string `json:"market_cap"`
	Volume               string `json:"volume"`
	TimeZone             string `json:"timezone"`
	TimeZoneName         string `json:"timezone_name"`
	GMTOffset            string `json:"gmt_offset"`
	LastTradeTime        string `json:"last_trade_time"`
}

type Response struct {
	StatusCode int               `json:"statusCode"`
	Headers    map[string]string `json:"headers"`
	Body       interface{}       `json:"body"`
}

type InputStruct struct {
	Symbol    string
	StockExng string
}
