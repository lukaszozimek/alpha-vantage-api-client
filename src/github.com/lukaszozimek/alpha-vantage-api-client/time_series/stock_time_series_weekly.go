package time_series

import "fmt"

func GetWeekly(symbol string, apiKey string, c *Client) *AlphaVantageTimeSeriesApiResponse {
	return makeApiCallGet(fmt.Sprintf(c.BaseURL.String()+"/query?function=TIME_SERIES_WEEKLY&symbol=%v&apikey=%v", symbol, apiKey), c)

}
