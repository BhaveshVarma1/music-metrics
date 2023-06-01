package service

import (
	"encoding/json"
	"music-metrics/model"
	"os"
	"testing"
)

func TestLoad(t *testing.T) {
	filePathStr := "../json-test/endsong_1.json"
	fileReader, err := os.Open(filePathStr)
	if err != nil {
		t.Errorf("Error opening file: %v", err)
	}

	var req []model.ExtendedStreamingObject
	err = json.NewDecoder(fileReader).Decode(&req)
	if err != nil {
		t.Errorf("Error decoding json: %v", err)
	}

	Load(req, GetUsername())
}
