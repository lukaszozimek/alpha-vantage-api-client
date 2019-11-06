package time_series

func GetDailyAdjusted(symbol string, c *Client) *AlphaVantageTimeSeriesApiResponse {

	return makeApiCallGet("https://www.alphavantage.co/query?function=TIME_SERIES_DAILY_ADJUSTED&symbol=MSFT&apikey=demo", c)

}
