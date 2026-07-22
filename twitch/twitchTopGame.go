package twitch

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/asimshrestha2/stream-set/save"
)

type TopGamesResponse struct {
	Data []TopGame `json:"data"`
	Pagination struct {
		Cursor string `json:"cursor"`
	} `json:"pagination"`
}

type TopGame struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	BoxArtURL   string `json:"box_art_url"`
}

func GetTopGames(limit int, offset int) TopGamesResponse {
	url := TwitchAPIURL + "/games/top?first=" + strconv.Itoa(limit)
	if offset > 0 {
		url += "&after=" + strconv.Itoa(offset)
	}
	body, err := Request("GET", url, nil)
	if err != nil {
		fmt.Println("GetTopGames error:", err)
		return TopGamesResponse{}
	}
	var topGamesResponse TopGamesResponse
	if err := json.Unmarshal([]byte(body), &topGamesResponse); err != nil {
		fmt.Println("GetTopGames decode error:", err)
		return TopGamesResponse{}
	}
	return topGamesResponse
}

func GetTopGamesNames() {
	if save.GameListExist() {
		err := save.LoadGameList(&GameDB)
		if err != nil {
			fmt.Println(err)
		}
		return
	}

	tgr := GetTopGames(100, 0)
	time.Sleep(500 * time.Millisecond)
	tgr1 := GetTopGames(100, 0)
	tgr.Data = append(tgr.Data, tgr1.Data...)
	for _, g := range tgr.Data {
		tempGame := DBGame{TwitchName: g.Name}
		GameDB = append(GameDB, tempGame)
	}

	go save.SaveGameList(GameDB)
}
