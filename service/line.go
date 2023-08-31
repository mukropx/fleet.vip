package service

import (
	"fmt"
	"net/http"
	"strings"
)

func SendLineNotifyMessage(message string) error {
	url := "https://notify-api.line.me/api/notify"
	method := "POST"

	payload := strings.NewReader(fmt.Sprintf("message=%s", message))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		return err
	}

	req.Header.Add("Authorization", "Bearer ")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	// Check the response status code for error handling
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("Line Notify API returned non-OK status: %s", res.Status)
	}

	return nil
}
