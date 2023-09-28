package service

import (
	"encoding/json"
	"music-metrics/model"
	"os"
	"testing"
)

func TestLoad(t *testing.T) {
	filePathStr := "../json-test/2021_0.json"
	fileReader, err := os.Open(filePathStr)
	if err != nil {
		t.Errorf("Error opening file: %v", err)
	}

	var req []model.ExtendedStreamingObject
	err = json.NewDecoder(fileReader).Decode(&req)
	if err != nil {
		t.Errorf("Error decoding json: %v", err)
	}

	Load(req, "1251455712")
}
