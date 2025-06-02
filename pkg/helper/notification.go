package helper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/farhansaleh/layanan_aptika_be/internal/domain"
)

func SendPushNotification(notifToken, title, body string) error {
	message := domain.ExpoPushMessage{
		To:    notifToken,
		Title: title,
		Body:  body,
	}
	jsonData, err := json.Marshal(message)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, "https://exp.host/--/api/v2/push/send", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to send notification: status %v", resp.Status)
	}

	fmt.Println("Push notification sent!")

	return nil
}