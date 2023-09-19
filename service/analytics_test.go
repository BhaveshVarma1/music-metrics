package service

import (
	"fmt"
	"testing"
)

func TestAnalytics(t *testing.T) {

	analytics := Analytics()
	fmt.Println(analytics)
}
