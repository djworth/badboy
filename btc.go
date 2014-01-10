package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func fetch() map[string]map[string]string {
	var data map[string]map[string]string

	resp, _ := http.Get("http://api.bitcoincharts.com/v1/weighted_prices.json")
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	json.Unmarshal(body, &data)

	return data

}

func CurrentBTCPrice(currency string) string {
	prices := fetch()

	return prices[currency]["24h"]

}
