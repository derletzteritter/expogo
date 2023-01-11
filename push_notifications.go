package expogo

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	DEFAULT_PRIORITY = "default"
	NORMAL_PRIORITY  = "normal"
	HIGH_PRIORITY    = "high"
)

type Notification struct {
	To         string `json:"to"`
	Title      string `json:"title,omitempty"`
	Body       string `json:"body,omitempty"`
	TTL        int    `json:"ttl,omitempty"`
	Expiration int    `json:"expiration,omitempty"`
	Priority   string `json:"priority,omitempty"`
}

func (client *ExpoClient) SendPushNotification(notification *Notification) {
	log.Println("Sending push notification")
	log.Println(notification.To)

	body, err := json.Marshal(notification)
	if err != nil {
		log.Println(err.Error())
	}

	log.Println(string(body))

	req, err := http.NewRequest("POST", client.host+client.pushPath, bytes.NewReader(body))
	if err != nil {
		log.Println(err.Error())
	}

	req.Header.Add("Content-Type", "application/json")

	resp, err := client.httpClient.Do(req)
	if err != nil {
		log.Println(err.Error())
	}

	defer resp.Body.Close()

	log.Println(resp.StatusCode)
	respBody, err := ioutil.ReadAll(resp.Body)

	log.Println(string(respBody))

	return
}

func (client *ExpoClient) SendPushNotifications() {
}
