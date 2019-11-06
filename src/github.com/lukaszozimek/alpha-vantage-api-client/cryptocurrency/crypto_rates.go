package forex

import (
	"encoding/json"
	"fmt"
)

func FxRealTimeCall(symbol string, interval string, c *Client) *AlphaVantageTimeSeriesApiResponse {

	makeApiCallGet("https://www.alphavantage.co/query?function=CURRENCY_EXCHANGE_RATE&from_currency=BTC&to_currency=CNY&apikey=demo", c)
}
func getData(body []byte) (*AlphaVantageRealTimeCurrencyExchange, error) {
	var s = new(AlphaVantageRealTimeCurrencyExchange)
	err := json.Unmarshal(body, &s)
	if err != nil {
		fmt.Println("whoops:", err)
	}
	return s, err
}
