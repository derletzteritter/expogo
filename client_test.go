package expogo

import (
	"fmt"
	"testing"
)

func TestNewClient(t *testing.T) {
	client := NewExpoClient(nil)

	responses, err := client.SendPushNotification(&Notification{
		To:       []string{""},
		Title:    "Hello World",
		Body:     "This is a test notification",
		Priority: HighPriority,
	})
	if err != nil {
		fmt.Println(err.Error())
	}

	for _, response := range responses {
		if response.Status == ReceiptResponseError {
			fmt.Println(response.Details.Error)
		}
	}
}
