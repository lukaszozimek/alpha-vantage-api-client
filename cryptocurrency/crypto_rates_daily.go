package cryptocurrency

import "fmt"

func timeSeriesDaily(symbol string, market string, apiKey string, c *Client) *AlphaVantageIntraExchangeRate {

	return makeApiCallGet(fmt.Sprintf("https://www.alphavantage.co/query?function=DIGITAL_CURRENCY_DAILY&symbol=%v&market=%v&apikey=%v", symbol, market, apiKey), c)
}
