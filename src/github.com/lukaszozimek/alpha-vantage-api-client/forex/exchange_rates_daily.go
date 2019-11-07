package forex

func timeSeriesDaily(symbol string, c *Client) *AlphaVantageIntraExchangeRate {

	return makeApiCallGet("https://www.alphavantage.co/query?function=FX_DAILY&from_symbol=EUR&to_symbol=USD&apikey=demo", c)
}
