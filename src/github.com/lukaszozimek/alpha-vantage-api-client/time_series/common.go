package time_series

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

type Client struct {
	BaseURL    *url.URL
	UserAgent  string
	HttpClient *http.Client
}
type AlphaVantageTimeSeriesApiResponse struct {
	Metadata Metadata `json:"Meta Data"`
	Result   map[time.Time]Result
}

type Metadata struct {
	Information   string `json:"1. Information"`
	Symbol        string `json:"2. Symbol"`
	LastRefreshed string `json:"3. Last Refreshed"`
	Interval      string `json:"4. Interval"`
	OutputSize    string `json:"5. Output Size"`
	TimeZone      string `json:"6. Time Zone"`
}

type Result struct {
	Open   float32 `json:"1. open"`
	High   float32 `json:"2. high"`
	Low    float32 `json:"3. low"`
	Close  float32 `json:"4. close"`
	Volume float32 `json:"5. volume"`
}

func makeApiCallGet(url string, c *Client) *AlphaVantageTimeSeriesApiResponse {

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
func getData(body []byte) (*AlphaVantageTimeSeriesApiResponse, error) {
	var s = new(AlphaVantageTimeSeriesApiResponse)
	err := json.Unmarshal(body, &s)
	if err != nil {
		fmt.Println("whoops:", err)
	}
	return s, err
}
