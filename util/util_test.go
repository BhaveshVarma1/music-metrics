package util

import (
	"fmt"
	"testing"
)

func Test_Email(t *testing.T) {
	emailBody := "Username: " + "silkypony7" + "\nDisplay name: " + "Lucy Pratt" + "\nTotal users: " + fmt.Sprintf("%d", 15)
	SendEmail("New Music Metrics User", emailBody)
}
