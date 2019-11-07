package time_series

import "fmt"

func timeSeriesDaily(symbol string, c *Client) *AlphaVantageTimeSeriesApiResponse {

	return makeApiCallGet(fmt.Sprintf("https://www.alphavantage.co/query?function=TIME_SERIES_DAILY&symbol=%v&apikey=demo", symbol), c)
}
