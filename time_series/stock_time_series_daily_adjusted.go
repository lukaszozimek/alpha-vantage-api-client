package time_series

import (
	"fmt"
	common "github.com/lukaszozimek/alpha-vantage-api-client"
)

func GetDailyAdjusted(symbol string, apiKey string, c *common.Client) *AlphaVantageTimeSeriesApiResponse {
	return makeApiCallGet(fmt.Sprintf(c.BaseURL.String()+"/query?function=TIME_SERIES_DAILY_ADJUSTED&symbol=%v&apikey=%v", symbol, apiKey), c)

}
