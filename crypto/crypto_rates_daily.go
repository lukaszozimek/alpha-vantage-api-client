package crypto

import (
	"fmt"
	common "github.com/lukaszozimek/alpha-vantage-api-client"
)

func TimeSeriesDaily(symbol string, market string, apiKey string, c *common.Client) *AlphaVantageCryptoExchangeRate {

	return makeApiCallGet(fmt.Sprintf("https://www.alphavantage.co/query?function=DIGITAL_CURRENCY_DAILY&symbol=%v&market=%v&apikey=%v", symbol, market, apiKey), c)
}
