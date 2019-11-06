package time_series

func timeSeriesDaily(symbol string, interval string, c *Client) *AlphaVantageTimeSeriesApiResponse {

	return makeApiCallGet("https://www.alphavantage.co/query?function=TIME_SERIES_DAILY&symbol=MSFT&apikey=demo", c)
}
