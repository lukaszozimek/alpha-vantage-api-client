package time_series

import "fmt"

const (
	ONE_MINUTE        = "1min"
	FIVE_MINUTE       = "5min"
	FIFITHTEEN_MINUTE = "15min"
	THIRTY_MINUTE     = "30min"
	SIXTY_MINUTE      = "60min"
)

func TimeSeriesIntraDayInterval1minute(symbol string, c *Client) *AlphaVantageTimeSeriesApiResponse {

	return timeSeriesIntraDay(symbol, ONE_MINUTE, c)
}
func TimeSeriesIntraDayInterval5minute(symbol string, c *Client) *AlphaVantageTimeSeriesApiResponse {

	return timeSeriesIntraDay(symbol, FIVE_MINUTE, c)
}
func TimeSeriesIntraDayIntervalFifteenMinute(symbol string, c *Client) *AlphaVantageTimeSeriesApiResponse {

	return timeSeriesIntraDay(symbol, FIFITHTEEN_MINUTE, c)
}
func TimeSeriesIntraDayIntervalThirtyMinute(symbol string, c *Client) *AlphaVantageTimeSeriesApiResponse {

	return timeSeriesIntraDay(symbol, THIRTY_MINUTE, c)
}
func TimeSeriesIntraDayIntervalSixtyMinute(symbol string, c *Client) *AlphaVantageTimeSeriesApiResponse {

	return timeSeriesIntraDay(symbol, SIXTY_MINUTE, c)
}
func timeSeriesIntraDay(symbol string, interval string, c *Client) *AlphaVantageTimeSeriesApiResponse {

	return makeApiCallGet(fmt.Sprintf("https://www.alphavantage.co/query?function=TIME_SERIES_INTRADAY&symbol=%v&interval%v&apikey=demo", symbol, interval), c)
}
