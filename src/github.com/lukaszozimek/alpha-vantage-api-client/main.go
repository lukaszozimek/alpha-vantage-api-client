package main

import (
	api "github.com/lukaszozimek/alpha-vantage-api-client/forex"
	"net/http"
)

func main() {
	client := &http.Client{}
	apiClient := &api.Client{nil, " ", client}
	var result = api.GetMonthly("EUR", "PLN", "29KSKPHV4ESREXEI", apiClient)
	print(result)

}
