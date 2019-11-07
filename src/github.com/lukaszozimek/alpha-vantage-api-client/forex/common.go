package forex

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Client struct {
	BaseURL    *url.URL
	UserAgent  string
	HttpClient *http.Client
}
type AlphaVantageRealTimeCurrencyExchange struct {
	ExchangeRateResult ExchangeRateResult
}
type ExchangeRateResult struct {
	FromCurrencyCode string  `json:"1. From_Currency Code"`
	FromCurrencyName string  `json:"2 From_Currency Name"`
	ToCurrencyCode   string  `json:"3. To_Currency Code"`
	ToCurrencyName   string  `json:"4. To_Currency Name"`
	ExchangeRate     string  `json:"5. Exchange Rate"`
	LastRefreshed    string  `json:"6. Last Refreshed"`
	TimeZone         string  `json:"7. Time Zone"`
	BadPrice         float32 `json:"8. Bid Price"`
	AskPrice         float32 `json:"9. Ask Price"`
}
type AlphaVantageIntraExchangeRate struct {
	ExchangeRateResult ExchangeRateResult
}
type Metadata struct {
	Information   string `json:"1. Information"`
	FromSymbol    string `json:"2. From Symbol"`
	ToSymbol      string `json:"3. To Symbol"`
	LastRefreshed string `json:"4. Last Refreshed"`
	Interval      string `json:"5. Interval"`
	OutputSize    string `json:"6. Output Size"`
	TimeZone      string `json:"7. Time Zone"`
}
type Result struct {
	Open  float32 `json:"1. open"`
	High  float32 `json:"2. high"`
	Low   float32 `json:"3. low"`
	Close float32 `json:"4. close"`
}

func makeApiCallGet(url string, c *Client) *AlphaVantageIntraExchangeRate {

	res, e := c.HttpClient.Get(url)
	if e != nil {
		panic(e.Error())
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}
	s, err := getData(body)
	if err != nil {
		panic(err.Error())
	}
	return s
}
func getData(body []byte) (*AlphaVantageIntraExchangeRate, error) {
	var s = new(AlphaVantageIntraExchangeRate)
	err := json.Unmarshal(body, &s)
	if err != nil {
		fmt.Println("whoops:", err)
	}
	return s, err
}
