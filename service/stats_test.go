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
