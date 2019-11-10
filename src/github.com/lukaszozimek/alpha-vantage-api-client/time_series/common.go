package time_series

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
type AlphaVantageTimeSeriesApiResponse struct {
	Metadata             Metadata             `json:"Meta Data"`
	ResultWrapperWrapper ResultWrapperWrapper `json:"result,omitempty"`
}
type ResultWrapperWrapper struct {
	Sprocket map[string]Result
	Partial  bool `json:"partial,omitempty"`
}
type Metadata struct {
	Information   string `json:"1. Information"`
	Symbol        string `json:"2. Symbol"`
	LastRefreshed string `json:"3. Last Refreshed"`
	OutputSize    string `json:"5. Output Size"`
	TimeZone      string `json:"6. Time Zone"`
	Interval      string `json:"4. Interval"`
}

type Result struct {
	Open   float32 `json:"1. open"`
	High   float32 `json:"2. high"`
	Low    float32 `json:"3. low"`
	Close  float32 `json:"4. close"`
	Volume float32 `json:"5. volume"`
}

func makeApiCallGet(url string, c *Client) (map[string]map[string]string, map[string]map[string]map[string]string) {

	res, e := c.HttpClient.Get(url)
	if e != nil {
		panic(e.Error())
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}
	metadata, result, err := getData(body)
	if err != nil {
		panic(err.Error())
	}
	return metadata, result
}
func getData(body []byte) (map[string]map[string]string, map[string]map[string]map[string]string, error) {
	//this is Messy API...
	var metadata map[string]map[string]string
	var result map[string]map[string]map[string]string
	err := json.Unmarshal(body, &metadata)
	err1 := json.Unmarshal(body, &result)
	if err != nil && err1 != nil {
		fmt.Println("whoops:", err)
	}

	return metadata, result, err
}
func mapMetadata(metadata map[string]map[string]string, response *AlphaVantageTimeSeriesApiResponse) {

}
func mapResultApi(result map[string]map[string]map[string]string, response *AlphaVantageTimeSeriesApiResponse) {

}
