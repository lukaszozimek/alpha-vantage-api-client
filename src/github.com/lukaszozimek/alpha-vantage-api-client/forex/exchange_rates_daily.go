package forex

import "fmt"

func timeSeriesDaily(fromSymbol string, toSymbol string, apiKey string, c *Client) *AlphaVantageIntraExchangeRate {
	return makeApiCallGet(fmt.Sprintf("https://www.alphavantage.co/query?function=FX_DAILY&from_symbol=%v&to_symbol=%v&&apikey=%v", fromSymbol, toSymbol, apiKey), c)
}
