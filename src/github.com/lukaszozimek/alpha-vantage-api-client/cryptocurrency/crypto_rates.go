package forex

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func FxRealTimeCall(fromCurrency string, toCurrency string, c *Client) *AlphaVantageRealTimeCurrencyExchange {

	return makeApiCallGetRealTime(fmt.Sprintf("https://www.alphavantage.co/query?function=CURRENCY_EXCHANGE_RATE&from_currency=%v&to_currency=%v&apikey=demo", fromCurrency, toCurrency), c)
}
func makeApiCallGetRealTime(url string, c *Client) *AlphaVantageRealTimeCurrencyExchange {

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
func getDataRealTime(body []byte) (*AlphaVantageRealTimeCurrencyExchange, error) {
	var s = new(AlphaVantageRealTimeCurrencyExchange)
	err := json.Unmarshal(body, &s)
	if err != nil {
		fmt.Println("whoops:", err)
	}
	return s, err
}
