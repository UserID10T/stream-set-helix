package twitch

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
)

type SearchCategoriesResponse struct {
	Data []SearchCategory `json:"data"`
}

type SearchCategory struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	BoxArtURL   string `json:"box_art_url"`
}

// SearchGames searches Twitch categories and maps the first result to a DBGame.
func SearchGames(query string) (DBGame, error) {
	q, _ := url.Parse(query)
	uri := TwitchAPIURL + "/search/categories?query=" + q.String()
	body, err := Request("GET", uri, nil)
	if err != nil {
		return DBGame{}, err
	}
	var sr SearchCategoriesResponse
	if err := json.Unmarshal([]byte(body), &sr); err != nil {
		return DBGame{}, err
	}
	if len(sr.Data) > 0 {
		return DBGame{TwitchName: sr.Data[0].Name}, nil
	}
	fmt.Println("SearchGames: no categories found for", query)
	return DBGame{}, errors.New("Not Found")
}
