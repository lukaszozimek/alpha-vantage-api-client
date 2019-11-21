package technical_indicator

import (
	"encoding/json"
	"fmt"
	common "github.com/lukaszozimek/alpha-vantage-api-client"
	"io/ioutil"
)

type AlphaVantageTechnicalIndicatorResponse struct {
	Metadata AlphaVantageTechnicalIndicatorMetadata `json:"Metadata"`
	Result   []*Result
}
type AlphaVantageTechnicalIndicatorMetadata struct {
	Symbol        string `json:"1: Symbol"`
	Indicator     string `json:"2: Indicator"`
	LastRefreshed string `json:"3: Last Refreshed"`
	Interval      string `json:"4: Interval"`
	TimePeriod    string `json:"5: Time Period"`
	SeriesType    string `json:"6: Series Type"`
	TimeZone      string `json:"7: Time Zone"`
}

func (c AlphaVantageTechnicalIndicatorMetadata) String() string {
	return fmt.Sprintf("{Symbol:\"%v\", Indicator:\"%v\",  LastRefreshed:\"%v\",  Interval:\"%v\", TimePeriod:\"%v\", SeriesType:\"%v\", TimeZone:\"%v\"}",
		c.Symbol,
		c.Indicator,
		c.LastRefreshed,
		c.Interval,
		c.TimePeriod,
		c.SeriesType,
		c.TimeZone)
}

type Result struct {
	Value     string
	Indicator string
	Date      string
}

const (
	ONE_MINUTE        = "1min"
	FIVE_MINUTE       = "5min"
	FIFITHTEEN_MINUTE = "15min"
	THIRTY_MINUTE     = "30min"
	SIXTY_MINUTE      = "60min"
	DAILY             = "daily"
	WEEKLY            = "weekly"
	MONTHLY           = "monthly"
	CLOSE             = "close"
	OPEN              = "open"
	HIGH              = "high"
	LOW               = "low"
)

func makeApiCallGet(url string, c *common.Client) *AlphaVantageTechnicalIndicatorResponse {

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
func getData(body []byte) (*AlphaVantageTechnicalIndicatorResponse, error) {
	var transformedModel = new(AlphaVantageTechnicalIndicatorResponse)
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

func mapMetadataAlphaVantageIntraExchangeRate(val interface{}, transformedModel *AlphaVantageTechnicalIndicatorResponse) {
	for nestedKey, nestedVal := range val.(map[string]interface{}) {
		if nestedKey == "1: Symbol" {
			transformedModel.Metadata.Symbol = nestedVal.(string)
		}
		if nestedKey == "2: Indicator" {
			transformedModel.Metadata.Indicator = nestedVal.(string)
		}
		if nestedKey == "3: Last Refreshed" {
			transformedModel.Metadata.LastRefreshed = nestedVal.(string)
		}
		if nestedKey == "4: Interval" {
			transformedModel.Metadata.Interval = nestedVal.(string)
		}
		if nestedKey == "5: Time Period" {
			transformedModel.Metadata.TimePeriod = nestedVal.(string)
		}
		if nestedKey == "6: Series Type" {
			transformedModel.Metadata.SeriesType = nestedVal.(string)
		}
		if nestedKey == "7: Time Zone" {
			transformedModel.Metadata.TimeZone = nestedVal.(string)
		}

	}
}
func mapAdiResultMetadataAlphaVantageIntraExchangeRate(val interface{}, transformedModel *AlphaVantageTechnicalIndicatorResponse) {
	var s []*Result
	for nestedKey, nestedVal := range val.(map[string]interface{}) {
		var result = new(Result)
		result.Date = nestedKey
		for keyResult, resultValue := range nestedVal.(map[string]interface{}) {
			result.Indicator = keyResult
			result.Value = resultValue.(string)

		}
		s = append(s, result)
	}
	transformedModel.Result = s
}
