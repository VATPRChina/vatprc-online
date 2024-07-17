package service

import (
	"VatprcOnline/model"
	"encoding/json"
)

func ParseVatsimResponse(response []byte) (*model.VatsimResponse, error) {
	var result model.VatsimResponse
	err := json.Unmarshal(response, &result)
	return &result, err
}
