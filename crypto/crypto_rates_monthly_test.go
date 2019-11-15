package crypto_test

import (
	"github.com/lukaszozimek/alpha-vantage-api-client/crypto"
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
	var c = &crypto.Client{BaseURL: result, HttpClient: DefaultClient}
	c.HttpClient.Transport = roundTripFunc(func(r *http.Request) (*http.Response, error) {
		assert.Equal(t, "https://www.alphavantage.co/query?function=DIGITAL_CURRENCY_MONTHLY&symbol=EUR&market=PLN&apikey=deomo", r.URL.String())
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(strings.NewReader(`{}`)),
		}, nil
	})
	item := crypto.GetMonthly("EUR", "PLN", "deomo", c)
	assert.Equal(t, "{Information:\"\", DigitalCurrencyCode:\"\", DigitalCurrencyName:\"\", MarketCode:\"\", MarketName:\"\", LastRefreshed:\"\", TimeZone:\"\"}", item.Metadata.String())
}
