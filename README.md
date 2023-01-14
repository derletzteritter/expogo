# ExpoGO

```
go get github.com/itschip/expogo
```

## Example

```go
package main

import (
        "log"

        "github.com/itschip/expogo"
)

func main() {
        client := expogo.NewExpoClient(nil)

        tickets, err := client.SendPushNotification(&expogo.Notification{
                To:       []string{"TOKEN_OR_TOKENS"},
                Title:    "Hello World",
                Body:     "This is a test notification",
                Priority: expogo.HighPriority,
        })
        if err != nil {
                log.Println(err.Error())
        }

        for _, ticket := range tickets {
                log.Println(ticket.ID)
                if ticket.Status == expogo.ReceiptResponseError {
                        log.Println(ticket.Details.Error)
                }
        }
}
```
