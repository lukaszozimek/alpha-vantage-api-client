package forex

import (
	"encoding/json"
	"fmt"
	common "github.com/lukaszozimek/alpha-vantage-api-client"
	"io/ioutil"
)

type RealTimeCurrencyExchangeMetadata struct {
	Information   string
	Symbol        string
	LastRefreshed string
	OutputSize    string
	TimeZone      string
}

func (c RealTimeCurrencyExchangeMetadata) String() string {
	return fmt.Sprintf("{Information:\"%v\", Symbol:\"%v\",  LastRefreshed:\"%v\",  OutputSize:\"%v\", TimeZone:\"%v\"}",
		c.Information,
		c.Symbol,
		c.LastRefreshed,
		c.OutputSize,
		c.TimeZone)
}

type AlphaVantageRealTimeCurrencyExchange struct {
	Result   []Result
	Metadata RealTimeCurrencyExchangeMetadata
}

type AlphaVantageIntraExchangeRate struct {
	Metadata Metadata
	Result   []Result
}
type AlphaVantageIntervalExchangeRate struct {
	Metadata Metadata
	Result   []*Result
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

func (c Metadata) String() string {
	return fmt.Sprintf("{Information:\"%v\", FromSymbol:\"%v\", ToSymbol:\"%v\", LastRefreshed:\"%v\", Interval:\"%v\", OutputSize:\"%v\", TimeZone:\"%v\"}", c.Information,
		c.FromSymbol,
		c.ToSymbol,
		c.LastRefreshed,
		c.Interval,
		c.OutputSize,
		c.TimeZone)
}

type Result struct {
	Date   string
	Open   string `json:"1. open"`
	High   string `json:"2. high"`
	Low    string `json:"3. low"`
	Close  string `json:"4. close"`
	Volume string `json:"5. volume"`
}

func makeApiCallGet(url string, c *common.Client) *AlphaVantageIntervalExchangeRate {

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
func getData(body []byte) (*AlphaVantageIntervalExchangeRate, error) {
	var transformedModel = new(AlphaVantageIntervalExchangeRate)
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

func mapMetadataAlphaVantageIntraExchangeRate(val interface{}, transformedModel *AlphaVantageIntervalExchangeRate) {
	for nestedKey, nestedVal := range val.(map[string]interface{}) {
		if nestedKey == "1. Information" {
			transformedModel.Metadata.Information = nestedVal.(string)
		}
		if nestedKey == "2. From Symbol" {
			transformedModel.Metadata.FromSymbol = nestedVal.(string)
		}
		if nestedKey == "3. To Symbol" {
			transformedModel.Metadata.ToSymbol = nestedVal.(string)
		}
		if nestedKey == "4. Last Refreshed" {
			transformedModel.Metadata.LastRefreshed = nestedVal.(string)
		}
		if nestedKey == "5. Interval" {
			transformedModel.Metadata.Interval = nestedVal.(string)
		}
		if nestedKey == "6. Output Size" {
			transformedModel.Metadata.TimeZone = nestedVal.(string)
		}
		if nestedKey == "7. Time Zone" {
			transformedModel.Metadata.TimeZone = nestedVal.(string)
		}

	}
}
func mapAdiResultMetadataAlphaVantageIntraExchangeRate(val interface{}, transformedModel *AlphaVantageIntervalExchangeRate) {
	var s []*Result
	for nestedKey, nestedVal := range val.(map[string]interface{}) {
		var result = new(Result)
		result.Date = nestedKey
		for keyResult, resultValue := range nestedVal.(map[string]interface{}) {

			if keyResult == "1. open" {
				result.Open = resultValue.(string)
			}
			if keyResult == "2. high" {
				result.High = resultValue.(string)
			}
			if keyResult == "3. low" {
				result.Low = resultValue.(string)
			}
			if keyResult == "4. close" {
				result.Close = resultValue.(string)
			}
			if keyResult == "5. volume" {
				result.Volume = resultValue.(string)
			}
		}
		s = append(s, result)
	}
	transformedModel.Result = s
}
