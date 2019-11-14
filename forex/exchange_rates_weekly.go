package forex

import "fmt"

func GetWeekly(fromSymbol string, toSymbol string, apiKey string, c *Client) *AlphaVantageIntervalExchangeRate {
	return makeApiCallGet(fmt.Sprintf(c.BaseURL.String()+"/query?function=FX_WEEKLY&from_symbol=%v&to_symbol=%v&apikey=%v", fromSymbol, toSymbol, apiKey), c)
}