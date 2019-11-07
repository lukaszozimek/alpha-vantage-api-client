package forex

import "fmt"

func GetWeekly(fromSymbol string, toSymbol string, c *Client) *AlphaVantageIntraExchangeRate {
	return makeApiCallGet(fmt.Sprintf("https://www.alphavantage.co/query?function=FX_WEEKLY&from_symbol=%v&to_symbol=%v&&apikey=demo", fromSymbol, toSymbol), c)
}
