package main

import (
	"encoding/json"
	"fmt"
	"music-metrics/model"
	"music-metrics/service"
	"music-metrics/util"
	"os"
	"path/filepath"
)

func visit(path string, info os.FileInfo, err error) error {
	if err != nil {
		fmt.Println(err)
		return nil
	}

	if path[len(path)-5:] != ".json" {
		return nil
	}

	username := util.GetUsername(path)
	if username == "" {
		return nil
	}

	fileReader, err := os.Open(path)
	if err != nil {
		fmt.Printf("Error opening file: %v", err)
	}

	var req []model.ExtendedStreamingObject
	err = json.NewDecoder(fileReader).Decode(&req)
	if err != nil {
		fmt.Printf("Error decoding json: %v", err)
	}

	go service.Load(req, username)

	return nil
}

func main() {

	dir := os.Args[1]

	err := filepath.Walk(dir, visit)
	if err != nil {
		fmt.Print(err)
	}

}
