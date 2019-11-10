package forex

import "fmt"

func TimeSeriesDaily(fromSymbol string, toSymbol string, apiKey string, c *Client) *AlphaVantageIntervalExchangeRate {
	return makeApiCallGet(fmt.Sprintf("https://www.alphavantage.co/query?function=FX_DAILY&from_symbol=%v&to_symbol=%v&&apikey=%v", fromSymbol, toSymbol, apiKey), c)
}
