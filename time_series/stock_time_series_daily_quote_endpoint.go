package time_series

import "fmt"

func GetDailyQuoteEndpoint(symbol string, apiKey string, c *Client) *AlphaVantageTimeSeriesApiResponse {
	return makeApiCallGet(fmt.Sprintf(c.BaseURL.String()+"/query?function=GLOBAL_QUOTE&symbol=%v&apikey=%v", symbol, apiKey), c)

}
