package twitch

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

// Request performs a Twitch API request and returns the response body.
func Request(method string, url string, body io.Reader, auth bool, context bool) (string, error) {
	fmt.Println("Twitch Request:", url)

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return "", err
	}
	req.Header.Set("Client-Id", ClientID)
	req.Header.Set("Accept", "application/json")
	if auth && Token != "" {
		req.Header.Set("Authorization", "Bearer "+Token)
	}
	if context {
		req.Header.Set("Content-Type", "application/json")
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	rbody, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		return "", readErr
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return "", fmt.Errorf("twitch api %s returned %s: %s", url, resp.Status, string(rbody))
	}

	return string(rbody), nil
}
