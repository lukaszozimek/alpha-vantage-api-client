package time_series

import "fmt"

const (
	ONE_MINUTE        = "1min"
	FIVE_MINUTE       = "5min"
	FIFITHTEEN_MINUTE = "15min"
	THIRTY_MINUTE     = "30min"
	SIXTY_MINUTE      = "60min"
)

func TimeSeriesIntraDayInterval1minute(symbol string, apiKey string, c *Client) *AlphaVantageTimeSeriesApiResponse {

	return timeSeriesIntraDay(symbol, ONE_MINUTE, apiKey, c)
}
func TimeSeriesIntraDayInterval5minute(symbol string, apiKey string, c *Client) *AlphaVantageTimeSeriesApiResponse {

	return timeSeriesIntraDay(symbol, FIVE_MINUTE, apiKey, c)
}
func TimeSeriesIntraDayIntervalFifteenMinute(symbol string, apiKey string, c *Client) *AlphaVantageTimeSeriesApiResponse {

	return timeSeriesIntraDay(symbol, FIFITHTEEN_MINUTE, apiKey, c)
}
func TimeSeriesIntraDayIntervalThirtyMinute(symbol string, apiKey string, c *Client) *AlphaVantageTimeSeriesApiResponse {

	return timeSeriesIntraDay(symbol, THIRTY_MINUTE, apiKey, c)
}
func TimeSeriesIntraDayIntervalSixtyMinute(symbol string, apiKey string, c *Client) *AlphaVantageTimeSeriesApiResponse {

	return timeSeriesIntraDay(symbol, SIXTY_MINUTE, apiKey, c)
}
func timeSeriesIntraDay(symbol string, interval string, apiKey string, c *Client) *AlphaVantageTimeSeriesApiResponse {
	return makeApiCallGet(fmt.Sprintf(c.BaseURL.String()+"/query?function=TIME_SERIES_INTRADAY&symbol=%v&interval=%v&apikey=%v", symbol, interval, apiKey), c)

}
