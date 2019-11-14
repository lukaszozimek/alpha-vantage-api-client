package forex

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func FxRealTimeCall(symbol string, apiKey string, c *Client) *AlphaVantageRealTimeCurrencyExchange {

	return makeApiCallGetRealTime(fmt.Sprintf(c.BaseURL.String()+"/query?function=TIME_SERIES_DAILY&symbol=%v&apikey=%v", symbol, apiKey), c)
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
	var transformedModel = new(AlphaVantageRealTimeCurrencyExchange)
	//this is Messy API...
	var result map[string]interface{}
	err := json.Unmarshal(body, &result)
	if err != nil {
		fmt.Println("whoops:", err)
	}
	for key, val := range result {
		if key == "Meta Data" {
			mapMetadata(val, transformedModel)
		} else {
			mapApiResultAlphaVantageRealTimeCurrencyExchange(val, transformedModel)
		}

	}
	return transformedModel, nil
}
func mapApiResultAlphaVantageRealTimeCurrencyExchange(val interface{}, transformedModel *AlphaVantageRealTimeCurrencyExchange) {
	var s []Result
	for nestedKey, nestedVal := range val.(map[string]interface{}) {
		var result Result
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

func mapMetadata(val interface{}, transformedModel *AlphaVantageRealTimeCurrencyExchange) {
	for nestedKey, nestedVal := range val.(map[string]interface{}) {
		if nestedKey == "1. Information" {
			transformedModel.Metadata.Information = nestedVal.(string)
		}
		if nestedKey == "2. Symbol" {
			transformedModel.Metadata.Symbol = nestedVal.(string)
		}
		if nestedKey == "3. Last Refreshed" {
			transformedModel.Metadata.LastRefreshed = nestedVal.(string)
		}
		if nestedKey == "4. Output Size" {
			transformedModel.Metadata.OutputSize = nestedVal.(string)
		}
		if nestedKey == "5. Time Zone" {
			transformedModel.Metadata.TimeZone = nestedVal.(string)
		}
	}
}
