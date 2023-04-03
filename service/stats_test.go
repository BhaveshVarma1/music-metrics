package service

import (
	"fmt"
	"testing"
)

func TestGetAverageYear(t *testing.T) {

	username := "prattnj"
	var service GetAverageYearService

	result := service.ExecuteService(username)
	fmt.Print(result)
}

func TestGetSongCounts(t *testing.T) {

	username := "prattnj"
	var service GetSongCountsService

	result := service.ExecuteService(username)
	fmt.Println(result)
}
