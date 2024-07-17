package service

import (
	"io/ioutil"
	"net/http"
	"time"
)

func FetchOnlineDataFromVatsim() ([]byte, error) {
	url := "https://data.vatsim.net/v3/vatsim-data.json"
	client := http.Client{
		Timeout: 1 * time.Second,
	}
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(resp.Body)
}

func FetchFutureAtcFromVatprc() ([]byte, error) {
	url := "https://atcapi.vatprc.net/v1/public/schedules"
	client := http.Client{
		Timeout: 1 * time.Second,
	}
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(resp.Body)
}
