package expogo

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	DefaultPriority = "default"
	NormalPriority  = "normal"
	HighPriority    = "high"
)

var (
	ReceiptResponseSuccess = "ok"
	ReceiptResponseError   = "error"

	DeviceNotRegistered = "DeviceNotRegistered"
	InvalidCredentials  = "InvalidCredentials"
	MessageTooBig       = "MessageTooBig"
	MessageRateExceeded = "MessageRateExceeded"
	InvalidData         = "InvalidData"
	MismatchSenderId    = "MismatchSenderId"
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
	Ticket  Notification
	Status  string `json:"status"`
	ID      string `json:"id"`
	Message string `json:"message"`
	Details struct {
		Error         string `json:"error"`
		ExpoPushToken string `json:"expoPushToken"`
	} `json:"details"`
}

type PushTicketError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type PushTicketResponse struct {
	Data   []PushTicket        `json:"data"`
	Errors []ServerTicketError `json:"errors"`
}

type ServerTicketError struct {
	Code        string `json:"code"`
	Message     string `json:"message"`
	IsTransient bool   `json:"isTransient"`
}

type ServerTicketErrorResponse struct {
	Errors []ServerTicketError `json:"errors"`
}

func (e *ServerTicketErrorResponse) Error() string {
	return e.Errors[0].Message
}

func NewServerTicketError(errors []ServerTicketError) *ServerTicketErrorResponse {
	return &ServerTicketErrorResponse{
		Errors: errors,
	}
}

func (client *ExpoClient) SendPushNotification(notification *Notification) ([]PushTicket, error) {
	body, err := json.Marshal(notification)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", client.host+client.pushPath, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Host", "exp.host")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Accept-Encoding", "gzip, deflate")
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	/*respBody, err := ioutil.ReadAll(resp.Body)
	log.Println(string(respBody)) */

	var receiptResponse PushTicketResponse
	// Use decoder instead of marshalling
	// since the response is a stream of JSON objects and not a JSON object in memory
	err = json.NewDecoder(resp.Body).Decode(&receiptResponse)
	if err != nil {
		log.Println("Failed to decode response body")
		return nil, err
	}

	if len(receiptResponse.Errors) > 0 {
		return nil, NewServerTicketError(receiptResponse.Errors)
	}

	if len(receiptResponse.Data) == 0 {
		return nil, fmt.Errorf("No tickets returned")
	}

	for i := range receiptResponse.Data {
		// Return tickets and the actual ticket data
		receiptResponse.Data[i].Ticket = *notification
	}

	return receiptResponse.Data, nil
}

func (client *ExpoClient) SendMultiplePushNotifications(notifications []*Notification) ([]PushTicket, error) {
	body, err := json.Marshal(notifications)
	if err != nil {
		return nil, err
	}

	log.Println("BODY")
	log.Println(string(body))

	// When sending multiple notifications, we should gzip the body to reduce upload bandwidth
	var buf bytes.Buffer
	gz := gzip.NewWriter(&buf)
	if _, err := gz.Write(body); err != nil {
		return nil, err
	}

	if err := gz.Close(); err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", client.host+client.pushPath, &buf)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Host", "exp.host")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Accept-Encoding", "gzip, deflate")
	req.Header.Add("Content-Encoding", "gzip")
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	respBody, _ := ioutil.ReadAll(resp.Body)
	log.Println("RESPONSE")
	log.Println(string(respBody))

	return nil, nil
}
