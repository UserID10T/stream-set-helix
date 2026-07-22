package twitch

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Request performs a Twitch Helix API request and returns the response body.
func Request(method, rawURL string, body io.Reader) (string, error) {
	fmt.Println("Twitch Request:", rawURL)

	req, err := http.NewRequest(method, rawURL, body)
	if err != nil {
		return "", err
	}

	req.Header.Set("Client-Id", ClientID)
	req.Header.Set("Authorization", "Bearer "+Token)
	req.Header.Set("Accept", "application/json")
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	rbody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return "", fmt.Errorf("twitch api %s returned %s: %s", rawURL, resp.Status, string(rbody))
	}

	return string(rbody), nil
}

func CurrentUser() (*User, error) {
	body, err := Request("GET", TwitchAPIURL+"/users", nil)
	if err != nil {
		return nil, err
	}

	var resp UserResponse
	if err := jsonUnmarshal([]byte(body), &resp); err != nil {
		return nil, err
	}
	if len(resp.Data) == 0 {
		return nil, fmt.Errorf("no users returned")
	}
	return &resp.Data[0], nil
}

func CurrentChannel(userID string) (*Channel, error) {
	q := url.Values{}
	q.Set("broadcaster_id", userID)

	body, err := Request("GET", TwitchAPIURL+"/channels?"+q.Encode(), nil)
	if err != nil {
		return nil, err
	}

	var resp ChannelResponse
	if err := jsonUnmarshal([]byte(body), &resp); err != nil {
		return nil, err
	}
	if len(resp.Data) == 0 {
		return nil, fmt.Errorf("no channels returned")
	}
	return &resp.Data[0], nil
}

func UpdateCurrentChannel(userID, gameID string) error {
	q := url.Values{}
	q.Set("broadcaster_id", userID)
	payload := fmt.Sprintf(`{"game_id":%q}`, gameID)
	_, err := Request("PATCH", TwitchAPIURL+"/channels?"+q.Encode(), bytesNewBufferString(payload))
	return err
}

func SearchCategories(query string) ([]SearchCategory, error) {
	q := url.Values{}
	q.Set("query", query)

	body, err := Request("GET", TwitchAPIURL+"/search/categories?"+q.Encode(), nil)
	if err != nil {
		return nil, err
	}

	var resp SearchCategoriesResponse
	if err := jsonUnmarshal([]byte(body), &resp); err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func GetTopGames(limit int, after string) (TopGamesResponse, error) {
	q := url.Values{}
	if limit > 0 {
		q.Set("first", fmt.Sprintf("%d", limit))
	}
	if after != "" {
		q.Set("after", after)
	}

	body, err := Request("GET", TwitchAPIURL+"/games/top?"+q.Encode(), nil)
	if err != nil {
		return TopGamesResponse{}, err
	}

	var resp TopGamesResponse
	if err := jsonUnmarshal([]byte(body), &resp); err != nil {
		return TopGamesResponse{}, err
	}
	return resp, nil
}

func jsonUnmarshal(data []byte, v any) error {
	return json.Unmarshal(data, v)
}

func bytesNewBufferString(s string) io.Reader {
	return stringsNewReader(s)
}

func stringsNewReader(s string) io.Reader {
	return strings.NewReader(s)
}
