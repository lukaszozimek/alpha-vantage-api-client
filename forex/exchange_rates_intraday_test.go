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

func TestTimeSeriesIntraDayIntervalFifteenMinute(t *testing.T) {
	var DefaultClient = &http.Client{}
	result, _ := url.Parse("https://www.alphavantage.co")
	var c = &common.Client{BaseURL: result, HttpClient: DefaultClient}
	c.HttpClient.Transport = roundTripFunc(func(r *http.Request) (*http.Response, error) {
		assert.Equal(t, "https://www.alphavantage.co/query?function=FX_INTRADAY&from_symbol=EUR&to_symbol=PLN&interval=15min&apikey=deomo", r.URL.String())
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(strings.NewReader(`{}`)),
		}, nil
	})
	item := forex.TimeSeriesIntraDayIntervalFifteenMinute("EUR", "PLN", "deomo", c)
	assert.Equal(t, "{Information:\"\", FromSymbol:\"\", ToSymbol:\"\", LastRefreshed:\"\", Interval:\"\", OutputSize:\"\", TimeZone:\"\"}", item.Metadata.String())
}

func TestTimeSeriesIntraDayIntervalSixtyMinute(t *testing.T) {
	var DefaultClient = &http.Client{}
	result, _ := url.Parse("https://www.alphavantage.co")
	var c = &common.Client{BaseURL: result, HttpClient: DefaultClient}
	c.HttpClient.Transport = roundTripFunc(func(r *http.Request) (*http.Response, error) {
		assert.Equal(t, "https://www.alphavantage.co/query?function=FX_INTRADAY&from_symbol=EUR&to_symbol=PLN&interval=60min&apikey=deomo", r.URL.String())
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(strings.NewReader(`{}`)),
		}, nil
	})
	item := forex.TimeSeriesIntraDayIntervalSixtyMinute("EUR", "PLN", "deomo", c)
	assert.Equal(t, "{Information:\"\", FromSymbol:\"\", ToSymbol:\"\", LastRefreshed:\"\", Interval:\"\", OutputSize:\"\", TimeZone:\"\"}", item.Metadata.String())
}
func TestTimeSeriesIntraDayIntervalThirtyMinute(t *testing.T) {
	var DefaultClient = &http.Client{}
	result, _ := url.Parse("https://www.alphavantage.co")
	var c = &common.Client{BaseURL: result, HttpClient: DefaultClient}
	c.HttpClient.Transport = roundTripFunc(func(r *http.Request) (*http.Response, error) {
		assert.Equal(t, "https://www.alphavantage.co/query?function=FX_INTRADAY&from_symbol=EUR&to_symbol=PLN&interval=30min&apikey=deomo", r.URL.String())
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(strings.NewReader(`{}`)),
		}, nil
	})
	item := forex.TimeSeriesIntraDayIntervalThirtyMinute("EUR", "PLN", "deomo", c)
	assert.Equal(t, "{Information:\"\", FromSymbol:\"\", ToSymbol:\"\", LastRefreshed:\"\", Interval:\"\", OutputSize:\"\", TimeZone:\"\"}", item.Metadata.String())
}

func TestTimeSeriesIntraDayInterval5minute(t *testing.T) {
	var DefaultClient = &http.Client{}
	result, _ := url.Parse("https://www.alphavantage.co")
	var c = &common.Client{BaseURL: result, HttpClient: DefaultClient}
	c.HttpClient.Transport = roundTripFunc(func(r *http.Request) (*http.Response, error) {
		assert.Equal(t, "https://www.alphavantage.co/query?function=FX_INTRADAY&from_symbol=EUR&to_symbol=PLN&interval=5min&apikey=deomo", r.URL.String())
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(strings.NewReader(`{}`)),
		}, nil
	})
	item := forex.TimeSeriesIntraDayInterval5minute("EUR", "PLN", "deomo", c)
	assert.Equal(t, "{Information:\"\", FromSymbol:\"\", ToSymbol:\"\", LastRefreshed:\"\", Interval:\"\", OutputSize:\"\", TimeZone:\"\"}", item.Metadata.String())
}

func TestTimeSeriesIntraDayInterval1minute(t *testing.T) {
	var DefaultClient = &http.Client{}
	result, _ := url.Parse("https://www.alphavantage.co")
	var c = &common.Client{BaseURL: result, HttpClient: DefaultClient}
	c.HttpClient.Transport = roundTripFunc(func(r *http.Request) (*http.Response, error) {
		assert.Equal(t, "https://www.alphavantage.co/query?function=FX_INTRADAY&from_symbol=EUR&to_symbol=PLN&interval=1min&apikey=deomo", r.URL.String())
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(strings.NewReader(`{}`)),
		}, nil
	})
	item := forex.TimeSeriesIntraDayInterval1minute("EUR", "PLN", "deomo", c)
	assert.Equal(t, "{Information:\"\", FromSymbol:\"\", ToSymbol:\"\", LastRefreshed:\"\", Interval:\"\", OutputSize:\"\", TimeZone:\"\"}", item.Metadata.String())
}
