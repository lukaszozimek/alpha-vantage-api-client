package forex_test

import (
	common "github.com/lukaszozimek/alpha-vantage-api-client"
	"github.com/lukaszozimek/alpha-vantage-api-client/forex"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"testing"
)

func TestTimeSeriesDaily(t *testing.T) {
	var DefaultClient = &http.Client{}
	result, _ := url.Parse("https://www.alphavantage.co")
	var c = &common.Client{BaseURL: result, HttpClient: DefaultClient}
	c.HttpClient.Transport = roundTripFunc(func(r *http.Request) (*http.Response, error) {
		assert.Equal(t, "https://www.alphavantage.co/query?function=FX_DAILY&from_symbol=EUR&to_symbol=PLN&apikey=deomo", r.URL.String())
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(strings.NewReader(`{}`)),
		}, nil
	})
	item := forex.TimeSeriesDaily("EUR", "PLN", "deomo", c)
	assert.Equal(t, "{Information:\"\", FromSymbol:\"\", ToSymbol:\"\", LastRefreshed:\"\", Interval:\"\", OutputSize:\"\", TimeZone:\"\"}", item.Metadata.String())
}
