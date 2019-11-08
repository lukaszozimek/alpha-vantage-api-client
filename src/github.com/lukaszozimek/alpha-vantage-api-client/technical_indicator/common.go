package technical_indicator

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

type AlphaVantageTechnicalIndicatorResponse struct {
	Metadata AlphaVantageTechnicalIndicatorMetadata `json:"Metadata"`
	Result   map[string]interface{}
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

type Result struct {
	value string
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

func makeApiCallGet(url string, c *Client) *AlphaVantageTechnicalIndicatorResponse {

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
	var s = new(AlphaVantageTechnicalIndicatorResponse)
	err := json.Unmarshal(body, &s)
	if err != nil {
		fmt.Println("whoops:", err)
	}
	return s, err
}
