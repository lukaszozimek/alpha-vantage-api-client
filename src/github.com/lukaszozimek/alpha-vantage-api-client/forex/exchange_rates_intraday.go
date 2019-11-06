package forex

const (
	ONE_MINUTE        = "1min"
	FIVE_MINUTE       = "5min"
	FIFITHTEEN_MINUTE = "15min"
	THIRTY_MINUTE     = "30min"
	SIXTY_MINUTE      = "60min"
)

func TimeSeriesIntraDayInterval1minute(fromSymbol string, toSymbol string, c *Client) *AlphaVantageTimeSeriesApiResponse {

	return timeSeriesIntraDay(fromSymbol, toSymbol, ONE_MINUTE, c)
}
func TimeSeriesIntraDayInterval5minute(fromSymbol string, toSymbol string, c *Client) *AlphaVantageTimeSeriesApiResponse {

	return timeSeriesIntraDay(fromSymbol, toSymbol, FIVE_MINUTE, c)
}
func TimeSeriesIntraDayIntervalFifteenMinute(fromSymbol string, toSymbol string, c *Client) *AlphaVantageTimeSeriesApiResponse {

	return timeSeriesIntraDay(fromSymbol, toSymbol, FIFITHTEEN_MINUTE, c)
}
func TimeSeriesIntraDayIntervalThirtyMinute(fromSymbol string, toSymbol string, c *Client) *AlphaVantageTimeSeriesApiResponse {

	return timeSeriesIntraDay(fromSymbol, toSymbol, THIRTY_MINUTE, c)
}
func TimeSeriesIntraDayIntervalSixtyMinute(fromSymbol string, toSymbol string, c *Client) *AlphaVantageTimeSeriesApiResponse {

	return timeSeriesIntraDay(fromSymbol, toSymbol, SIXTY_MINUTE, c)
}
func timeSeriesIntraDay(fromSymbol string, toSymbol string, interval string, c *Client) *AlphaVantageTimeSeriesApiResponse {

	return makeApiCallGet("https://www.alphavantage.co/query?function=FX_INTRADAY&from_symbol=EUR&to_symbol=USD&interval=5min&apikey=demo", c)
}