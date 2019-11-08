package time_series

import "fmt"

func GetWeekly(symbol string, apiKey string, c *Client) *AlphaVantageTimeSeriesApiResponse {

	return makeApiCallGet(fmt.Sprintf("https://www.alphavantage.co/query?function=TIME_SERIES_WEEKLY&symbol=%v&apikey=demo", symbol), c)

}
