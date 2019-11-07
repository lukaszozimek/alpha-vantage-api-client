package forex

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func FxRealTimeCall(symbol string, interval string, c *Client) *AlphaVantageRealTimeCurrencyExchange {

	return makeApiCallGetRealTime("https://www.alphavantage.co/query?function=TIME_SERIES_DAILY&symbol=MSFT&apikey=demo", c)
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
