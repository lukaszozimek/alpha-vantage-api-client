package main

import (
	api "github.com/lukaszozimek/alpha-vantage-api-client/time_series"
	"net/http"
)

func main() {
	client := &http.Client{}
	apiClient := &api.Client{nil, " ", client}
	var result = api.TimeSeriesIntraDayIntervalFifteenMinute("MSFT", "29KSKPHV4ESREXEI", apiClient)
	print(result)

}
