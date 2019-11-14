package cryptocurrency

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func FxRealTimeCall(fromCurrency string, toCurrency string, apiKey string, c *Client) *ExchangeRateResult {

	return makeApiCallGetRealTime(fmt.Sprintf("https://www.alphavantage.co/query?function=CURRENCY_EXCHANGE_RATE&from_currency=%v&to_currency=%v&apikey=%v", fromCurrency, toCurrency, apiKey), c)
}
func makeApiCallGetRealTime(url string, c *Client) *ExchangeRateResult {

	res, e := c.HttpClient.Get(url)
	if e != nil {
		panic(e.Error())
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}
	s, err := getDataRealTime(body)
	if err != nil {
		panic(err.Error())
	}
	return s
}
func getDataRealTime(body []byte) (*ExchangeRateResult, error) {
	var s = new(ExchangeRateResult)
	err := json.Unmarshal(body, &s)
	if err != nil {
		fmt.Println("whoops:", err)
	}
	return s, err
}
