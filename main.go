package slack

import (
    "bytes"
    "encoding/json"
    "net/http"
	"fmt"
	"go-slack-message/models"
)

func Send(webhookUrl string, message models.SlackMessage) error {
	payload, err := json.Marshal(message)
    if err != nil {
        return err
    }

	resp, err := http.Post(webhookUrl, "application/json", bytes.NewBuffer(payload))
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
    }

    return nil
}