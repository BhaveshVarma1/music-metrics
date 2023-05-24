package service

import (
	"fmt"
	"music-metrics/model"
	"time"
)

func Load(history model.ExtendedStreamingHistory, username string) {
	// todo
	time.Sleep(30 * time.Second)
	fmt.Println("Received history for user " + username + " with " + fmt.Sprint(len(history.StreamingHistory)) + " entries")
}
