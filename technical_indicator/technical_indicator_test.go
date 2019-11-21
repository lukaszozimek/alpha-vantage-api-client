package technical_indicator_test

import (
	common "github.com/lukaszozimek/alpha-vantage-api-client"
	"github.com/lukaszozimek/alpha-vantage-api-client/technical_indicator"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"testing"
)

func TestGetSMAReport(t *testing.T) {
	var DefaultClient = &http.Client{}
	result, _ := url.Parse("https://www.alphavantage.co")
	var c = &common.Client{BaseURL: result, HttpClient: DefaultClient}
	c.HttpClient.Transport = roundTripFunc(func(r *http.Request) (*http.Response, error) {
		assert.Equal(t, "https://www.alphavantage.coquery?function=SMA&symbol=EUR&interval=1min&time_period=60&series_type=open&apikey=deomo", r.URL.String())
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(strings.NewReader(`{}`)),
		}, nil
	})
	item := technical_indicator.GetSMAReport("EUR", technical_indicator.ONE_MINUTE, "60", technical_indicator.OPEN, "deomo", c)
	assert.Equal(t, "{Symbol:\"\", Indicator:\"\",  LastRefreshed:\"\",  Interval:\"\", TimePeriod:\"\", SeriesType:\"\", TimeZone:\"\"}", item.Metadata.String())

}
func TestGetEMAReport(t *testing.T) {
	var DefaultClient = &http.Client{}
	result, _ := url.Parse("https://www.alphavantage.co")
	var c = &common.Client{BaseURL: result, HttpClient: DefaultClient}
	c.HttpClient.Transport = roundTripFunc(func(r *http.Request) (*http.Response, error) {
		assert.Equal(t, "https://www.alphavantage.co/query?function=EMA&symbol=EUR&interval=1min&time_period=60&series_type=open&apikey=deomo", r.URL.String())
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(strings.NewReader(`{}`)),
		}, nil
	})
	item := technical_indicator.GetEMAReport("EUR", technical_indicator.ONE_MINUTE, "60", technical_indicator.OPEN, "deomo", c)
	assert.Equal(t, "{Symbol:\"\", Indicator:\"\",  LastRefreshed:\"\",  Interval:\"\", TimePeriod:\"\", SeriesType:\"\", TimeZone:\"\"}", item.Metadata.String())
}
func TestGetWMAReport(t *testing.T) {
	var DefaultClient = &http.Client{}
	result, _ := url.Parse("https://www.alphavantage.co")
	var c = &common.Client{BaseURL: result, HttpClient: DefaultClient}
	c.HttpClient.Transport = roundTripFunc(func(r *http.Request) (*http.Response, error) {
		assert.Equal(t, "https://www.alphavantage.co/query?function=WMA&symbol=EUR&interval=1min&time_period=60&series_type=open&apikey=deomo", r.URL.String())
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(strings.NewReader(`{}`)),
		}, nil
	})
	item := technical_indicator.GetWMAReport("EUR", technical_indicator.ONE_MINUTE, "60", technical_indicator.OPEN, "deomo", c)
	assert.Equal(t, "{Symbol:\"\", Indicator:\"\",  LastRefreshed:\"\",  Interval:\"\", TimePeriod:\"\", SeriesType:\"\", TimeZone:\"\"}", item.Metadata.String())
}
func TestGetDEMAReport(t *testing.T) {
	var DefaultClient = &http.Client{}
	result, _ := url.Parse("https://www.alphavantage.co")
	var c = &common.Client{BaseURL: result, HttpClient: DefaultClient}
	c.HttpClient.Transport = roundTripFunc(func(r *http.Request) (*http.Response, error) {
		assert.Equal(t, "https://www.alphavantage.co/query?function=DEMA&symbol=EUR&interval=1min&time_period=60&series_type=open&apikey=deomo", r.URL.String())
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(strings.NewReader(`{}`)),
		}, nil
	})
	item := technical_indicator.GetDEMAReport("EUR", technical_indicator.ONE_MINUTE, "60", technical_indicator.OPEN, "deomo", c)
	assert.Equal(t, "{Symbol:\"\", Indicator:\"\",  LastRefreshed:\"\",  Interval:\"\", TimePeriod:\"\", SeriesType:\"\", TimeZone:\"\"}", item.Metadata.String())
}
func TestGetTEMAReport(t *testing.T) {
	var DefaultClient = &http.Client{}
	result, _ := url.Parse("https://www.alphavantage.co")
	var c = &common.Client{BaseURL: result, HttpClient: DefaultClient}
	c.HttpClient.Transport = roundTripFunc(func(r *http.Request) (*http.Response, error) {
		assert.Equal(t, "https://www.alphavantage.co/query?function=TEMA&symbol=EUR&interval=1min&time_period=60&series_type=open&apikey=deomo", r.URL.String())
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(strings.NewReader(`{}`)),
		}, nil
	})
	item := technical_indicator.GetTEMAReport("EUR", technical_indicator.ONE_MINUTE, "60", technical_indicator.OPEN, "deomo", c)
	assert.Equal(t, "{Symbol:\"\", Indicator:\"\",  LastRefreshed:\"\",  Interval:\"\", TimePeriod:\"\", SeriesType:\"\", TimeZone:\"\"}", item.Metadata.String())
}
func TestGetTRIMAReport(t *testing.T) {
	var DefaultClient = &http.Client{}
	result, _ := url.Parse("https://www.alphavantage.co")
	var c = &common.Client{BaseURL: result, HttpClient: DefaultClient}
	c.HttpClient.Transport = roundTripFunc(func(r *http.Request) (*http.Response, error) {
		assert.Equal(t, "https://www.alphavantage.co/query?function=TRIMA&symbol=EUR&interval=1min&time_period=60&series_type=open&apikey=deomo", r.URL.String())
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(strings.NewReader(`{}`)),
		}, nil
	})
	item := technical_indicator.GetTRIMAReport("EUR", technical_indicator.ONE_MINUTE, "60", technical_indicator.OPEN, "deomo", c)
	assert.Equal(t, "{Symbol:\"\", Indicator:\"\",  LastRefreshed:\"\",  Interval:\"\", TimePeriod:\"\", SeriesType:\"\", TimeZone:\"\"}", item.Metadata.String())
}
func TestGetKAMAReport(t *testing.T) {
	var DefaultClient = &http.Client{}
	result, _ := url.Parse("https://www.alphavantage.co")
	var c = &common.Client{BaseURL: result, HttpClient: DefaultClient}
	c.HttpClient.Transport = roundTripFunc(func(r *http.Request) (*http.Response, error) {
		assert.Equal(t, "https://www.alphavantage.co/query?function=KAMA&symbol=EUR&interval=1min&time_period=60&series_type=open&apikey=deomo", r.URL.String())
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(strings.NewReader(`{}`)),
		}, nil
	})
	item := technical_indicator.GetKAMAReport("EUR", technical_indicator.ONE_MINUTE, "60", technical_indicator.OPEN, "deomo", c)
	assert.Equal(t, "{Symbol:\"\", Indicator:\"\",  LastRefreshed:\"\",  Interval:\"\", TimePeriod:\"\", SeriesType:\"\", TimeZone:\"\"}", item.Metadata.String())
}
func TestGetMAMAReport(t *testing.T) {
	var DefaultClient = &http.Client{}
	result, _ := url.Parse("https://www.alphavantage.co")
	var c = &common.Client{BaseURL: result, HttpClient: DefaultClient}
	c.HttpClient.Transport = roundTripFunc(func(r *http.Request) (*http.Response, error) {
		assert.Equal(t, "https://www.alphavantage.co/query?function=MAMA&symbol=EUR&interval=1min&time_period=60&series_type=open&apikey=deomo", r.URL.String())
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(strings.NewReader(`{}`)),
		}, nil
	})
	item := technical_indicator.GetMAMAReport("EUR", technical_indicator.ONE_MINUTE, "60", technical_indicator.OPEN, "deomo", c)
	assert.Equal(t, "{Symbol:\"\", Indicator:\"\",  LastRefreshed:\"\",  Interval:\"\", TimePeriod:\"\", SeriesType:\"\", TimeZone:\"\"}", item.Metadata.String())
}
func TestGetVWAPReport(t *testing.T) {
	var DefaultClient = &http.Client{}
	result, _ := url.Parse("https://www.alphavantage.co")
	var c = &common.Client{BaseURL: result, HttpClient: DefaultClient}
	c.HttpClient.Transport = roundTripFunc(func(r *http.Request) (*http.Response, error) {
		assert.Equal(t, "https://www.alphavantage.co/query?function=VWAP&symbol=EUR&interval=1min&time_period=60&series_type=open&apikey=deomo", r.URL.String())
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(strings.NewReader(`{}`)),
		}, nil
	})
	item := technical_indicator.GetVWAPReport("EUR", technical_indicator.ONE_MINUTE, "60", technical_indicator.OPEN, "deomo", c)
	assert.Equal(t, "{Symbol:\"\", Indicator:\"\",  LastRefreshed:\"\",  Interval:\"\", TimePeriod:\"\", SeriesType:\"\", TimeZone:\"\"}", item.Metadata.String())
}
func TestGetT3Report(t *testing.T) {
	var DefaultClient = &http.Client{}
	result, _ := url.Parse("https://www.alphavantage.co")
	var c = &common.Client{BaseURL: result, HttpClient: DefaultClient}
	c.HttpClient.Transport = roundTripFunc(func(r *http.Request) (*http.Response, error) {
		assert.Equal(t, "https://www.alphavantage.co/query?function=T3&symbol=EUR&interval=1min&time_period=60&series_type=open&apikey=deomo", r.URL.String())
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(strings.NewReader(`{}`)),
		}, nil
	})
	item := technical_indicator.GetT3Report("EUR", technical_indicator.ONE_MINUTE, "60", technical_indicator.OPEN, "deomo", c)
	assert.Equal(t, "{Symbol:\"\", Indicator:\"\",  LastRefreshed:\"\",  Interval:\"\", TimePeriod:\"\", SeriesType:\"\", TimeZone:\"\"}", item.Metadata.String())
}
func TestGetMACDReport(t *testing.T) {
	var DefaultClient = &http.Client{}
	result, _ := url.Parse("https://www.alphavantage.co")
	var c = &common.Client{BaseURL: result, HttpClient: DefaultClient}
	c.HttpClient.Transport = roundTripFunc(func(r *http.Request) (*http.Response, error) {
		assert.Equal(t, "https://www.alphavantage.co/query?function=MACD&symbol=EUR&interval=1min&time_period=60&series_type=open&apikey=deomo", r.URL.String())
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(strings.NewReader(`{}`)),
		}, nil
	})
	item := technical_indicator.GetMACDReport("EUR", technical_indicator.ONE_MINUTE, "60", technical_indicator.OPEN, "deomo", c)
	assert.Equal(t, "{Symbol:\"\", Indicator:\"\",  LastRefreshed:\"\",  Interval:\"\", TimePeriod:\"\", SeriesType:\"\", TimeZone:\"\"}", item.Metadata.String())
}
func TestGetMACDEXTReport(t *testing.T) {
	var DefaultClient = &http.Client{}
	result, _ := url.Parse("https://www.alphavantage.co")
	var c = &common.Client{BaseURL: result, HttpClient: DefaultClient}
	c.HttpClient.Transport = roundTripFunc(func(r *http.Request) (*http.Response, error) {
		assert.Equal(t, "https://www.alphavantage.co/query?function=MACDEXT&symbol=EUR&interval=1min&time_period=60&series_type=open&apikey=deomo", r.URL.String())
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(strings.NewReader(`{}`)),
		}, nil
	})
	item := technical_indicator.GetMACDEXTReport("EUR", technical_indicator.ONE_MINUTE, "60", technical_indicator.OPEN, "deomo", c)
	assert.Equal(t, "{Symbol:\"\", Indicator:\"\",  LastRefreshed:\"\",  Interval:\"\", TimePeriod:\"\", SeriesType:\"\", TimeZone:\"\"}", item.Metadata.String())
}
func TestGetSTOCHReport(t *testing.T) {
	var DefaultClient = &http.Client{}
	result, _ := url.Parse("https://www.alphavantage.co")
	var c = &common.Client{BaseURL: result, HttpClient: DefaultClient}
	c.HttpClient.Transport = roundTripFunc(func(r *http.Request) (*http.Response, error) {
		assert.Equal(t, "https://www.alphavantage.co/query?function=STOCH&symbol=EUR&interval=1min&time_period=60&series_type=open&apikey=deomo", r.URL.String())
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(strings.NewReader(`{}`)),
		}, nil
	})
	item := technical_indicator.GetSTOCHReport("EUR", technical_indicator.ONE_MINUTE, "60", technical_indicator.OPEN, "deomo", c)
	assert.Equal(t, "{Symbol:\"\", Indicator:\"\",  LastRefreshed:\"\",  Interval:\"\", TimePeriod:\"\", SeriesType:\"\", TimeZone:\"\"}", item.Metadata.String())
}
func TestGetSTOCHFReport(t *testing.T) {
	var DefaultClient = &http.Client{}
	result, _ := url.Parse("https://www.alphavantage.co")
	var c = &common.Client{BaseURL: result, HttpClient: DefaultClient}
	c.HttpClient.Transport = roundTripFunc(func(r *http.Request) (*http.Response, error) {
		assert.Equal(t, "https://www.alphavantage.co/query?function=STOCHF&symbol=EUR&interval=1min&time_period=60&series_type=open&apikey=deomo", r.URL.String())
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(strings.NewReader(`{}`)),
		}, nil
	})
	item := technical_indicator.GetSTOCHFReport("EUR", technical_indicator.ONE_MINUTE, "60", technical_indicator.OPEN, "deomo", c)
	assert.Equal(t, "{Symbol:\"\", Indicator:\"\",  LastRefreshed:\"\",  Interval:\"\", TimePeriod:\"\", SeriesType:\"\", TimeZone:\"\"}", item.Metadata.String())
}
func TestGetRSIReport(t *testing.T) {
	var DefaultClient = &http.Client{}
	result, _ := url.Parse("https://www.alphavantage.co")
	var c = &common.Client{BaseURL: result, HttpClient: DefaultClient}
	c.HttpClient.Transport = roundTripFunc(func(r *http.Request) (*http.Response, error) {
		assert.Equal(t, "https://www.alphavantage.co/query?function=RSI&symbol=EUR&interval=1min&time_period=60&series_type=open&apikey=deomo", r.URL.String())
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(strings.NewReader(`{}`)),
		}, nil
	})
	item := technical_indicator.GetRSIReport("EUR", technical_indicator.ONE_MINUTE, "60", technical_indicator.OPEN, "deomo", c)
	assert.Equal(t, "{Symbol:\"\", Indicator:\"\",  LastRefreshed:\"\",  Interval:\"\", TimePeriod:\"\", SeriesType:\"\", TimeZone:\"\"}", item.Metadata.String())
}
func TestGetSTOCHRSIReport(t *testing.T) {
	var DefaultClient = &http.Client{}
	result, _ := url.Parse("https://www.alphavantage.co")
	var c = &common.Client{BaseURL: result, HttpClient: DefaultClient}
	c.HttpClient.Transport = roundTripFunc(func(r *http.Request) (*http.Response, error) {
		assert.Equal(t, "https://www.alphavantage.co/query?function=STOCHRSI&symbol=EUR&interval=1min&time_period=60&series_type=open&apikey=deomo", r.URL.String())
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(strings.NewReader(`{}`)),
		}, nil
	})
	item := technical_indicator.GetSTOCHRSIReport("EUR", technical_indicator.ONE_MINUTE, "60", technical_indicator.OPEN, "deomo", c)
	assert.Equal(t, "{Symbol:\"\", Indicator:\"\",  LastRefreshed:\"\",  Interval:\"\", TimePeriod:\"\", SeriesType:\"\", TimeZone:\"\"}", item.Metadata.String())
}
func TestGetWILLRReport(t *testing.T) {
	var DefaultClient = &http.Client{}
	result, _ := url.Parse("https://www.alphavantage.co")
	var c = &common.Client{BaseURL: result, HttpClient: DefaultClient}
	c.HttpClient.Transport = roundTripFunc(func(r *http.Request) (*http.Response, error) {
		assert.Equal(t, "https://www.alphavantage.co/query?function=WILLR&symbol=EUR&interval=1min&time_period=60&series_type=open&apikey=deomo", r.URL.String())
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(strings.NewReader(`{}`)),
		}, nil
	})
	item := technical_indicator.GetWILLRReport("EUR", technical_indicator.ONE_MINUTE, "60", technical_indicator.OPEN, "deomo", c)
	assert.Equal(t, "{Symbol:\"\", Indicator:\"\",  LastRefreshed:\"\",  Interval:\"\", TimePeriod:\"\", SeriesType:\"\", TimeZone:\"\"}", item.Metadata.String())
}
func TestGetADXReport(t *testing.T) {
	var DefaultClient = &http.Client{}
	result, _ := url.Parse("https://www.alphavantage.co")
	var c = &common.Client{BaseURL: result, HttpClient: DefaultClient}
	c.HttpClient.Transport = roundTripFunc(func(r *http.Request) (*http.Response, error) {
		assert.Equal(t, "https://www.alphavantage.co/query?function=ADX&symbol=EUR&interval=1min&time_period=60&series_type=open&apikey=deomo", r.URL.String())
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(strings.NewReader(`{}`)),
		}, nil
	})
	item := technical_indicator.GetADXReport("EUR", technical_indicator.ONE_MINUTE, "60", technical_indicator.OPEN, "deomo", c)
	assert.Equal(t, "{Symbol:\"\", Indicator:\"\",  LastRefreshed:\"\",  Interval:\"\", TimePeriod:\"\", SeriesType:\"\", TimeZone:\"\"}", item.Metadata.String())
}
func TestGetADXRReport(t *testing.T) {
	var DefaultClient = &http.Client{}
	result, _ := url.Parse("https://www.alphavantage.co")
	var c = &common.Client{BaseURL: result, HttpClient: DefaultClient}
	c.HttpClient.Transport = roundTripFunc(func(r *http.Request) (*http.Response, error) {
		assert.Equal(t, "https://www.alphavantage.co/query?function=ADXR&symbol=EUR&interval=1min&time_period=60&series_type=open&apikey=deomo", r.URL.String())
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(strings.NewReader(`{}`)),
		}, nil
	})
	item := technical_indicator.GetADXRReport("EUR", technical_indicator.ONE_MINUTE, "60", technical_indicator.OPEN, "deomo", c)
	assert.Equal(t, "{Symbol:\"\", Indicator:\"\",  LastRefreshed:\"\",  Interval:\"\", TimePeriod:\"\", SeriesType:\"\", TimeZone:\"\"}", item.Metadata.String())
}
func TestGetAPOReport(t *testing.T) {
	var DefaultClient = &http.Client{}
	result, _ := url.Parse("https://www.alphavantage.co")
	var c = &common.Client{BaseURL: result, HttpClient: DefaultClient}
	c.HttpClient.Transport = roundTripFunc(func(r *http.Request) (*http.Response, error) {
		assert.Equal(t, "https://www.alphavantage.co/query?function=APO&symbol=EUR&interval=1min&time_period=60&series_type=open&apikey=deomo", r.URL.String())
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(strings.NewReader(`{}`)),
		}, nil
	})
	item := technical_indicator.GetAPOReport("EUR", technical_indicator.ONE_MINUTE, "60", technical_indicator.OPEN, "deomo", c)
	assert.Equal(t, "{Symbol:\"\", Indicator:\"\",  LastRefreshed:\"\",  Interval:\"\", TimePeriod:\"\", SeriesType:\"\", TimeZone:\"\"}", item.Metadata.String())
}
func TestGetPPOReport(t *testing.T) {
	var DefaultClient = &http.Client{}
	result, _ := url.Parse("https://www.alphavantage.co")
	var c = &common.Client{BaseURL: result, HttpClient: DefaultClient}
	c.HttpClient.Transport = roundTripFunc(func(r *http.Request) (*http.Response, error) {
		assert.Equal(t, "https://www.alphavantage.co/query?function=PPO&symbol=EUR&interval=1min&time_period=60&series_type=open&apikey=deomo", r.URL.String())
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(strings.NewReader(`{}`)),
		}, nil
	})
	item := technical_indicator.GetPPOReport("EUR", technical_indicator.ONE_MINUTE, "60", technical_indicator.OPEN, "deomo", c)
	assert.Equal(t, "{Symbol:\"\", Indicator:\"\",  LastRefreshed:\"\",  Interval:\"\", TimePeriod:\"\", SeriesType:\"\", TimeZone:\"\"}", item.Metadata.String())
}
func TestGetMOMReport(t *testing.T) {
	var DefaultClient = &http.Client{}
	result, _ := url.Parse("https://www.alphavantage.co")
	var c = &common.Client{BaseURL: result, HttpClient: DefaultClient}
	c.HttpClient.Transport = roundTripFunc(func(r *http.Request) (*http.Response, error) {
		assert.Equal(t, "https://www.alphavantage.co/query?function=MOM&symbol=EUR&interval=1min&time_period=60&series_type=open&apikey=deomo", r.URL.String())
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(strings.NewReader(`{}`)),
		}, nil
	})
	item := technical_indicator.GetMOMReport("EUR", technical_indicator.ONE_MINUTE, "60", technical_indicator.OPEN, "deomo", c)
	assert.Equal(t, "{Symbol:\"\", Indicator:\"\",  LastRefreshed:\"\",  Interval:\"\", TimePeriod:\"\", SeriesType:\"\", TimeZone:\"\"}", item.Metadata.String())
}
func TestGetBOPReport(t *testing.T) {
	var DefaultClient = &http.Client{}
	result, _ := url.Parse("https://www.alphavantage.co")
	var c = &common.Client{BaseURL: result, HttpClient: DefaultClient}
	c.HttpClient.Transport = roundTripFunc(func(r *http.Request) (*http.Response, error) {
		assert.Equal(t, "https://www.alphavantage.co/query?function=BOP&symbol=EUR&interval=1min&time_period=60&series_type=open&apikey=deomo", r.URL.String())
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(strings.NewReader(`{}`)),
		}, nil
	})
	item := technical_indicator.GetBOPReport("EUR", technical_indicator.ONE_MINUTE, "60", technical_indicator.OPEN, "deomo", c)
	assert.Equal(t, "{Symbol:\"\", Indicator:\"\",  LastRefreshed:\"\",  Interval:\"\", TimePeriod:\"\", SeriesType:\"\", TimeZone:\"\"}", item.Metadata.String())
}
func TestGetCCIReport(t *testing.T) {
	var DefaultClient = &http.Client{}
	result, _ := url.Parse("https://www.alphavantage.co")
	var c = &common.Client{BaseURL: result, HttpClient: DefaultClient}
	c.HttpClient.Transport = roundTripFunc(func(r *http.Request) (*http.Response, error) {
		assert.Equal(t, "https://www.alphavantage.co/query?function=CCI&symbol=EUR&interval=1min&time_period=60&series_type=open&apikey=deomo", r.URL.String())
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(strings.NewReader(`{}`)),
		}, nil
	})
	item := technical_indicator.GetCCIReport("EUR", technical_indicator.ONE_MINUTE, "60", technical_indicator.OPEN, "deomo", c)
	assert.Equal(t, "{Symbol:\"\", Indicator:\"\",  LastRefreshed:\"\",  Interval:\"\", TimePeriod:\"\", SeriesType:\"\", TimeZone:\"\"}", item.Metadata.String())
}
func TestGetCMOReport(t *testing.T) {
	var DefaultClient = &http.Client{}
	result, _ := url.Parse("https://www.alphavantage.co")
	var c = &common.Client{BaseURL: result, HttpClient: DefaultClient}
	c.HttpClient.Transport = roundTripFunc(func(r *http.Request) (*http.Response, error) {
		assert.Equal(t, "https://www.alphavantage.co/query?function=CMO&symbol=EUR&interval=1min&time_period=60&series_type=open&apikey=deomo", r.URL.String())
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(strings.NewReader(`{}`)),
		}, nil
	})
	item := technical_indicator.GetCMOReport("EUR", technical_indicator.ONE_MINUTE, "60", technical_indicator.OPEN, "deomo", c)
	assert.Equal(t, "{Symbol:\"\", Indicator:\"\",  LastRefreshed:\"\",  Interval:\"\", TimePeriod:\"\", SeriesType:\"\", TimeZone:\"\"}", item.Metadata.String())
}
func TestGetROCReport(t *testing.T) {
	var DefaultClient = &http.Client{}
	result, _ := url.Parse("https://www.alphavantage.co")
	var c = &common.Client{BaseURL: result, HttpClient: DefaultClient}
	c.HttpClient.Transport = roundTripFunc(func(r *http.Request) (*http.Response, error) {
		assert.Equal(t, "https://www.alphavantage.co/query?function=ROC&symbol=EUR&interval=1min&time_period=60&series_type=open&apikey=deomo", r.URL.String())
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(strings.NewReader(`{}`)),
		}, nil
	})
	item := technical_indicator.GetROCReport("EUR", technical_indicator.ONE_MINUTE, "60", technical_indicator.OPEN, "deomo", c)
	assert.Equal(t, "{Symbol:\"\", Indicator:\"\",  LastRefreshed:\"\",  Interval:\"\", TimePeriod:\"\", SeriesType:\"\", TimeZone:\"\"}", item.Metadata.String())
}
func TestGetROCRReport(t *testing.T) {
	var DefaultClient = &http.Client{}
	result, _ := url.Parse("https://www.alphavantage.co")
	var c = &common.Client{BaseURL: result, HttpClient: DefaultClient}
	c.HttpClient.Transport = roundTripFunc(func(r *http.Request) (*http.Response, error) {
		assert.Equal(t, "https://www.alphavantage.co/query?function=ROCR&symbol=EUR&interval=1min&time_period=60&series_type=open&apikey=deomo", r.URL.String())
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(strings.NewReader(`{}`)),
		}, nil
	})
	item := technical_indicator.GetROCRReport("EUR", technical_indicator.ONE_MINUTE, "60", technical_indicator.OPEN, "deomo", c)
	assert.Equal(t, "{Symbol:\"\", Indicator:\"\",  LastRefreshed:\"\",  Interval:\"\", TimePeriod:\"\", SeriesType:\"\", TimeZone:\"\"}", item.Metadata.String())
}
func TestGetAROONReport(t *testing.T) {
	var DefaultClient = &http.Client{}
	result, _ := url.Parse("https://www.alphavantage.co")
	var c = &common.Client{BaseURL: result, HttpClient: DefaultClient}
	c.HttpClient.Transport = roundTripFunc(func(r *http.Request) (*http.Response, error) {
		assert.Equal(t, "https://www.alphavantage.co/query?function=AROON&symbol=EUR&interval=1min&time_period=60&series_type=open&apikey=deomo", r.URL.String())
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(strings.NewReader(`{}`)),
		}, nil
	})
	item := technical_indicator.GetAROONReport("EUR", technical_indicator.ONE_MINUTE, "60", technical_indicator.OPEN, "deomo", c)
	assert.Equal(t, "{Symbol:\"\", Indicator:\"\",  LastRefreshed:\"\",  Interval:\"\", TimePeriod:\"\", SeriesType:\"\", TimeZone:\"\"}", item.Metadata.String())
}
func TestGetAROONOSCReport(t *testing.T) {
	var DefaultClient = &http.Client{}
	result, _ := url.Parse("https://www.alphavantage.co")
	var c = &common.Client{BaseURL: result, HttpClient: DefaultClient}
	c.HttpClient.Transport = roundTripFunc(func(r *http.Request) (*http.Response, error) {
		assert.Equal(t, "https://www.alphavantage.co/query?function=AROONOSC&symbol=EUR&interval=1min&time_period=60&series_type=open&apikey=deomo", r.URL.String())
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(strings.NewReader(`{}`)),
		}, nil
	})
	item := technical_indicator.GetAROONOSCReport("EUR", technical_indicator.ONE_MINUTE, "60", technical_indicator.OPEN, "deomo", c)
	assert.Equal(t, "{Symbol:\"\", Indicator:\"\",  LastRefreshed:\"\",  Interval:\"\", TimePeriod:\"\", SeriesType:\"\", TimeZone:\"\"}", item.Metadata.String())
}
func TestGetMFIReport(t *testing.T) {
	var DefaultClient = &http.Client{}
	result, _ := url.Parse("https://www.alphavantage.co")
	var c = &common.Client{BaseURL: result, HttpClient: DefaultClient}
	c.HttpClient.Transport = roundTripFunc(func(r *http.Request) (*http.Response, error) {
		assert.Equal(t, "https://www.alphavantage.co/query?function=MFI&symbol=EUR&interval=1min&time_period=60&series_type=open&apikey=deomo", r.URL.String())
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(strings.NewReader(`{}`)),
		}, nil
	})
	item := technical_indicator.GetMFIReport("EUR", technical_indicator.ONE_MINUTE, "60", technical_indicator.OPEN, "deomo", c)
	assert.Equal(t, "{Symbol:\"\", Indicator:\"\",  LastRefreshed:\"\",  Interval:\"\", TimePeriod:\"\", SeriesType:\"\", TimeZone:\"\"}", item.Metadata.String())
}
func TestGetTRIXReport(t *testing.T) {
	var DefaultClient = &http.Client{}
	result, _ := url.Parse("https://www.alphavantage.co")
	var c = &common.Client{BaseURL: result, HttpClient: DefaultClient}
	c.HttpClient.Transport = roundTripFunc(func(r *http.Request) (*http.Response, error) {
		assert.Equal(t, "https://www.alphavantage.co/query?function=TRIX&symbol=EUR&interval=1min&time_period=60&series_type=open&apikey=deomo", r.URL.String())
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(strings.NewReader(`{}`)),
		}, nil
	})
	item := technical_indicator.GetTRIXReport("EUR", technical_indicator.ONE_MINUTE, "60", technical_indicator.OPEN, "deomo", c)
	assert.Equal(t, "{Symbol:\"\", Indicator:\"\",  LastRefreshed:\"\",  Interval:\"\", TimePeriod:\"\", SeriesType:\"\", TimeZone:\"\"}", item.Metadata.String())
}
func TestGetULTOSCReport(t *testing.T) {
	var DefaultClient = &http.Client{}
	result, _ := url.Parse("https://www.alphavantage.co")
	var c = &common.Client{BaseURL: result, HttpClient: DefaultClient}
	c.HttpClient.Transport = roundTripFunc(func(r *http.Request) (*http.Response, error) {
		assert.Equal(t, "https://www.alphavantage.co/query?function=ULTOSC&symbol=EUR&interval=1min&time_period=60&series_type=open&apikey=deomo", r.URL.String())
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(strings.NewReader(`{}`)),
		}, nil
	})
	item := technical_indicator.GetULTOSCReport("EUR", technical_indicator.ONE_MINUTE, "60", technical_indicator.OPEN, "deomo", c)
	assert.Equal(t, "{Symbol:\"\", Indicator:\"\",  LastRefreshed:\"\",  Interval:\"\", TimePeriod:\"\", SeriesType:\"\", TimeZone:\"\"}", item.Metadata.String())
}
func TestGetDXReport(t *testing.T) {
	var DefaultClient = &http.Client{}
	result, _ := url.Parse("https://www.alphavantage.co")
	var c = &common.Client{BaseURL: result, HttpClient: DefaultClient}
	c.HttpClient.Transport = roundTripFunc(func(r *http.Request) (*http.Response, error) {
		assert.Equal(t, "https://www.alphavantage.co/query?function=DX&symbol=EUR&interval=1min&time_period=60&series_type=open&apikey=deomo", r.URL.String())
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(strings.NewReader(`{}`)),
		}, nil
	})
	item := technical_indicator.GetDXReport("EUR", technical_indicator.ONE_MINUTE, "60", technical_indicator.OPEN, "deomo", c)
	assert.Equal(t, "{Symbol:\"\", Indicator:\"\",  LastRefreshed:\"\",  Interval:\"\", TimePeriod:\"\", SeriesType:\"\", TimeZone:\"\"}", item.Metadata.String())
}
func TestGetminusDireport(t *testing.T) {
	var DefaultClient = &http.Client{}
	result, _ := url.Parse("https://www.alphavantage.co")
	var c = &common.Client{BaseURL: result, HttpClient: DefaultClient}
	c.HttpClient.Transport = roundTripFunc(func(r *http.Request) (*http.Response, error) {
		assert.Equal(t, "https://www.alphavantage.co/query?function=MINUS_DI&symbol=EUR&interval=1min&time_period=60&series_type=open&apikey=deomo", r.URL.String())
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(strings.NewReader(`{}`)),
		}, nil
	})
	item := technical_indicator.GetminusDireport("EUR", technical_indicator.ONE_MINUTE, "60", technical_indicator.OPEN, "deomo", c)
	assert.Equal(t, "{Symbol:\"\", Indicator:\"\",  LastRefreshed:\"\",  Interval:\"\", TimePeriod:\"\", SeriesType:\"\", TimeZone:\"\"}", item.Metadata.String())
}
func TestGetplusDireport(t *testing.T) {
	var DefaultClient = &http.Client{}
	result, _ := url.Parse("https://www.alphavantage.co")
	var c = &common.Client{BaseURL: result, HttpClient: DefaultClient}
	c.HttpClient.Transport = roundTripFunc(func(r *http.Request) (*http.Response, error) {
		assert.Equal(t, "https://www.alphavantage.co/query?function=PLUS_DI&symbol=EUR&interval=1min&time_period=60&series_type=open&apikey=deomo", r.URL.String())
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(strings.NewReader(`{}`)),
		}, nil
	})
	item := technical_indicator.GetplusDireport("EUR", technical_indicator.ONE_MINUTE, "60", technical_indicator.OPEN, "deomo", c)
	assert.Equal(t, "{Symbol:\"\", Indicator:\"\",  LastRefreshed:\"\",  Interval:\"\", TimePeriod:\"\", SeriesType:\"\", TimeZone:\"\"}", item.Metadata.String())
}
func TestGetminusDmreport(t *testing.T) {
	var DefaultClient = &http.Client{}
	result, _ := url.Parse("https://www.alphavantage.co")
	var c = &common.Client{BaseURL: result, HttpClient: DefaultClient}
	c.HttpClient.Transport = roundTripFunc(func(r *http.Request) (*http.Response, error) {
		assert.Equal(t, "https://www.alphavantage.co/query?function=MINUS_DM&symbol=EUR&interval=1min&time_period=60&series_type=open&apikey=deomo", r.URL.String())
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(strings.NewReader(`{}`)),
		}, nil
	})
	item := technical_indicator.GetminusDmreport("EUR", technical_indicator.ONE_MINUTE, "60", technical_indicator.OPEN, "deomo", c)
	assert.Equal(t, "{Symbol:\"\", Indicator:\"\",  LastRefreshed:\"\",  Interval:\"\", TimePeriod:\"\", SeriesType:\"\", TimeZone:\"\"}", item.Metadata.String())
}
func TestGetplusDmreport(t *testing.T) {
	var DefaultClient = &http.Client{}
	result, _ := url.Parse("https://www.alphavantage.co")
	var c = &common.Client{BaseURL: result, HttpClient: DefaultClient}
	c.HttpClient.Transport = roundTripFunc(func(r *http.Request) (*http.Response, error) {
		assert.Equal(t, "https://www.alphavantage.co/query?function=PLUS_DM&symbol=EUR&interval=1min&time_period=60&series_type=open&apikey=deomo", r.URL.String())
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(strings.NewReader(`{}`)),
		}, nil
	})
	item := technical_indicator.GetplusDmreport("EUR", technical_indicator.ONE_MINUTE, "60", technical_indicator.OPEN, "deomo", c)
	assert.Equal(t, "{Symbol:\"\", Indicator:\"\",  LastRefreshed:\"\",  Interval:\"\", TimePeriod:\"\", SeriesType:\"\", TimeZone:\"\"}", item.Metadata.String())
}
func TestGetBBANDSReport(t *testing.T) {
	var DefaultClient = &http.Client{}
	result, _ := url.Parse("https://www.alphavantage.co")
	var c = &common.Client{BaseURL: result, HttpClient: DefaultClient}
	c.HttpClient.Transport = roundTripFunc(func(r *http.Request) (*http.Response, error) {
		assert.Equal(t, "https://www.alphavantage.co/query?function=BBANDS&symbol=EUR&interval=1min&time_period=60&series_type=open&apikey=deomo", r.URL.String())
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(strings.NewReader(`{}`)),
		}, nil
	})
	item := technical_indicator.GetBBANDSReport("EUR", technical_indicator.ONE_MINUTE, "60", technical_indicator.OPEN, "deomo", c)
	assert.Equal(t, "{Symbol:\"\", Indicator:\"\",  LastRefreshed:\"\",  Interval:\"\", TimePeriod:\"\", SeriesType:\"\", TimeZone:\"\"}", item.Metadata.String())
}
func TestGetMIDPOINTReport(t *testing.T) {
	var DefaultClient = &http.Client{}
	result, _ := url.Parse("https://www.alphavantage.co")
	var c = &common.Client{BaseURL: result, HttpClient: DefaultClient}
	c.HttpClient.Transport = roundTripFunc(func(r *http.Request) (*http.Response, error) {
		assert.Equal(t, "https://www.alphavantage.co/query?function=MIDPOINT&symbol=EUR&interval=1min&time_period=60&series_type=open&apikey=deomo", r.URL.String())
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(strings.NewReader(`{}`)),
		}, nil
	})
	item := technical_indicator.GetMIDPOINTReport("EUR", technical_indicator.ONE_MINUTE, "60", technical_indicator.OPEN, "deomo", c)
	assert.Equal(t, "{Symbol:\"\", Indicator:\"\",  LastRefreshed:\"\",  Interval:\"\", TimePeriod:\"\", SeriesType:\"\", TimeZone:\"\"}", item.Metadata.String())
}
func TestGetMIDPRICEReport(t *testing.T) {
	var DefaultClient = &http.Client{}
	result, _ := url.Parse("https://www.alphavantage.co")
	var c = &common.Client{BaseURL: result, HttpClient: DefaultClient}
	c.HttpClient.Transport = roundTripFunc(func(r *http.Request) (*http.Response, error) {
		assert.Equal(t, "https://www.alphavantage.co/query?function=MIDPRICE&symbol=EUR&interval=1min&time_period=60&series_type=open&apikey=deomo", r.URL.String())
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(strings.NewReader(`{}`)),
		}, nil
	})
	item := technical_indicator.GetMIDPRICEReport("EUR", technical_indicator.ONE_MINUTE, "60", technical_indicator.OPEN, "deomo", c)
	assert.Equal(t, "{Symbol:\"\", Indicator:\"\",  LastRefreshed:\"\",  Interval:\"\", TimePeriod:\"\", SeriesType:\"\", TimeZone:\"\"}", item.Metadata.String())
}
func TestGetSARReport(t *testing.T) {
	var DefaultClient = &http.Client{}
	result, _ := url.Parse("https://www.alphavantage.co")
	var c = &common.Client{BaseURL: result, HttpClient: DefaultClient}
	c.HttpClient.Transport = roundTripFunc(func(r *http.Request) (*http.Response, error) {
		assert.Equal(t, "https://www.alphavantage.co/query?function=SAR&symbol=EUR&interval=1min&time_period=60&series_type=open&apikey=deomo", r.URL.String())
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(strings.NewReader(`{}`)),
		}, nil
	})
	item := technical_indicator.GetSARReport("EUR", technical_indicator.ONE_MINUTE, "60", technical_indicator.OPEN, "deomo", c)
	assert.Equal(t, "{Symbol:\"\", Indicator:\"\",  LastRefreshed:\"\",  Interval:\"\", TimePeriod:\"\", SeriesType:\"\", TimeZone:\"\"}", item.Metadata.String())
}
func TestGetTRANGEReport(t *testing.T) {
	var DefaultClient = &http.Client{}
	result, _ := url.Parse("https://www.alphavantage.co")
	var c = &common.Client{BaseURL: result, HttpClient: DefaultClient}
	c.HttpClient.Transport = roundTripFunc(func(r *http.Request) (*http.Response, error) {
		assert.Equal(t, "https://www.alphavantage.co/query?function=TRANGE&symbol=EUR&interval=1min&time_period=60&series_type=open&apikey=deomo", r.URL.String())
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(strings.NewReader(`{}`)),
		}, nil
	})
	item := technical_indicator.GetTRANGEReport("EUR", technical_indicator.ONE_MINUTE, "60", technical_indicator.OPEN, "deomo", c)
	assert.Equal(t, "{Symbol:\"\", Indicator:\"\",  LastRefreshed:\"\",  Interval:\"\", TimePeriod:\"\", SeriesType:\"\", TimeZone:\"\"}", item.Metadata.String())
}
func TestGetATRReport(t *testing.T) {
	var DefaultClient = &http.Client{}
	result, _ := url.Parse("https://www.alphavantage.co")
	var c = &common.Client{BaseURL: result, HttpClient: DefaultClient}
	c.HttpClient.Transport = roundTripFunc(func(r *http.Request) (*http.Response, error) {
		assert.Equal(t, "https://www.alphavantage.co/query?function=ATR&symbol=EUR&interval=1min&time_period=60&series_type=open&apikey=deomo", r.URL.String())
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(strings.NewReader(`{}`)),
		}, nil
	})
	item := technical_indicator.GetATRReport("EUR", technical_indicator.ONE_MINUTE, "60", technical_indicator.OPEN, "deomo", c)
	assert.Equal(t, "{Symbol:\"\", Indicator:\"\",  LastRefreshed:\"\",  Interval:\"\", TimePeriod:\"\", SeriesType:\"\", TimeZone:\"\"}", item.Metadata.String())
}
func TestGetNATRReport(t *testing.T) {
	var DefaultClient = &http.Client{}
	result, _ := url.Parse("https://www.alphavantage.co")
	var c = &common.Client{BaseURL: result, HttpClient: DefaultClient}
	c.HttpClient.Transport = roundTripFunc(func(r *http.Request) (*http.Response, error) {
		assert.Equal(t, "https://www.alphavantage.co/query?function=NATR&symbol=EUR&interval=1min&time_period=60&series_type=open&apikey=deomo", r.URL.String())
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(strings.NewReader(`{}`)),
		}, nil
	})
	item := technical_indicator.GetNATRReport("EUR", technical_indicator.ONE_MINUTE, "60", technical_indicator.OPEN, "deomo", c)
	assert.Equal(t, "{Symbol:\"\", Indicator:\"\",  LastRefreshed:\"\",  Interval:\"\", TimePeriod:\"\", SeriesType:\"\", TimeZone:\"\"}", item.Metadata.String())
}
func TestGetADReport(t *testing.T) {
	var DefaultClient = &http.Client{}
	result, _ := url.Parse("https://www.alphavantage.co")
	var c = &common.Client{BaseURL: result, HttpClient: DefaultClient}
	c.HttpClient.Transport = roundTripFunc(func(r *http.Request) (*http.Response, error) {
		assert.Equal(t, "https://www.alphavantage.co/query?function=AD&symbol=EUR&interval=1min&time_period=60&series_type=open&apikey=deomo", r.URL.String())
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(strings.NewReader(`{}`)),
		}, nil
	})
	item := technical_indicator.GetADReport("EUR", technical_indicator.ONE_MINUTE, "60", technical_indicator.OPEN, "deomo", c)
	assert.Equal(t, "{Symbol:\"\", Indicator:\"\",  LastRefreshed:\"\",  Interval:\"\", TimePeriod:\"\", SeriesType:\"\", TimeZone:\"\"}", item.Metadata.String())
}
func TestGetADOSCReport(t *testing.T) {
	var DefaultClient = &http.Client{}
	result, _ := url.Parse("https://www.alphavantage.co")
	var c = &common.Client{BaseURL: result, HttpClient: DefaultClient}
	c.HttpClient.Transport = roundTripFunc(func(r *http.Request) (*http.Response, error) {
		assert.Equal(t, "https://www.alphavantage.co/query?function=ADOSC&symbol=EUR&interval=1min&time_period=60&series_type=open&apikey=deomo", r.URL.String())
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(strings.NewReader(`{}`)),
		}, nil
	})
	item := technical_indicator.GetADOSCReport("EUR", technical_indicator.ONE_MINUTE, "60", technical_indicator.OPEN, "deomo", c)
	assert.Equal(t, "{Symbol:\"\", Indicator:\"\",  LastRefreshed:\"\",  Interval:\"\", TimePeriod:\"\", SeriesType:\"\", TimeZone:\"\"}", item.Metadata.String())
}
func TestGetOBVReport(t *testing.T) {
	var DefaultClient = &http.Client{}
	result, _ := url.Parse("https://www.alphavantage.co")
	var c = &common.Client{BaseURL: result, HttpClient: DefaultClient}
	c.HttpClient.Transport = roundTripFunc(func(r *http.Request) (*http.Response, error) {
		assert.Equal(t, "https://www.alphavantage.co/query?function=OBV&symbol=EUR&interval=1min&time_period=60&series_type=open&apikey=deomo", r.URL.String())
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(strings.NewReader(`{}`)),
		}, nil
	})
	item := technical_indicator.GetOBVReport("EUR", technical_indicator.ONE_MINUTE, "60", technical_indicator.OPEN, "deomo", c)
	assert.Equal(t, "{Symbol:\"\", Indicator:\"\",  LastRefreshed:\"\",  Interval:\"\", TimePeriod:\"\", SeriesType:\"\", TimeZone:\"\"}", item.Metadata.String())
}
func TestGethtTrendlinereport(t *testing.T) {
	var DefaultClient = &http.Client{}
	result, _ := url.Parse("https://www.alphavantage.co")
	var c = &common.Client{BaseURL: result, HttpClient: DefaultClient}
	c.HttpClient.Transport = roundTripFunc(func(r *http.Request) (*http.Response, error) {
		assert.Equal(t, "https://www.alphavantage.co/query?function=HT_TRENDLINE&symbol=EUR&interval=1min&time_period=60&series_type=open&apikey=deomo", r.URL.String())
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(strings.NewReader(`{}`)),
		}, nil
	})
	item := technical_indicator.GethtTrendlinereport("EUR", technical_indicator.ONE_MINUTE, "60", technical_indicator.OPEN, "deomo", c)
	assert.Equal(t, "{Symbol:\"\", Indicator:\"\",  LastRefreshed:\"\",  Interval:\"\", TimePeriod:\"\", SeriesType:\"\", TimeZone:\"\"}", item.Metadata.String())
}
func TestGethtSinereport(t *testing.T) {
	var DefaultClient = &http.Client{}
	result, _ := url.Parse("https://www.alphavantage.co")
	var c = &common.Client{BaseURL: result, HttpClient: DefaultClient}
	c.HttpClient.Transport = roundTripFunc(func(r *http.Request) (*http.Response, error) {
		assert.Equal(t, "https://www.alphavantage.co/query?function=HT_SINE&symbol=EUR&interval=1min&time_period=60&series_type=open&apikey=deomo", r.URL.String())
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(strings.NewReader(`{}`)),
		}, nil
	})
	item := technical_indicator.GethtSinereport("EUR", technical_indicator.ONE_MINUTE, "60", technical_indicator.OPEN, "deomo", c)
	assert.Equal(t, "{Symbol:\"\", Indicator:\"\",  LastRefreshed:\"\",  Interval:\"\", TimePeriod:\"\", SeriesType:\"\", TimeZone:\"\"}", item.Metadata.String())
}
func TestGethtTrendmodereport(t *testing.T) {
	var DefaultClient = &http.Client{}
	result, _ := url.Parse("https://www.alphavantage.co")
	var c = &common.Client{BaseURL: result, HttpClient: DefaultClient}
	c.HttpClient.Transport = roundTripFunc(func(r *http.Request) (*http.Response, error) {
		assert.Equal(t, "https://www.alphavantage.co/query?function=HT_TRENDMODE&symbol=EUR&interval=1min&time_period=60&series_type=open&apikey=deomo", r.URL.String())
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(strings.NewReader(`{}`)),
		}, nil
	})
	item := technical_indicator.GethtTrendmodereport("EUR", technical_indicator.ONE_MINUTE, "60", technical_indicator.OPEN, "deomo", c)
	assert.Equal(t, "{Symbol:\"\", Indicator:\"\",  LastRefreshed:\"\",  Interval:\"\", TimePeriod:\"\", SeriesType:\"\", TimeZone:\"\"}", item.Metadata.String())
}
func TestGethtDcperiodreport(t *testing.T) {
	var DefaultClient = &http.Client{}
	result, _ := url.Parse("https://www.alphavantage.co")
	var c = &common.Client{BaseURL: result, HttpClient: DefaultClient}
	c.HttpClient.Transport = roundTripFunc(func(r *http.Request) (*http.Response, error) {
		assert.Equal(t, "https://www.alphavantage.co/query?function=HT_DCPERIOD&symbol=EUR&interval=1min&time_period=60&series_type=open&apikey=deomo", r.URL.String())
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(strings.NewReader(`{}`)),
		}, nil
	})
	item := technical_indicator.GethtDcperiodreport("EUR", technical_indicator.ONE_MINUTE, "60", technical_indicator.OPEN, "deomo", c)
	assert.Equal(t, "{Symbol:\"\", Indicator:\"\",  LastRefreshed:\"\",  Interval:\"\", TimePeriod:\"\", SeriesType:\"\", TimeZone:\"\"}", item.Metadata.String())
}
func TestGethtDcphasereport(t *testing.T) {
	var DefaultClient = &http.Client{}
	result, _ := url.Parse("https://www.alphavantage.co")
	var c = &common.Client{BaseURL: result, HttpClient: DefaultClient}
	c.HttpClient.Transport = roundTripFunc(func(r *http.Request) (*http.Response, error) {
		assert.Equal(t, "https://www.alphavantage.co/query?function=HT_DCPHASE&symbol=EUR&interval=1min&time_period=60&series_type=open&apikey=deomo", r.URL.String())
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(strings.NewReader(`{}`)),
		}, nil
	})
	item := technical_indicator.GethtDcphasereport("EUR", technical_indicator.ONE_MINUTE, "60", technical_indicator.OPEN, "deomo", c)
	assert.Equal(t, "{Symbol:\"\", Indicator:\"\",  LastRefreshed:\"\",  Interval:\"\", TimePeriod:\"\", SeriesType:\"\", TimeZone:\"\"}", item.Metadata.String())
}
func TestGethtPhasorreport(t *testing.T) {
	var DefaultClient = &http.Client{}
	result, _ := url.Parse("https://www.alphavantage.co")
	var c = &common.Client{BaseURL: result, HttpClient: DefaultClient}
	c.HttpClient.Transport = roundTripFunc(func(r *http.Request) (*http.Response, error) {
		assert.Equal(t, "https://www.alphavantage.co/query?function=HT_PHASOR&symbol=EUR&interval=1min&time_period=60&series_type=open&apikey=deomo", r.URL.String())
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(strings.NewReader(`{}`)),
		}, nil
	})
	item := technical_indicator.GethtPhasorreport("EUR", technical_indicator.ONE_MINUTE, "60", technical_indicator.OPEN, "deomo", c)
	assert.Equal(t, "{Symbol:\"\", Indicator:\"\",  LastRefreshed:\"\",  Interval:\"\", TimePeriod:\"\", SeriesType:\"\", TimeZone:\"\"}", item.Metadata.String())
}
