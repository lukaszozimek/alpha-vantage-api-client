package time_series

import "fmt"

func GetMonthly(symbol string, apiKey string, c *Client) *AlphaVantageTimeSeriesApiResponse {
	return makeApiCallGet(fmt.Sprintf(c.BaseURL.String()+"/query?function=TIME_SERIES_MONTHLY&symbol=%v&apikey=%v", symbol, apiKey), c)

}
