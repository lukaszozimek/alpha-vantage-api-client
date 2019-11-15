package time_series

import "fmt"

func GetDailyQuoteSearchEndpoint(keyWord string, apiKey string, c *Client) *AlphaVantageTimeSeriesApiResponse {
	return makeApiCallGet(fmt.Sprintf(c.BaseURL.String()+"/query?function=SYMBOL_SEARCH&keywords=%v&apikey=%v", keyWord, apiKey), c)

}
