package forex

import "fmt"

func GetWeekly(symbol string, market string, c *Client) *AlphaVantageIntraExchangeRate {
	return makeApiCallGet(fmt.Sprintf("https://www.alphavantage.co/query?function=DIGITAL_CURRENCY_WEEKLY&symbol=%v&market=%v&apikey=demo", symbol, market), c)
}
