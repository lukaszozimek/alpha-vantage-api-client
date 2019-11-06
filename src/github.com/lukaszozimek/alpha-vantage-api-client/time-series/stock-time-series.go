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

const (
	ONE_MINUTE        = "1min"
	FIVE_MINUTE       = "5min"
	FIFITHTEEN_MINUTE = "15min"
	THIRTY_MINUTE     = "30min"
	SIXTY_MINUTE      = "60min"
)

type AlphaVantageTimeSeriesApiResponse struct {
	metadata Metadata `json:"Meta Data"`
	result   map[time.Time]Result
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
	open   float32 `json:"1. open"`
	high   float32 `json:"2. high"`
	low    float32 `json:"3. low"`
	close  float32 `json:"4. close"`
	volume float32 `json:"5. volume"`
}

func TimeSeriesIntraDayInterval1minute(symbol string, c *Client) *AlphaVantageTimeSeriesApiResponse {

	return timeSeriesIntraDay(symbol, ONE_MINUTE, c)
}
func TimeSeriesIntraDayInterval5minute(symbol string, c *Client) *AlphaVantageTimeSeriesApiResponse {

	return timeSeriesIntraDay(symbol, FIVE_MINUTE, c)
}
func TimeSeriesIntraDayIntervalFifithteenMinute(symbol string, c *Client) *AlphaVantageTimeSeriesApiResponse {

	return timeSeriesIntraDay(symbol, FIFITHTEEN_MINUTE, c)
}
func TimeSeriesIntraDayIntervalThirtyMinute(symbol string, c *Client) *AlphaVantageTimeSeriesApiResponse {

	return timeSeriesIntraDay(symbol, THIRTY_MINUTE, c)
}
func TimeSeriesIntraDayIntervalSixtyMinute(symbol string, c *Client) *AlphaVantageTimeSeriesApiResponse {

	return timeSeriesIntraDay(symbol, SIXTY_MINUTE, c)
}
func timeSeriesIntraDay(symbol string, interval string, c *Client) *AlphaVantageTimeSeriesApiResponse {

	res, e := c.HttpClient.Get("https://www.alphavantage.co/query?function=TIME_SERIES_INTRADAY&symbol=MSFT&interval=5min&apikey=demo")
	if e != nil {
		panic(e.Error())
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}
	s, err := getStations(body)
	if err != nil {
		panic(err.Error())
	}
	return s
}
func getStations(body []byte) (*AlphaVantageTimeSeriesApiResponse, error) {
	var s = new(AlphaVantageTimeSeriesApiResponse)
	err := json.Unmarshal(body, &s)
	if err != nil {
		fmt.Println("whoops:", err)
	}
	return s, err
}
