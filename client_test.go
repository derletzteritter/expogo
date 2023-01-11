package expogo

import "testing"

func TestNewClient(t *testing.T) {
	client := NewExpoClient(nil)

	client.SendPushNotification(&Notification{
		To:       "",
		Title:    "Hello World",
		Body:     "This is a test notification",
		Priority: HIGH_PRIORITY,
	})
}
