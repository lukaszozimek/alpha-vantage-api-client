package time_series

import "fmt"

func GetWeeklyAdjusted(symbol string, apiKey string, c *Client) *AlphaVantageTimeSeriesApiResponse {
	var result = new(AlphaVantageTimeSeriesApiResponse)
	metadata, resultApi := makeApiCallGet(fmt.Sprintf("https://www.alphavantage.co/query?function=TIME_SERIES_DAILY&symbol=%v&apikey=%v", symbol, apiKey), c)
	mapMetadata(metadata, result)
	mapResultApi(resultApi, result)
	return result
}
