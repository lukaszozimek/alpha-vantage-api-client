package crypto

import "fmt"

func GetMonthly(symbol string, market string, apiKey string, c *Client) *AlphaVantageCryptoExchangeRate {
	return makeApiCallGet(fmt.Sprintf("https://www.alphavantage.co/query?function=DIGITAL_CURRENCY_MONTHLY&symbol=%v&market=%v&apikey=%v", symbol, market, apiKey), c)
}
