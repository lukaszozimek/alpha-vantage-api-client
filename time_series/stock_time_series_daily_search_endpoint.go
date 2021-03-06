package time_series

import (
	"fmt"
	common "github.com/lukaszozimek/alpha-vantage-api-client"
)

func GetDailyQuoteSearchEndpoint(keyWord string, apiKey string, c *common.Client) *AlphaVantageTimeSeriesApiResponse {
	return makeApiCallGet(fmt.Sprintf(c.BaseURL.String()+"/query?function=SYMBOL_SEARCH&keywords=%v&apikey=%v", keyWord, apiKey), c)

}
