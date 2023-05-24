package service

import (
	"encoding/json"
	"fmt"
	"music-metrics/model"
	"os"
	"testing"
)

func TestLoad(t *testing.T) {
	filePathStr := "../json-test/endsong_0.json"
	fileReader, err := os.Open(filePathStr)
	if err != nil {
		t.Errorf("Error opening file: %v", err)
	}

	var req []model.ExtendedStreamingObject
	err = json.NewDecoder(fileReader).Decode(&req)
	if err != nil {
		t.Errorf("Error decoding json: %v", err)
	}

	fmt.Println("made it this far")
}
