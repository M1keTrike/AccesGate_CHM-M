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
		fmt.Printf("Error al convertir mensaje a JSON: %v\n", err)
		return err
	}

	resp, err := http.Post(s.url, "application/json", bytes.NewBuffer(jsonMessage))
	if err != nil {
		fmt.Printf("Error al enviar mensaje a la API: %v\n", err)
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("La API respondió con un estado no exitoso: %d\n", resp.StatusCode)
		return fmt.Errorf("error en la API: código %d", resp.StatusCode)
	}

	fmt.Printf("Mensaje enviado exitosamente a la API: %s\n", jsonMessage)
	return nil
}
