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

func TestFxRealTimeCall(t *testing.T) {
	var DefaultClient = &http.Client{}
	result, _ := url.Parse("https://www.alphavantage.co")
	var c = &crypto.Client{BaseURL: result, HttpClient: DefaultClient}
	c.HttpClient.Transport = roundTripFunc(func(r *http.Request) (*http.Response, error) {
		assert.Equal(t, "https://www.alphavantage.co/query?function=CURRENCY_EXCHANGE_RATE&from_currency=EUR&to_currency=PLN&apikey=deomo", r.URL.String())
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(strings.NewReader(`{}`)),
		}, nil
	})
	crypto.FxRealTimeCall("EUR", "PLN", "deomo", c)

}
