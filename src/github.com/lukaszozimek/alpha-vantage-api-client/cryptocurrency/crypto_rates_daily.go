package forex

func timeSeriesDaily(symbol string, c *Client) *AlphaVantageIntraExchangeRate {

	return makeApiCallGet("https://www.alphavantage.co/query?function=DIGITAL_CURRENCY_DAILY&symbol=BTC&market=CNY&apikey=demo", c)
}
