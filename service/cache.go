package service

import (
	"fmt"
	"time"
)

var cache map[string][]byte

func init() {
	cache = make(map[string][]byte)
}

func getExpirationKey(key string) string {
	return fmt.Sprintf("%s%s", key, "Expiration")
}

func GetDataFromCache(key string) ([]byte, error) {
	expiration, found := cache[getExpirationKey(key)]
	if !found {
		return nil, nil
	}
	if time.Now().Format("2006-01-02 15:04:05") > string(expiration) {
		// Expired
		return nil, nil
	}

	data, found := cache[key]
	if !found {
		return nil, nil
	}
	return data, nil
}

func PutDataToCache(key string, data []byte, expirationSeconds int) error {
	cache[key] = data
	expData := time.Now().Add(time.Duration(expirationSeconds) * time.Second).Format("2006-01-02 15:04:05")
	cache[getExpirationKey(key)] = []byte(expData)
	return nil
}
