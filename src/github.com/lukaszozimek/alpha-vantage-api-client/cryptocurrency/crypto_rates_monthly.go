package forex

func GetMonthly(symbol string, c *Client) *AlphaVantageTimeSeriesApiResponse {
	return makeApiCallGet("https://www.alphavantage.co/query?function=DIGITAL_CURRENCY_MONTHLY&symbol=BTC&market=CNY&apikey=demo", c)
}
