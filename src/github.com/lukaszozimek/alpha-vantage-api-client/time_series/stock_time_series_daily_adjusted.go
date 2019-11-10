package time_series

import "fmt"

func GetDailyAdjusted(symbol string, apiKey string, c *Client) *AlphaVantageTimeSeriesApiResponse {
	var result = new(AlphaVantageTimeSeriesApiResponse)
	metadata, resultApi := makeApiCallGet(fmt.Sprintf("https://www.alphavantage.co/query?function=TIME_SERIES_DAILY_ADJUSTED&symbol=%v&apikey=demo", symbol), c)
	mapMetadata(metadata, result)
	mapResultApi(resultApi, result)
	return result
}
