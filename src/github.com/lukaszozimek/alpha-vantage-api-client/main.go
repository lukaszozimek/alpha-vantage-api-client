package main

import (
	api "github.com/lukaszozimek/alpha-vantage-api-client/time_series"
	"net/http"
	"reflect"
)

func main() {
	client := &http.Client{}
	apiClient := &api.Client{nil, " ", client}
	var result = api.TimeSeriesDaily("MSFT", "29KSKPHV4ESREXEI", apiClient)
	reflect.ValueOf(result).MapKeys()

}
