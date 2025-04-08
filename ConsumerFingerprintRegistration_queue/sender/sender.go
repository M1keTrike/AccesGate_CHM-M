package sender

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Sender struct {
	url string
}

func NewSender(url string) *Sender {
	return &Sender{
		url: url,
	}
}

func (s *Sender) Send(body []byte) error {
	message := map[string]string{
		"message": string(body),
	}

	jsonMessage, err := json.Marshal(message)
	if err != nil {
		fmt.Printf("Error marshaling message to JSON: %v\n", err)
		return err
	}

	resp, err := http.Post(s.url, "application/json", bytes.NewBuffer(jsonMessage))
	if err != nil {
		fmt.Printf("Error sending message to API: %v\n", err)
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("API responded with non-success status: %d\n", resp.StatusCode)
		return fmt.Errorf("API error: code %d", resp.StatusCode)
	}

	fmt.Printf("Message successfully sent to API: %s\n", jsonMessage)
	return nil
}