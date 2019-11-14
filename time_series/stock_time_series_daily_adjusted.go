package time_series

import "fmt"

func GetDailyAdjusted(symbol string, apiKey string, c *Client) *AlphaVantageTimeSeriesApiResponse {
	return makeApiCallGet(fmt.Sprintf(c.BaseURL.String()+"/query?function=TIME_SERIES_DAILY_ADJUSTED&symbol=%v&apikey=demo", symbol), c)

}