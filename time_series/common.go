package time_series

import (
	"encoding/json"
	"fmt"
	common "github.com/lukaszozimek/alpha-vantage-api-client"
	"io/ioutil"
)

type AlphaVantageTimeSeriesApiResponse struct {
	Metadata Metadata `json:"Meta Data"`
	Result   []Result
}
type Metadata struct {
	Information   string `json:"1. Information"`
	Symbol        string `json:"2. Symbol"`
	LastRefreshed string `json:"3. Last Refreshed"`
	OutputSize    string `json:"5. Output Size"`
	TimeZone      string `json:"6. Time Zone"`
	Interval      string `json:"4. Interval"`
}

func (c Metadata) String() string {
	return fmt.Sprintf("{Information:\"%v\", Symbol:\"%v\", LastRefreshed:\"%v\", OutputSize:\"%v\", TimeZone:\"%v\", Interval:\"%v\"}",
		c.Information,
		c.Symbol,
		c.LastRefreshed,
		c.OutputSize,
		c.TimeZone,
		c.Interval)
}

type Result struct {
	Date   string
	Open   string `json:"1. open"`
	High   string `json:"2. high"`
	Low    string `json:"3. low"`
	Close  string `json:"4. close"`
	Volume string `json:"5. volume"`
}

func makeApiCallGet(url string, c *common.Client) *AlphaVantageTimeSeriesApiResponse {

	res, e := c.HttpClient.Get(url)
	if e != nil {
		panic(e.Error())
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}
	metadata, err := getData(body)
	if err != nil {
		panic(err.Error())
	}
	return metadata
}
func getData(body []byte) (*AlphaVantageTimeSeriesApiResponse, error) {
	var transformedModel = new(AlphaVantageTimeSeriesApiResponse)
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
			mapApiResult(val, transformedModel)
		}

	}
	return transformedModel, nil
}

func mapApiResult(val interface{}, transformedModel *AlphaVantageTimeSeriesApiResponse) {
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

func mapMetadata(val interface{}, transformedModel *AlphaVantageTimeSeriesApiResponse) {
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
		if nestedKey == "5. Output Size" {
			transformedModel.Metadata.OutputSize = nestedVal.(string)
		}
		if nestedKey == "6. Time Zone" {
			transformedModel.Metadata.TimeZone = nestedVal.(string)
		}
		if nestedKey == "4. Interval" {
			transformedModel.Metadata.Interval = nestedVal.(string)
		}

	}
}
