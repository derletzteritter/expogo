package expogo

import (
	"fmt"
	"testing"
)

func TestNewClient(t *testing.T) {
	client := NewExpoClient(nil)

	_, err := client.SendPushNotification(&Notification{
		To:       []string{""},
		Title:    "Hello World",
		Body:     "This is a test notification",
		Priority: HighPriority,
	})

	fmt.Println(err.Error())
}
