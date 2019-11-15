package time_series_test

import (
	"github.com/lukaszozimek/alpha-vantage-api-client/time_series"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"testing"
)

func TestGetMonthly(t *testing.T) {
	var DefaultClient = &http.Client{}
	result, _ := url.Parse("https://www.alphavantage.co")
	var c = &time_series.Client{BaseURL: result, HttpClient: DefaultClient}
	c.HttpClient.Transport = roundTripFunc(func(r *http.Request) (*http.Response, error) {
		assert.Equal(t, "https://www.alphavantage.co/query?function=TIME_SERIES_MONTHLY&symbol=EUR&apikey=deomo", r.URL.String())
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(strings.NewReader(`{}`)),
		}, nil
	})
	item := time_series.GetMonthly("EUR", "deomo", c)
	assert.Equal(t, "{Information:\"\", Symbol:\"\", LastRefreshed:\"\", OutputSize:\"\", TimeZone:\"\", Interval:\"\"}", item.Metadata.String())
}
