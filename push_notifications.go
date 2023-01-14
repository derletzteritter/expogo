package expogo

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

const (
	DEFAULT_PRIORITY = "default"
	NORMAL_PRIORITY  = "normal"
	HIGH_PRIORITY    = "high"
)

var (
	ReceiptResponseSuccess = "ok"
	ReceiptResponseError   = "error"
)

type Notification struct {
	To         []string `json:"to"`
	Title      string   `json:"title,omitempty"`
	Body       string   `json:"body,omitempty"`
	TTL        int      `json:"ttl,omitempty"`
	Expiration int      `json:"expiration,omitempty"`
	Badge      int      `json:"badge,omitempty"`
	Priority   string   `json:"priority,omitempty"`
}

type PushTicket struct {
	Status  string `json:"status"`
	ID      string `json:"id"`
	Message string `json:"message"`
	Details struct {
		Error string `json:"error"`
	} `json:"details"`
}

type PushTicketError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type PushTicketResponse struct {
	Data []PushTicket `json:"data"`
}

func (client *ExpoClient) SendPushNotification(notification *Notification) *PushTicketResponse {
	body, err := json.Marshal(notification)
	if err != nil {
		log.Println(err.Error())
	}

	req, err := http.NewRequest("POST", client.host+client.pushPath, bytes.NewReader(body))
	if err != nil {
		log.Println(err.Error())
	}
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.httpClient.Do(req)
	if err != nil {
		log.Println(err.Error())
	}

	var receiptResponse PushTicketResponse
	// Use decoder instead of marshalling
	// since the response is a stream of JSON objects and not a JSON object in memory
	err = json.NewDecoder(resp.Body).Decode(&receiptResponse)
	if err != nil {
		log.Println("Failed to decode response body")
		log.Println(err.Error())
	}

	// Check for any errors in the response, loop through them
	// and return the first one
	for _, ticket := range receiptResponse.Data {
		if ticket.Status == ReceiptResponseError {
			log.Println("Error sending push notification")
		}
	}

	return nil
}
