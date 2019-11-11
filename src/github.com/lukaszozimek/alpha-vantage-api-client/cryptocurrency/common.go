package cryptocurrency

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
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
type AlphaVantageCryptoExchangeRate struct {
	Metadata Metadata
	Result   []*Result
}
type Metadata struct {
	Information         string //"1. Information ": "Monthly Prices and Volumes for Digital Currency",
	DigitalCurrencyCode string //"2.Digital Currency Code ": "BTC",
	DigitalCurrencyName string //"3.Digital Currency Name ": "Bitcoin",
	MarketCode          string //"4.Market Code ": "CNY",
	MarketName          string //"5.Market Name ": "Chinese Yuan",
	LastRefreshed       string //"6.Last Refreshed ": "2019-11-11 00:00:00",
	TimeZone            string //"7.Time Zone": "UTC"
}
type Result struct {
	Date           string
	OpenCurrencyA  string //"1a. open (CNY)": "63933.00301200",
	OpenCurrencyB  string //"1b. open (USD)": "9140.86000000",
	HighCurrencyA  string //"2a. high (CNY)": "66540.58065600",
	HighCurrencyB  string //"2b. high (USD)": "9513.68000000",
	LowCurrencyA   string //"3a. low (CNY)": "60821.56320000",
	LowCurrencyB   string //"3b. low (USD)": "8696.00000000",
	CloseCurrencyA string //"4a. close (CNY)": "63239.38819800",
	CloseCurrencyB string //"4b. close (USD)": "9041.69000000",
	Volume         string //"5. volume": "396312.86658300",
	MarketCap      string //"6. market cap (USD)": "396312.86658300"
}

func makeApiCallGet(url string, c *Client) *AlphaVantageCryptoExchangeRate {

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
func getData(body []byte) (*AlphaVantageCryptoExchangeRate, error) {
	var transformedModel = new(AlphaVantageCryptoExchangeRate)
	//this is Messy API...
	var result map[string]interface{}
	err := json.Unmarshal(body, &result)
	if err != nil {
		fmt.Println("whoops:", err)
	}
	for key, val := range result {
		if key == "Meta Data" {
			mapMetadataAlphaVantageIntraExchangeRate(val, transformedModel)
		} else {
			mapAdiResultMetadataAlphaVantageIntraExchangeRate(val, transformedModel)
		}

	}
	return transformedModel, nil
}

func mapMetadataAlphaVantageIntraExchangeRate(val interface{}, transformedModel *AlphaVantageCryptoExchangeRate) {
	for nestedKey, nestedVal := range val.(map[string]interface{}) {
		if nestedKey == "1. Information" {
			transformedModel.Metadata.Information = nestedVal.(string)
		}
		if nestedKey == "2. Digital Currency Code" {
			transformedModel.Metadata.DigitalCurrencyCode = nestedVal.(string)
		}
		if nestedKey == "3. Digital Currency Name" {
			transformedModel.Metadata.DigitalCurrencyName = nestedVal.(string)
		}
		if nestedKey == "4. Market Code" {
			transformedModel.Metadata.MarketCode = nestedVal.(string)
		}
		if nestedKey == "5. Market Name" {
			transformedModel.Metadata.MarketName = nestedVal.(string)
		}
		if nestedKey == "6. Last Refreshed" {
			transformedModel.Metadata.LastRefreshed = nestedVal.(string)
		}
		if nestedKey == "7. Time Zone" {
			transformedModel.Metadata.TimeZone = nestedVal.(string)
		}

	}
}
func mapAdiResultMetadataAlphaVantageIntraExchangeRate(val interface{}, transformedModel *AlphaVantageCryptoExchangeRate) {
	var s []*Result
	for nestedKey, nestedVal := range val.(map[string]interface{}) {
		var result = new(Result)
		result.Date = nestedKey
		for keyResult, resultValue := range nestedVal.(map[string]interface{}) {

			if strings.HasPrefix(keyResult, "1a.") {
				result.OpenCurrencyA = resultValue.(string)
			}
			if strings.HasPrefix(keyResult, "1b.") {
				result.OpenCurrencyB = resultValue.(string)
			}
			if strings.HasPrefix(keyResult, "2a.") {
				result.HighCurrencyA = resultValue.(string)
			}
			if strings.HasPrefix(keyResult, "2b.") {
				result.HighCurrencyB = resultValue.(string)
			}
			if strings.HasPrefix(keyResult, "3a.") {
				result.LowCurrencyA = resultValue.(string)
			}
			if strings.HasPrefix(keyResult, "3b.") {
				result.LowCurrencyB = resultValue.(string)
			}
			if strings.HasPrefix(keyResult, "4a.") {
				result.CloseCurrencyA = resultValue.(string)
			}
			if strings.HasPrefix(keyResult, "4b.") {
				result.CloseCurrencyB = resultValue.(string)
			}
			if strings.HasPrefix(keyResult, "5.") {
				result.Volume = resultValue.(string)
			}
			if strings.HasPrefix(keyResult, "6.") {
				result.MarketCap = resultValue.(string)
			}
		}
		s = append(s, result)
	}
	transformedModel.Result = s
}
