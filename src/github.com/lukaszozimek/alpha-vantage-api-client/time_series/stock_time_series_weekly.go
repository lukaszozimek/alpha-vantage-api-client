package time_series

import "fmt"

func GetWeekly(symbol string, apiKey string, c *Client) *AlphaVantageTimeSeriesApiResponse {
	var result = new(AlphaVantageTimeSeriesApiResponse)
	metadata, resultApi := makeApiCallGet(fmt.Sprintf("https://www.alphavantage.co/query?function=TIME_SERIES_WEEKLY&symbol=%v&apikey=%v", symbol, apiKey), c)
	mapMetadata(metadata, result)
	mapResultApi(resultApi, result)
	return result
}
