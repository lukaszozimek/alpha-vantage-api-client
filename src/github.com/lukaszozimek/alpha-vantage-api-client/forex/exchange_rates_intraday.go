package forex

import "fmt"

const (
	ONE_MINUTE        = "1min"
	FIVE_MINUTE       = "5min"
	FIFITHTEEN_MINUTE = "15min"
	THIRTY_MINUTE     = "30min"
	SIXTY_MINUTE      = "60min"
)

func TimeSeriesIntraDayInterval1minute(fromSymbol string, toSymbol string, apiKey string, c *Client) *AlphaVantageIntervalExchangeRate {

	return timeSeriesIntraDay(fromSymbol, toSymbol, ONE_MINUTE, apiKey, c)
}
func TimeSeriesIntraDayInterval5minute(fromSymbol string, toSymbol string, apiKey string, c *Client) *AlphaVantageIntervalExchangeRate {

	return timeSeriesIntraDay(fromSymbol, toSymbol, FIVE_MINUTE, apiKey, c)
}
func TimeSeriesIntraDayIntervalFifteenMinute(fromSymbol string, toSymbol string, apiKey string, c *Client) *AlphaVantageIntervalExchangeRate {

	return timeSeriesIntraDay(fromSymbol, toSymbol, FIFITHTEEN_MINUTE, apiKey, c)
}
func TimeSeriesIntraDayIntervalThirtyMinute(fromSymbol string, toSymbol string, apiKey string, c *Client) *AlphaVantageIntervalExchangeRate {

	return timeSeriesIntraDay(fromSymbol, toSymbol, THIRTY_MINUTE, apiKey, c)
}
func TimeSeriesIntraDayIntervalSixtyMinute(fromSymbol string, toSymbol string, apiKey string, c *Client) *AlphaVantageIntervalExchangeRate {

	return timeSeriesIntraDay(fromSymbol, toSymbol, SIXTY_MINUTE, apiKey, c)
}
func timeSeriesIntraDay(fromSymbol string, toSymbol string, interval string, apiKey string, c *Client) *AlphaVantageIntervalExchangeRate {

	return makeApiCallGet(fmt.Sprintf("https://www.alphavantage.co/query?function=FX_INTRADAY&from_symbol=%v&to_symbol=%v&interval=%v&apikey=%v", fromSymbol, toSymbol, interval, apiKey), c)
}
