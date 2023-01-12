package expogo

import "testing"

func TestNewClient(t *testing.T) {
	client := NewExpoClient(nil)

	client.SendPushNotification(&Notification{
		To:       []string{"ExponentPushToken[jgjpK8NQcb4475rh2nn4K9]"},
		Title:    "Hello World",
		Body:     "This is a test notification",
		Priority: HIGH_PRIORITY,
	})
}
