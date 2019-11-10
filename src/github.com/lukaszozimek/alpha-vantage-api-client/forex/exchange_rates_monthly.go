package forex

import "fmt"

func GetMonthly(fromSymbol string, toSymbol string, apiKey string, c *Client) *AlphaVantageIntervalExchangeRate {
	return makeApiCallGet(fmt.Sprintf("https://www.alphavantage.co/query?function=FX_MONTHLY&from_symbol=%v&to_symbol=%v&&apikey=%v", fromSymbol, toSymbol, apiKey), c)

}
