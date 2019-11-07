package forex

func GetWeekly(symbol string, c *Client) *AlphaVantageIntraExchangeRate {
	return makeApiCallGet("https://www.alphavantage.co/query?function=FX_WEEKLY&from_symbol=EUR&to_symbol=USD&apikey=demo", c)
}
