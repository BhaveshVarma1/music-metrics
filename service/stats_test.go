package service

import (
	"fmt"
	"testing"
)

func TestGetAverageYear(t *testing.T) {

	username := "prattnj"

	result := GetAverageYear(username)
	fmt.Print(result)
}

func TestGetSongCounts(t *testing.T) {

	username := "prattnj"

	result := GetSongCounts(username)
	fmt.Println(result)
}
