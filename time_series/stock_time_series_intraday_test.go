package time_series_test

import (
	common "github.com/lukaszozimek/alpha-vantage-api-client"

	"github.com/lukaszozimek/alpha-vantage-api-client/time_series"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"testing"
)

func TestTimeSeriesIntraDayInterval1minute(t *testing.T) {
	var DefaultClient = &http.Client{}
	result, _ := url.Parse("https://www.alphavantage.co")
	var c = &common.Client{BaseURL: result, HttpClient: DefaultClient}
	c.HttpClient.Transport = roundTripFunc(func(r *http.Request) (*http.Response, error) {
		assert.Equal(t, "https://www.alphavantage.co/query?function=TIME_SERIES_INTRADAY&symbol=EUR&interval=1min&apikey=deomo", r.URL.String())
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(strings.NewReader(`{}`)),
		}, nil
	})
	item := time_series.TimeSeriesIntraDayInterval1minute("EUR", "deomo", c)
	assert.Equal(t, "{Information:\"\", Symbol:\"\", LastRefreshed:\"\", OutputSize:\"\", TimeZone:\"\", Interval:\"\"}", item.Metadata.String())
}

func TestTimeSeriesIntraDayInterval5minute(t *testing.T) {
	var DefaultClient = &http.Client{}
	result, _ := url.Parse("https://www.alphavantage.co")
	var c = &common.Client{BaseURL: result, HttpClient: DefaultClient}
	c.HttpClient.Transport = roundTripFunc(func(r *http.Request) (*http.Response, error) {
		assert.Equal(t, "https://www.alphavantage.co/query?function=TIME_SERIES_INTRADAY&symbol=EUR&interval=5min&apikey=deomo", r.URL.String())
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(strings.NewReader(`{}`)),
		}, nil
	})
	item := time_series.TimeSeriesIntraDayInterval5minute("EUR", "deomo", c)
	assert.Equal(t, "{Information:\"\", Symbol:\"\", LastRefreshed:\"\", OutputSize:\"\", TimeZone:\"\", Interval:\"\"}", item.Metadata.String())
}

func TestTimeSeriesIntraDayIntervalFifteenMinute(t *testing.T) {
	var DefaultClient = &http.Client{}
	result, _ := url.Parse("https://www.alphavantage.co")
	var c = &common.Client{BaseURL: result, HttpClient: DefaultClient}
	c.HttpClient.Transport = roundTripFunc(func(r *http.Request) (*http.Response, error) {
		assert.Equal(t, "https://www.alphavantage.co/query?function=TIME_SERIES_INTRADAY&symbol=EUR&interval=15min&apikey=deomo", r.URL.String())
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(strings.NewReader(`{}`)),
		}, nil
	})
	item := time_series.TimeSeriesIntraDayIntervalFifteenMinute("EUR", "deomo", c)
	assert.Equal(t, "{Information:\"\", Symbol:\"\", LastRefreshed:\"\", OutputSize:\"\", TimeZone:\"\", Interval:\"\"}", item.Metadata.String())
}
func TestTimeSeriesIntraDayIntervalThirtyMinute(t *testing.T) {
	var DefaultClient = &http.Client{}
	result, _ := url.Parse("https://www.alphavantage.co")
	var c = &common.Client{BaseURL: result, HttpClient: DefaultClient}
	c.HttpClient.Transport = roundTripFunc(func(r *http.Request) (*http.Response, error) {
		assert.Equal(t, "https://www.alphavantage.co/query?function=TIME_SERIES_INTRADAY&symbol=EUR&interval=30min&apikey=deomo", r.URL.String())
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(strings.NewReader(`{}`)),
		}, nil
	})
	item := time_series.TimeSeriesIntraDayIntervalThirtyMinute("EUR", "deomo", c)
	assert.Equal(t, "{Information:\"\", Symbol:\"\", LastRefreshed:\"\", OutputSize:\"\", TimeZone:\"\", Interval:\"\"}", item.Metadata.String())
}
func TestTimeSeriesIntraDayIntervalSixtyMinute(t *testing.T) {
	var DefaultClient = &http.Client{}
	result, _ := url.Parse("https://www.alphavantage.co")
	var c = &common.Client{BaseURL: result, HttpClient: DefaultClient}
	c.HttpClient.Transport = roundTripFunc(func(r *http.Request) (*http.Response, error) {
		assert.Equal(t, "https://www.alphavantage.co/query?function=TIME_SERIES_INTRADAY&symbol=EUR&interval=60min&apikey=deomo", r.URL.String())
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(strings.NewReader(`{}`)),
		}, nil
	})
	item := time_series.TimeSeriesIntraDayIntervalSixtyMinute("EUR", "deomo", c)
	assert.Equal(t, "{Information:\"\", Symbol:\"\", LastRefreshed:\"\", OutputSize:\"\", TimeZone:\"\", Interval:\"\"}", item.Metadata.String())
}
