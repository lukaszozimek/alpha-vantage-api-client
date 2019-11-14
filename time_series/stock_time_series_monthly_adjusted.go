package time_series

import "fmt"

func GetMonthlyAdjusted(symbol string, apiKey string, c *Client) *AlphaVantageTimeSeriesApiResponse {
	return makeApiCallGet(fmt.Sprintf(c.BaseURL.String()+"/query?function=TIME_SERIES_DAILY&symbol=%v&apikey=demo", symbol), c)

}
