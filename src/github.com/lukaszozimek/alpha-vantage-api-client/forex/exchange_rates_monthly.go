package forex

func GetMonthly(symbol string, c *Client) *AlphaVantageTimeSeriesApiResponse {
	return makeApiCallGet("https://www.alphavantage.co/query?function=FX_MONTHLY&from_symbol=EUR&to_symbol=USD&apikey=demo", c)
}
