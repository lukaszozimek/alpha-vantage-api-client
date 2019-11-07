package forex

func GetWeekly(symbol string, c *Client) *AlphaVantageIntraExchangeRate {
	return makeApiCallGet("https://www.alphavantage.co/query?function=DIGITAL_CURRENCY_WEEKLY&symbol=BTC&market=CNY&apikey=demo", c)
}
