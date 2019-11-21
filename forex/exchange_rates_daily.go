package forex

import (
	"fmt"
	common "github.com/lukaszozimek/alpha-vantage-api-client"
)

func TimeSeriesDaily(fromSymbol string, toSymbol string, apiKey string, c *common.Client) *AlphaVantageIntervalExchangeRate {
	return makeApiCallGet(fmt.Sprintf(c.BaseURL.String()+"/query?function=FX_DAILY&from_symbol=%v&to_symbol=%v&apikey=%v", fromSymbol, toSymbol, apiKey), c)
}
