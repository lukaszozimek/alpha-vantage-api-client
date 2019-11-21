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

func TestFxRealTimeCall(t *testing.T) {
	var DefaultClient = &http.Client{}
	result, _ := url.Parse("https://www.alphavantage.co")
	var c = &common.Client{BaseURL: result, HttpClient: DefaultClient}
	c.HttpClient.Transport = roundTripFunc(func(r *http.Request) (*http.Response, error) {
		assert.Equal(t, "https://www.alphavantage.co/query?function=TIME_SERIES_DAILY&symbol=EUR&apikey=deomo", r.URL.String())
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(strings.NewReader(`{}`)),
		}, nil
	})
	item := forex.FxRealTimeCall("EUR", "deomo", c)
	assert.Equal(t, "{Information:\"\", Symbol:\"\",  LastRefreshed:\"\",  OutputSize:\"\", TimeZone:\"\"}", item.Metadata.String())
}
